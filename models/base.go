package models

import "time"

type User struct {
	ID        string `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string `json:"User"`
	Lists     []List
}

type Game struct {
	ID            string `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time
	Name          string `json:"Game"`
	Year          uint   `json:"Year"`
	Publisher     string `json:"Publisher"`
	Developer     string `json:"Dev"`
	Platform      string `json:"Platform"`
	GameLink      string `json:"GameLink"`
	DevLink       string `json:"DevLink"`
	PublisherLink string `json:"PublisherLink"`
	PlatformLink  string `json:"PlatformLink"`
}

type List struct {
	ID        string `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string `json:"List"`
	Games     []Game
}

// SAMPLE  string `json:""`
