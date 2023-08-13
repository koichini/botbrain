package firestore

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

type User struct {
    First string `firestore:"first"`
    Last  string `firestore:"last"`
    Born  int    `firestore:"born"`
}

func Client(projectID string) (ctx context.Context, client *firestore.Client){
	// Use the application default credentials
	ctx = context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(client)
	return
}

func Read(collection string, client *firestore.Client) {
	ctx := context.Background()
	iter := client.Collection(collection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}

func Write(collection string, data interface{}, client *firestore.Client) {
	ctx := context.Background()
	_, _, err1 := client.Collection("users").Add(ctx, data)
	if err1 != nil {
        log.Fatalf("Failed adding alovelace: %v", err1)
	}
}