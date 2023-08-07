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

// Update処理は重複判定なし
// DBに依存して処理できる想定
func Update(album Album) {
	var updateAlbums []Album
	albums := Read()
	// for i, a := range albums {
	// 	fmt.Println(i, a)
	// 	if a.ID == album.ID {
	// 		updateAlbums = append()
	// 	}
	// }
	updateAlbums = append(albums, album)
	Create(updateAlbums)
}