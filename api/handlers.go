package api

import (
	"fmt"
	"github.com/Bnei-Baruch/udb/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var u models.User
	if c.BindJSON(&u) == nil {
		udb := c.MustGet("UDB").(*gorm.DB)
		udb.Create(&u)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func GetUsers(c *gin.Context) {
	var u []models.User
	udb := c.MustGet("UDB").(*gorm.DB)
	if err := udb.Find(&u).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, u)
	}
}

func GetUser(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("id"), 10, 0)
	if e != nil {
		NewBadRequestError(errors.Wrap(e, "id expects int64")).Abort(c)
		return
	}
	udb := c.MustGet("UDB").(*gorm.DB)
	var u models.User
	if err := udb.Where("id = ?", id).First(&u).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, u)
	}
}

func GetIngest(c *gin.Context) {
	id := c.Params.ByName("id")
	udb := c.MustGet("UDB").(*gorm.DB)
	var ingest models.Ingest
	if err := udb.Where("capture_id = ?", id).First(&ingest).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, ingest)
	}
}

func GetTrimmer(c *gin.Context) {
	id := c.Params.ByName("id")
	udb := c.MustGet("UDB").(*gorm.DB)
	var t models.Trimmer
	if err := udb.Where("trim_id = ?", id).First(&t).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, t)
	}
}

func PutTrimmer(c *gin.Context) {
	var t models.Trimmer
	if c.BindJSON(&t) == nil {
		udb := c.MustGet("UDB").(*gorm.DB)
		udb.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&t)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}