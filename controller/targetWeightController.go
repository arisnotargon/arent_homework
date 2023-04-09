package controller

import (
	"net/http"

	dbmodel "github.com/arisnotargon/arent_homework/model/db"
	requestmodel "github.com/arisnotargon/arent_homework/model/request"
	"github.com/davecgh/go-spew/spew"
	"github.com/shopspring/decimal"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TargetWeightController struct {
	Db *gorm.DB
}

func (ctr TargetWeightController) Store(c *gin.Context) {
	reqInfo := &requestmodel.TargetWeightStore{}
	e := c.Bind(reqInfo)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}

	targetWeightDecimal, e := decimal.NewFromString(reqInfo.TargetWeight)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}

	NowWeightDecimal, e := decimal.NewFromString(reqInfo.NowWeight)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}

	if targetWeightDecimal.GreaterThanOrEqual(NowWeightDecimal) {
		c.JSON(http.StatusBadRequest, gin.H{"messege": "target weight can not bigger than now weight"})
		return
	}

	userId, _ := c.Get("userId")
	userIdInt, ok := userId.(uint)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"messege": "user id invalid"})
		return
	}

	tx := ctr.Db.Begin()

	defer func() {
		if e == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	targetWeight := &dbmodel.TargetWeight{
		UserId:       uint(userIdInt),
		TargetWeight: reqInfo.TargetWeight,
		NowWeight:    reqInfo.NowWeight,
		Status:       uint8(dbmodel.TargetWeightStatusValid),
	}

	e = tx.Table("target_weights").Where("user_id=? and status=?", userId, dbmodel.TargetWeightStatusValid).
		Update("status", dbmodel.TargetWeightStatusInvalid).Error

	spew.Dump("after update===>", e)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}

	e = tx.Save(targetWeight).Error

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
