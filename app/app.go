package app

import (
	"github.com/gin-gonic/gin"
	"github.com/psinthorn/gostack_users-api/configs"

	"github.com/psinthorn/gostack_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApp(port string) {
	router.LoadHTMLGlob("views/*/*.html")
	router.Static("/assets/", "./assets/")
	urlsMapping()

	logger.Info("Start GoGolang Web Application...")
	router.Run(":" + configs.ServerPort.PortSelector(port))
}
