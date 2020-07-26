package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// config setup
	v := viper.New()
	loadConfigure(v)

	// http server setup
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/health", healthHandler)
		v1.POST("/mailgun/permanent-failure", permanentFailureHandler)
	}

	r.Run(":" + v.GetString("port"))
}
