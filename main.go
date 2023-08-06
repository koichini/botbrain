package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/koichini/botbrain/json"
)

const (
    FILENAME string = "create"
)

func main() {
	albums := json.Read()
    fmt.Println("==============")
    fmt.Printf("%T\n", albums)
    fmt.Println(albums)
    fmt.Println("==============")
	album := json.Album{ID: "4", Title: "Copper seas", Artist: "blue ocean", Price: 1.99}
    json.Update(album)

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
