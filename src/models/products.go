package models

import "time"

type Product struct {
	Id           int        `db:"id" json:"id"`
	Name         *string    `db:"name" json:"name" form:"name"`
	BasePrice    *int       `db:"basePrice" json:"basePrice" form:"basePrice"`
	Description  *string    `db:"description" json:"description" form:"description"`
	Image        *string    `db:"image" json:"image" form:"image"`
	IsBestSeller *bool      `db:"isBestSeller" json:"isBestSeller" form:"isBestSeller"`
	Discount     *int       `db:"discount" json:"discount" form:"discount"`
	CreatedAt    *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt    *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllProducts() ([]Product, error) {
	sql := `SELECT * FROM "products" ORDER BY "id" ASC`
	data := []Product{}
	err := db.Select(&data, sql)
	return data, err
}

func FindOneProduct(id int) (Product, error) {
	sql := `SELECT * FROM "products" WHERE "id"=$1`
	data := Product{}
	err := db.Get(&data, sql, id)
	return data, err
}
