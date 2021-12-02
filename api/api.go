package api

import (
	"main/db"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {

	db.SetupDB()
	SetupDemoAPI(router)
	SetupAuthenAPI(router)
	SetupProductAPI(router)
	// SetupTransactionAPI(router)
}
