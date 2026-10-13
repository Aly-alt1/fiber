package domain

type Mouse struct {
	ID     string  `gorm:"column ID" json:"id"`
	Avatar string  `gorm:"column AVATAR" json:"avatar"`
	Name   string  `gorm:"column NAME" json:"name"`
	Amount int     `gorm:"column AMOUNT" json:"amount"`
	Vendor string  `gorm:"column VENDOR" json:"vendor"`
	Status string  `gorm:"column STATUS" json:"status"`
	Rating float64 `gorm:"column RATING" json:"rating"`
}
