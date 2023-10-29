package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string `json:"User"`
	Lists     []List `gorm:"foreignKey:UserRefer"`
}

type Game struct {
	ID            uint `json:"Id" gorm:"primaryKey"`
	CreatedAt     time.Time
	Name          string  `json:"Game"`
	Year          uint    `json:"Year"`
	Publisher     string  `json:"Publisher"`
	Developer     string  `json:"Dev"`
	Platform      string  `json:"Platform"`
	GameLink      string  `json:"GameLink"`
	DevLink       string  `json:"DevLink"`
	PublisherLink string  `json:"PublisherLink"`
	PlatformLink  string  `json:"PlatformLink"`
	Lists         []*List `gorm:"many2many:game_lists;"`
}

type List struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string `json:"List"`
	Public    bool   `json:"Public"`
	UserRefer uint
	Games     []*Game `gorm:"many2many:game_lists;"`
}
