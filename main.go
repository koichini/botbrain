package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
    FILENAME string = "create"
)

func main() {
	// albums := json.Read()
    // fmt.Printf("%T\n", albums)
	// album := json.Album{ID: "4", Title: "Copper Copper seas", Artist: "blue ocean", Price: 1.99}
    // json.Update(album)

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
    fmt.Println("======= main end =======")
}
