# nim-griz
- نیم گریز یک نرم افزار برای نگهداری مجموعه از تمامی بازی ها از nes تا ps3 است
- نوشته شده با go , chi router , gorm , sqlite , htmx

# .env sample
```
PORT=
```

# Models
```
To unmarshal JSON into an interface value, Unmarshal stores one of these in the interface value:

bool, for JSON booleans
float64, for JSON numbers
string, for JSON strings
[]interface{}, for JSON arrays
map[string]interface{}, for JSON objects
nil for JSON null
```

```go
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
```

# Show Case