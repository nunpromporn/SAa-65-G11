package main

import (
	"github.com/Panadsada/sa-65-example/controller"
	"github.com/Panadsada/sa-65-example/entity"
	"github.com/Panadsada/sa-65-example/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{

			// User Routes
			protected.GET("/users", controller.ListUsers)
			protected.GET("/user/:id", controller.GetUser)
			
			// Patient Routes
			protected.GET("/patients", controller.ListPatients)

			// Department Routes
			protected.GET("/departments", controller.ListDepartments)

			// Doctor Routes
			protected.GET("/doctors", controller.ListDoctors)

			// Tenderness Routes
			protected.GET("/tenderness", controller.ListTenderness)

			// Symptom Routes
			protected.POST("/createsymptoms", controller.CreateSymptom)
			protected.GET("/symptoms", controller.ListSymptom)


		}
	}

	
	//ไว้สมัคร user doctor
	//r.POST("/doctors", controller.CreateDoctor)
	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}