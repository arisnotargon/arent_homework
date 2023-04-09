package controller

import (
	"net/http"
	"time"

	dbmodel "github.com/arisnotargon/arent_homework/model/db"
	requestmodel "github.com/arisnotargon/arent_homework/model/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AteFoodController struct {
	Db *gorm.DB
}

func (ctr AteFoodController) Store(c *gin.Context) {
	reqInfo := &requestmodel.AteFoodStore{}
	e := c.Bind(reqInfo)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}

	userId, _ := c.Get("userId")
	userIdInt, ok := userId.(uint)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"messege": "user id invalid"})
		return
	}

	ateAt, _ := time.Parse("2006-01-02", reqInfo.AteAt)

	ateFood := &dbmodel.AteFood{
		UserId: userIdInt,
		Name:   reqInfo.Name,
		Type:   reqInfo.Type,
		AteAt:  ateAt,
		Pic:    reqInfo.Pic,
	}

	e = ctr.Db.Save(ateFood).Error
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ctr AteFoodController) List(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdInt, ok := userId.(uint)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"messege": "user id invalid"})
		return
	}

	ateFoodList := make([]dbmodel.AteFood, 0)
	var count int64 = 0
	e := ctr.Db.Model(ateFoodList).Where("user_id=?", userIdInt).Count(&count).Error
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}

	e = ctr.Db.Where("user_id=?", userIdInt).Order("id desc").Find(&ateFoodList).Error
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": count, "ate_food_list": ateFoodList})

}
