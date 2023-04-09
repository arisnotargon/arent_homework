package controller

import (
	"net/http"

	requestmodel "github.com/arisnotargon/arent_homework/model/request"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"

	dbmodel "github.com/arisnotargon/arent_homework/model/db"
	"github.com/arisnotargon/arent_homework/utils"
)

type UserController struct {
	Db *gorm.DB
}

func (ctr UserController) SignUp(c *gin.Context) {
	reqInfo := &requestmodel.SignUp{}
	e := c.Bind(reqInfo)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}

	spew.Dump("after", reqInfo)

	hashedPassword, e := bcrypt.GenerateFromPassword([]byte(reqInfo.Password), bcrypt.DefaultCost)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messege": "failed to generate hashed password"})
		return
	}

	User := &dbmodel.User{
		Username: reqInfo.Username,
		Password: string(hashedPassword),
	}

	var count int64 = 0
	ctr.Db.Where("username = ?", reqInfo.Username).Table("users").Count(&count)

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"messege": "username Duplicate"})
		return
	}

	e = ctr.Db.Create(User).Error
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messege": "failed to create user"})
		return
	}

	// jwt トケン作成
	jwtToken, e := utils.CreateToken(1440, int(User.ID))

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messege": "failed to create jwt token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

func (ctr UserController) Login(c *gin.Context) {
	reqInfo := &requestmodel.Login{}
	e := c.Bind(reqInfo)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}

	User := &dbmodel.User{}
	e = ctr.Db.Where("username = ?", reqInfo.Username).First(User).Error
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
	}

	// 暗号番号チェック
	if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(reqInfo.Password)); err != nil {
		c.JSON(401, gin.H{"error": "invalid password"})
		return
	}

	// jwt トケン作成
	jwtToken, e := utils.CreateToken(1440, int(User.ID))

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messege": "failed to create jwt token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

func (ctr UserController) GetUserInfo(c *gin.Context) {
	userId, _ := c.Get("userId")
	username, _ := c.Get("username")

	c.JSON(http.StatusOK, gin.H{"username": username, "user_id": userId})
}
