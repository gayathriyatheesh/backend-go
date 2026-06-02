package main

type Product struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Name     string  `gorm:"column:name;type:varchar(100)" json:"name"`
	Price    float64 `gorm:"column:price" json:"price"`
	Category string  `gorm:"column:category" json:"category"`
}

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Email string `gorm:"column:email;uniqueIndex" json:"email"`
}
