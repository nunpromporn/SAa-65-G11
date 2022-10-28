package controller

 

import (

             "github.com/Panadsada/sa-65-example/entity"

           "github.com/gin-gonic/gin"

           "net/http"

)
// POST /tenderness

// func CreateTenderness(c *gin.Context) {

// 	var tenderness entity.Tenderness

// 	if err := c.ShouldBindJSON(&tenderness); err != nil {

// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 		   return

// 	}



// 	if err := entity.DB().Create(&tenderness).Error; err != nil {

// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 		   return

// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": tenderness})

// }
// GET /tenderness/:id

// func GetTenderness(c *gin.Context) {

// 	var tenderness entity.Tenderness

// 	id := c.Param("id")

// 	if err := entity.DB().Raw("SELECT * FROM tendernesses WHERE id = ?", id).Scan(&tenderness).Error; err != nil {

// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 		   return

// 	}



// 	c.JSON(http.StatusOK, gin.H{"data": tenderness})

// }
// GET /tenderness

func ListTenderness(c *gin.Context) {

	var tenderness []entity.Tenderness

	if err := entity.DB().Raw("SELECT * FROM tendernesses").Scan(&tenderness).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}

	c.JSON(http.StatusOK, gin.H{"data": tenderness})

}
