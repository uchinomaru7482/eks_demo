package main

import (
	"fmt"
	"log"
	"eks_demo/configs"
	"eks_demo/internal/app/handler"
	"eks_demo/internal/pkg/utils"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 設定情報
	appConfig := &configs.AppConfig{}
	appConfig.LoadAppConfig()
	// タイムゾーン設定
	if err := utils.InitTimeLocationAsiaTokyo(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	handler.RegisterHandler(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", strconv.Itoa(appConfig.ListenPort))))
}
