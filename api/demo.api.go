package api

import (
	"fmt"
	"main/db"
	"main/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupDemoAPI(router *gin.Engine) {
	demoAPI := router.Group("/api/v1")
	{
		demoAPI.GET("/demo", getDemo)
		demoAPI.GET("/demo/:id", getDemoById)
		demoAPI.POST("/demo", createDemo)
		demoAPI.PUT("/demo", editDemo)
		demoAPI.DELETE("/demo/:id", deleteDemo)
	}
}

func panicQueryHandle(c *gin.Context) {
	errResult := recover()
	c.JSON(200, gin.H{"result": "nok", "error": fmt.Sprintf("%s", errResult)})
}

func getDemo(c *gin.Context) {
	defer panicQueryHandle(c)
	var demo []model.Demo
	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		if err := db.GetDB().Where("name like ?", keyword).Order("created_at DESC").Find(&demo).Error; err != nil {
			panic(err)
		}
	} else {
		if err := db.GetDB().Order("created_at DESC").Find(&demo).Error; err != nil {
			panic(err)
		}
	}
	c.JSON(200, demo)
}

func getDemoById(c *gin.Context) {
	defer panicQueryHandle(c)
	var demo model.Demo
	if err := db.GetDB().Where("id = ?", c.Param("id")).First(&demo).Error; err != nil {
		panic(err)
	}
	c.JSON(200, demo)
}

func createDemo(c *gin.Context) {
	demo := model.Demo{}
	demo.Name = c.PostForm("name")
	demo.CreatedAt = time.Now()
	db.GetDB().Create(&demo)
	c.JSON(http.StatusOK, gin.H{"result": demo})
}

func editDemo(c *gin.Context) {
	var demo model.Demo
	id, _ := strconv.ParseUint(c.PostForm("id"), 10, 32)
	demo.ID = uint32(id)
	demo.Name = c.PostForm("name")
	demo.CreatedAt = time.Now()
	db.GetDB().Updates(&demo)
	c.JSON(http.StatusOK, gin.H{"result": demo})
}

func deleteDemo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	db.GetDB().Delete(&model.Demo{}, id)
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}
