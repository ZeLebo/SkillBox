package database

type Model struct {
	// primary key in gorm
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

// Friend for one-many relation
type Friend struct {
	First  int `gorm:"column:friend1"`
	Second int `gorm:"column:friend2"`
}
