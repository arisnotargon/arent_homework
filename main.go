package main

import (
	_ "time"

	"github.com/arisnotargon/arent_homework/controller"
	"github.com/arisnotargon/arent_homework/middleware"
	dbmodel "github.com/arisnotargon/arent_homework/model/db"
	"github.com/davecgh/go-spew/spew"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/arisnotargon/arent_homework/config"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		panic("read config failed:" + err.Error())
	}

	spew.Dump("main conf===>", config.Config)

	db, err := gorm.Open(mysql.Open(config.Config.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// create table
	db.AutoMigrate(
		&dbmodel.User{},
		&dbmodel.Notification{},
		&dbmodel.TargetWeight{},
		&dbmodel.BodyInfo{},
		&dbmodel.AteFood{},
	)

	r := gin.Default()

	usrCtr := &controller.UserController{
		Db: db,
	}

	ntfCtr := &controller.NotificationController{
		Db: db,
	}

	twCtr := &controller.TargetWeightController{
		Db: db,
	}
	bfiCtr := &controller.BodyInfoController{
		Db: db,
	}

	afCtr := &controller.AteFoodController{
		Db: db,
	}

	r.POST("signup", usrCtr.SignUp)
	r.POST("login", usrCtr.Login)

	r.Group("user").Use(middleware.JwtMiddlewareGen(db)).
		GET("info", usrCtr.GetUserInfo)

	r.Group("notification").Use(middleware.JwtMiddlewareGen(db)).
		GET("unread_num", ntfCtr.GetUnreadNum).
		GET("list", ntfCtr.List)

	r.Group("target_weight").Use(middleware.JwtMiddlewareGen(db)).
		POST("", twCtr.Store)

	r.Group("body_info").Use(middleware.JwtMiddlewareGen(db)).
		POST("", bfiCtr.Store).
		GET("banner_info", bfiCtr.BannerInfo)

	r.Group("ate_food").Use(middleware.JwtMiddlewareGen(db)).
		POST("", afCtr.Store).
		GET("", afCtr.List)

	r.Run(":9999")
}
