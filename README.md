# nim-griz
- نیم گریز یک نرم افزار برای نگهداری مجموعه از تمامی بازی ها از nes تا ps3 است
- نوشته شده با go , chi router , gorm , sqlite , htmx

# .env sample
```
PORT=
```

# Models
```go
type User struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string `json:"User"`
	Lists     []List
}

type Game struct {
	ID            string `gorm:"primaryKey"`
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
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string `json:"List"`
	Games     []Game
}
```

# Show Case