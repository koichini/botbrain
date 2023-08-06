package json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

func Create() {
	fmt.Println("test")
	// 書き込むデータ
	album := []Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	
	// JSONファイルを新規作成(既に存在する場合は上書き)
	file, err := os.Create("db/create.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// JSONファイルに書き込み
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(album); err != nil {
		log.Fatal(err)
	}

	fmt.Println("writted")

}

func Read() {

	file, err := os.Open("db/create.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var a []Album
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&a); err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", a)
	fmt.Println("readed")
}