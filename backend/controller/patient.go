package controller

 

import (

            "github.com/Panadsada/sa-65-example/entity"

           "github.com/gin-gonic/gin"

           "net/http"
		   //"golang.org/x/crypto/bcrypt"
		   

)
// // POST /patients

// func CreatePatient(c *gin.Context) {

// 	var patient entity.Patient

// 	if err := c.ShouldBindJSON(&patient); err != nil {

// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 		   return

// 	}



// 	if err := entity.DB().Create(&patient).Error; err != nil {

// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 		   return

// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": patient})

// }
// // GET /patient/:id

// func GetPatient(c *gin.Context) {

// 	var patient entity.Patient

// 	id := c.Param("id")

// 	if err := entity.DB().Raw("SELECT * FROM patients WHERE id = ?", id).Scan(&patient).Error; err != nil {

// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 		   return

// 	}



// 	c.JSON(http.StatusOK, gin.H{"data": patient})

// }
// GET /patients

func ListPatients(c *gin.Context) {

	var patients []entity.Patient

	if err := entity.DB().Raw("SELECT * FROM patients").Scan(&patients).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": patients})

}
