package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
)

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
var temperatures = []temperature{
	{ID: "1", Celsius: 22, Fahrenheit: 71.6, Kelvin: 295.15},
}

var Temperatures = []Temperature{
	{
		Celsius:    22,
		Fahrenheit: 71.6,
		Kelvin:     295.15},
}

func main() {

	router := gin.Default()

	router.GET("/temperatures", getTemperature)
	router.GET("/c=:c", postTemperature)

	router.Run("localhost:8080")
}

// Get temperature
func getTemperature(c *gin.Context) {
	c.XML(http.StatusOK, Temperatures)
}

// func getStatus(c *gin.Context) {
// 	c.XML(http.StatusOK, c.Status(http.StatusOK))
// }

// postAlbums adds an album from JSON received in the request body.
func postTemperature(c *gin.Context) {
	var temp Temperature
	cel, err := strconv.Atoi(c.Param("c"))
	if err != nil {
		golog.Info(err)
		// c.XML(http.StatusInternalServerError, err)
		return
	}
	// C = (F-32)/1.8
	// K = C + 273.15
	temp.Celsius = float32(cel)
	temp.Kelvin = float32(cel) + 273.15
	temp.Fahrenheit = (float32(cel) * 1.8) + 32

	// if err := c.BindXML(&temp); err != nil {
	// 	c.XML(http.StatusInternalServerError, err)
	// 	return
	// }

	// Add the new album to the slice.
	Temperatures = append(Temperatures, temp)
	c.XML(http.StatusOK, Temperatures)

}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
