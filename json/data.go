package json

import (
	"encoding/json"
	"log"
	"os"
)

const (
	DBDIR string = "db/"
	FILENAME string = "album"
)

type Album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

func Create(albums []Album) {
	// JSONファイルを新規作成(既に存在する場合は上書き)
	file, err := os.Create(DBDIR + FILENAME + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// JSONファイルに書き込み
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(albums); err != nil {
		log.Fatal(err)
	}
}

func Read() (alubm []Album) {

	file, err := os.Open(DBDIR + FILENAME +  ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&alubm); err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", alubm)
	
	return
}

func Update(album Album) {
	albums := Read()
	albums = append(albums, album)
	Create(albums)
}