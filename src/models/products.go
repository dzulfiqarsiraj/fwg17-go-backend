package models

import (
	"database/sql"
	"time"
)

type Product struct {
	Id           int          `db:"id" json:"id"`
	Name         string       `db:"name" json:"name" form:"name"`
	BasePrice    string       `db:"basePrice" json:"basePrice" form:"basePrice"`
	Description  string       `db:"description" json:"description" form:"description"`
	Image        string       `db:"image" json:"image" form:"image"`
	IsBestSeller string       `db:"isBestSeller" json:"isBestSeller" form:"isBestSeller"`
	Discount     string       `db:"discount" json:"discount" form:"discount"`
	CreatedAt    time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt    sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoProduct struct {
	Data  []Product
	Count int
}

func FindAllProducts(limit int, offset int) (InfoProduct, error) {
	sql := `SELECT * FROM "products" LIMIT $1 OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "products"`
	result := InfoProduct{}
	dataProduct := []Product{}
	err := db.Select(&dataProduct, sql, limit, offset)

	result.Data = dataProduct
	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)
	return result, err

}
