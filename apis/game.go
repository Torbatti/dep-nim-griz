package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/torbatti/nim-griz/models"
	"gorm.io/gorm"
)

// TODO: FIGURE OUT why is this not working
type JsonGame struct {
	Game          string
	GameLink      string
	Year          float64
	Dev           string
	DevLink       string
	Publisher     string
	PublisherLink string
	Platform      string
	PlatformLink  string
}

func Start(db *gorm.DB, path string) {
	var filePaths []string

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		if info.Name() != path {
			filePaths = append(filePaths, info.Name())
		}
		return nil
	})

	// UnmarshalJson(db, "datas/3DSGames.json", &modelGames)
	var modelGames []models.Game
	for _, v := range filePaths {
		UnmarshalJson(db, path+"/"+v, &modelGames)
	}
	db.CreateInBatches(&modelGames, 1023)

}

func UnmarshalJson(db *gorm.DB, path string, modelGames *[]models.Game) {

	jsonData, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("read error: %s\n", err)
	}

	var JsonGames []JsonGame
	err = json.Unmarshal([]byte(jsonData), &JsonGames)
	if err != nil {
		fmt.Println(path)

		fmt.Printf("unmarshal error: %s\n", err)
		return
	}

	for _, v := range JsonGames {
		guyear := uint(int(v.Year))

		game := models.Game{
			Name:          v.Game,
			GameLink:      v.GameLink,
			Year:          guyear,
			Developer:     v.Dev,
			DevLink:       v.DevLink,
			Publisher:     v.Publisher,
			PublisherLink: v.PublisherLink,
			Platform:      v.Platform,
			PlatformLink:  v.PlatformLink,
		}
		*modelGames = append(*modelGames, game)
	}
	println("- X " + path)
}
