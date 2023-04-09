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

	// spew.Dump("db===>", db)
	db.AutoMigrate(&dbmodel.User{})

	r := gin.Default()

	usrCtr := &controller.UserController{
		Db: db,
	}

	r.POST("signup", usrCtr.SignUp)
	r.POST("login", usrCtr.Login)

	r.Group("user").Use(middleware.JwtMiddlewareGen(db)).GET("info", usrCtr.GetUserInfo)

	r.Run(":9999")
}
