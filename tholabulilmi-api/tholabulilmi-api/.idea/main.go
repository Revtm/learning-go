package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DoaDto struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Arabic    string `json:"arabic"`
	Translate string `json:"translate"`
}

var doaDtos = []DoaDto{
	{0, "Doa Sebelum Makan", "Bismillah", "Dengan menyebut nama Allah"},
	{1, "Doa Setelah Makan", "Alhamdulillah", "Segala puji bagi Allah"},
}

func getDoaDtos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, doaDtos)
}

func getDoaDto(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Format Id salah!"})
		return
	}

	for _, doa := range doaDtos {
		if doa.Id == id{
			c.IndentedJSON(http.StatusOK, doa)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Doa tidak ditemukan!"})
}

func addDoaDto(c *gin.Context){
	var doaDto DoaDto
	err := c.BindJSON(&doaDto)

	if err != nil {
		return
	}

	doaDtos = append(doaDtos, doaDto)
	c.IndentedJSON(http.StatusCreated, doaDto)
}

func main() {
	router := gin.Default()
	router.GET("/doa", getDoaDtos)
	router.POST("/doa", addDoaDto)
	router.GET("/doa/:id",getDoaDto)

	router.Run("localhost:8080")
}
