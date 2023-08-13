package main

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"

	"github.com/koichini/botbrain/firestore"
)

const (
	projectID = "twbot-395603"
)

type User struct {
    First string `firestore:"first"`
    Last  string `firestore:"last"`
    Born  int    `firestore:"born"`
}

func main() {
    fmt.Println("======= init firestore =======")
	// Use the application default credentials
	ctx, client := firestore.Client(projectID)
	fmt.Println(client)
	fmt.Println("===== type =====")
	P(client)
	P(ctx)

	firestore.Read("users", client)

	// user := User{
	// 	First: "ziro",
    //     Last:  "suzuki",
    //     Born:  1827,
	// }
	// firestore.Write("users", &user, client)

	defer client.Close()

    fmt.Println("======= gin =======")
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
    fmt.Println("======= main end =======")
}

func P(t interface{}) {
    fmt.Println(reflect.TypeOf(t))
}