// Package fcServer - Interact with flowercare BLE API's
package fcServer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (fcs *FlowerCare) healthCheckAPI() {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		if fcs.errorCount > 0 {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.AbortWithStatus(http.StatusOK)
	})

	r.Run()
}
