package controller

import (
	"net/http"

	dbmodel "github.com/arisnotargon/arent_homework/model/db"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NotificationController struct {
	Db *gorm.DB
}

func (ctr NotificationController) GetUnreadNum(c *gin.Context) {
	spew.Dump("in GetUnreadNum")
	userId, _ := c.Get("userId")
	var count int64 = 0
	e := ctr.Db.Table("notifications").Where("user_id=? and read_at is null", userId).Count(&count).Error
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messege": e.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"unread_notification_num": count})
}

func (ctr NotificationController) List(c *gin.Context) {
	userId, _ := c.Get("userId")

	list := make([]dbmodel.Notification, 0)
	e := ctr.Db.Where("user_id=? ", userId).Find(&list).Error
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messege": e.Error()})
		return
	}
	spew.Dump("notification list====>", list)

	c.JSON(http.StatusOK, gin.H{"notification_list": list})
}
