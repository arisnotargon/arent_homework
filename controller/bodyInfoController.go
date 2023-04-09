package controller

import (
	"net/http"
	"time"

	dbmodel "github.com/arisnotargon/arent_homework/model/db"
	requestmodel "github.com/arisnotargon/arent_homework/model/request"
	responseModel "github.com/arisnotargon/arent_homework/model/response"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BodyInfoController struct {
	Db *gorm.DB
}

func (ctr BodyInfoController) Store(c *gin.Context) {
	spew.Dump("in BodyInfoController Store")
	reqInfo := &requestmodel.BodyInfoStore{}
	e := c.Bind(reqInfo)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}

	weight, e := decimal.NewFromString(reqInfo.Weight)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}
	if weight.LessThanOrEqual(decimal.NewFromInt(0)) {
		c.JSON(http.StatusBadRequest, gin.H{"messege": "weight can not less than 0"})
		return
	}

	fatRate, e := decimal.NewFromString(reqInfo.FatRate)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}
	if fatRate.LessThanOrEqual(decimal.NewFromInt(0)) || fatRate.GreaterThanOrEqual(decimal.NewFromInt(100)) {
		c.JSON(http.StatusBadRequest, gin.H{"messege": "fat rate range error"})
		return
	}

	userId, _ := c.Get("userId")
	userIdInt, ok := userId.(uint)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"messege": "user id invalid"})
		return
	}

	bodyInfo := &dbmodel.BodyInfo{
		UserId:    userIdInt,
		Weight:    reqInfo.Weight,
		FatRate:   reqInfo.FatRate,
		RecodedAt: time.Now(),
	}

	e = ctr.Db.Save(bodyInfo).Error
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messege": e.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// BannerInfo 達成率、体重体脂肪率記録取得
func (ctr BodyInfoController) BannerInfo(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdInt, ok := userId.(uint)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"messege": "user id invalid"})
		return
	}
	now := time.Now()
	spew.Dump(now.Format("2006-01-02"))

	responseObj := &responseModel.BannerInfo{
		NowDate: now.Format("2006-01-02"),
	}

	bodyInfoList := make([]dbmodel.BodyInfo, 0)

	e := ctr.Db.Where("user_id=?", userIdInt).Order("id desc").Find(&bodyInfoList).Error
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messege": e.Error()})
		return
	}

	targetWeight := &dbmodel.TargetWeight{}
	e = ctr.Db.Where("user_id=? and status=?", userIdInt, dbmodel.TargetWeightStatusValid).Find(targetWeight).Error
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messege": e.Error()})
		return
	}
	if len(bodyInfoList) > 0 && targetWeight.ID > 0 {
		lastBodyInfo := bodyInfoList[0]
		lastWetightDecimal, e := decimal.NewFromString(lastBodyInfo.Weight)
		if e != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"messege": e.Error()})
			return
		}
		targetWeightDecimal, e := decimal.NewFromString(targetWeight.TargetWeight)
		if e != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"messege": e.Error()})
			return
		}

		startWeightDecimal, e := decimal.NewFromString(targetWeight.NowWeight)
		if e != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"messege": e.Error()})
			return
		}

		if lastWetightDecimal.LessThanOrEqual(targetWeightDecimal) {
			responseObj.CompleteRate = "100"
		} else {
			// (lastWetight - targetWeight) / (startWeight - targetWeight) 小数2桁を保留
			responseObj.CompleteRate = lastWetightDecimal.Sub(targetWeightDecimal).
				Div(startWeightDecimal.Sub(targetWeightDecimal)).Mul(decimal.NewFromInt(100)).Round(2).String()
		}
	}

	responseObj.BodyHistory = bodyInfoList
	c.JSON(http.StatusOK, responseObj)
}
