//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package routers

import (
	ap "github.com/moemoe89/simple-go-clean-arch/api"
	mw "github.com/moemoe89/simple-go-clean-arch/api/middleware"
	"github.com/moemoe89/simple-go-clean-arch/api/v1/user"
	conf "github.com/moemoe89/simple-go-clean-arch/config"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// GetRouter will create a variable that represent the gin.Engine
func GetRouter() *gin.Engine {

	dbR, dbW, err := conf.InitDB()
	if err != nil {
		panic(err)
	}

	lang, err := conf.InitLang()
	if err != nil {
		panic(err)
	}

	log := conf.InitLog()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(mw.CORS)
	r.GET("/", ap.Ping)
	r.GET("/ping", ap.Ping)

	api := r.Group("/api")
	apiV1 := api.Group("/v1")

	userRepo := user.NewMysqlRepository(dbR, dbW)
	userSvc := user.NewService(log, userRepo)
	usr := user.NewUserCtrl(lang, log, userSvc)

	apiV1.POST("/user", usr.Create)
	apiV1.GET("/user", usr.List)
	apiV1.GET("/user/:id", usr.Detail)
	apiV1.PUT("/user/:id", usr.Update)
	apiV1.DELETE("/user/:id", usr.Delete)

	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
