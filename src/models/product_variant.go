package models

import (
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type ProductVariant struct {
	Id              int        `db:"id" json:"id"`
	Name            *string    `db:"name" json:"name" form:"name"`
	AdditionalPrice *int       `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
	CreatedAt       *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt       *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllProductVariant(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "productVariant"
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productVariant"`
	result := services.Info{}
	data := []ProductVariant{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneProductVariant(id int) (ProductVariant, error) {
	sql := `SELECT * FROM "productVariant" WHERE "id"=$1`
	data := ProductVariant{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneProductVariantByName(name string) (ProductVariant, error) {
	sql := `SELECT * FROM "productVariant" WHERE "name" = $1`
	data := ProductVariant{}
	err := db.Get(&data, sql, name)
	return data, err
}

func FindProductVariantNameById(id int) (ProductVariant, error) {
	sql := `SELECT "name" FROM "productVariant" WHERE "id" = $1`
	data := ProductVariant{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductVariant(data ProductVariant) (ProductVariant, error) {
	sql := `
	INSERT INTO "productVariant" ("name","additionalPrice") VALUES
	(:name, :additionalPrice)
	RETURNING *`

	result := ProductVariant{}
	rows, err := db.NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateProductVariant(data ProductVariant) (ProductVariant, error) {
	sql := `
	UPDATE "productVariant" SET
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"additionalPrice"=COALESCE(:additionalPrice,"additionalPrice"),
	"updatedAt"=NOW()
	WHERE id = :id
	RETURNING *
	`
	result := ProductVariant{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteProductVariant(id int) (ProductVariant, error) {
	sql := `DELETE FROM "productVariant" WHERE "id" = $1 RETURNING *`
	data := ProductVariant{}
	err := db.Get(&data, sql, id)
	return data, err
}
