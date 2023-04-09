package middleware

import (
	"net/http"

	"github.com/arisnotargon/arent_homework/utils"
	"github.com/davecgh/go-spew/spew"
	"gorm.io/gorm"

	"strings"

	dbmodel "github.com/arisnotargon/arent_homework/model/db"
	"github.com/gin-gonic/gin"
)

func JwtMiddlewareGen(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := strings.Split(c.Request.Header.Get("Authorization"), " ")
		if len(authorizationHeader) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"messsege": "Unauthorized"})
			return
		}

		token := authorizationHeader[1]
		spew.Dump(token)

		claims, e := utils.VarifyToken(token)

		if e != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"messsege": e.Error()})
			return
		}

		User := &dbmodel.User{}

		e = db.Where("id = ?", claims.Id).Find(User).Error
		if e != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"messsege": e.Error()})
			return
		}

		c.Set("userId", User.ID)
		c.Set("username", User.Username)

		c.Next()
	}
}
