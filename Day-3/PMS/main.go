// package main

// import (
// 	"github.com/gin-gonic/gin"
// )


// type car struct{
// 	ID string `json:"id"`
// 	Number string `json:"number"`
// 	Model string `json:"model"`
// 	Class string `json:"class"`
// }
// func main() {
// 	r := gin.Default()
// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Hello World!",
// 		})
// 	})
// 	r.Run(":8000") 
// }


package main

import (
	//"context"
	"fmt"
	//"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Config




type Car struct {
	ID     string `json:"id"`
	Number string `json:"number"`
	Model  string `json:"model"`
	Type   string `json:"type"`
}

// Global slice of cars
var cars []Car = []Car{
	{ID: "id001", Number: "KA03 A1010", Model: "Maruti Suzuki", Type: "CUV"},
	{ID: "id002", Number: "KA40 Z2233", Model: "Toyota Innova", Type: "SUV"},
}

// POST /cars
func createCar(c *gin.Context) {
	var jbodyCar Car

	// Bind JSON body to jbodyCar
	if err := c.BindJSON(&jbodyCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new car object
	createdCar := Car{ID: jbodyCar.ID, Number: jbodyCar.Number, Model: jbodyCar.Model, Type: jbodyCar.Type}

	// Append to the global cars slice
	cars = append(cars, createdCar)

	// Return created car
	c.JSON(http.StatusCreated, gin.H{
		"message": "Car created successfully",
		"car":     createdCar,
	})
}

// GET /cars
func readAllCars(c *gin.Context) {
	// var cars []Car = []Car{
	// 	{ID: "id001", Number: "KA03 A1010", Model: "Maruti Suzuki", Type: "CUV"},
	// 	{ID: "id002", Number: "KA40 Z2233", Model: "Toyoto Innova", Type: "SUV"},
	// }

	c.JSON(http.StatusOK, cars)
}

// GET /cars/:id
func readCarById(c *gin.Context) {
	id := c.Param("id")
	
	    
		for _, s := range cars{
			if s.ID == id {
				var car Car = Car{ID: id, Number: s.Number, Model: s.Model, Type: s.Type}
				c.JSON(http.StatusOK, car)
				return
			} else {
				c.JSON(http.StatusNotFound, gin.H{"error": "Not found ID"})
			}
			
		}
	
	// var car Car = Car{ID: id, Number: "KA03 A1010", Model: "Maruti Suzuki", Type: "CUV"}
	// c.JSON(http.StatusOK, car)
}

// PUT /cars/:id
func updateCar(c *gin.Context) {
	id := c.Param("id")
	var jbodyCar Car
	if err := c.BindJSON(&jbodyCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var oldCar Car = Car{ID: id, Number: "KA03 A1010", Model: "Maruti Suzuki", Type: "CUV"}

	oldCar.Number = jbodyCar.Number
	oldCar.Model = jbodyCar.Model
	oldCar.Type = jbodyCar.Type

	// Return updated car
	c.JSON(http.StatusOK, gin.H{
		"message": "Car updated successfully",
		"car":     oldCar,
	})
}

// DELETE /cars/:id
func deleteCar(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}

func main() {
	// Set up Gin router
	r := gin.Default()
	// CORS Configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// // Routes
	// r.POST("/cars", createCar)
	// r.GET("/cars", readAllCars)
	r.GET("/cars/:id", readCarById)
	// r.PUT("/cars/:id", updateCar)
	// r.DELETE("/cars/:id", deleteCar)



	// Start server
	r.Run(":8080")
}