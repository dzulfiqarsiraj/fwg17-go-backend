package models

import (
	"fmt"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/lib/pq"
)

type Product struct {
	Id           int        `db:"id" json:"id"`
	Name         *string    `db:"name" json:"name" form:"name"`
	BasePrice    *int       `db:"basePrice" json:"basePrice" form:"basePrice"`
	Description  *string    `db:"description" json:"description" form:"description"`
	Image        *string    `db:"image" json:"image" form:"image"`
	IsBestSeller *bool      `db:"isBestSeller" json:"isBestSeller" form:"isBestSeller"`
	Discount     *float64   `db:"discount" json:"discount" form:"discount"`
	CreatedAt    *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt    *time.Time `db:"updatedAt" json:"updatedAt"`
}

type ProductWithCategories struct {
	Id           int             `db:"id" json:"id"`
	Name         *string         `db:"name" json:"name" form:"name"`
	Category     *pq.StringArray `db:"category" json:"category" form:"category"`
	BasePrice    *int            `db:"basePrice" json:"basePrice" form:"basePrice"`
	Description  *string         `db:"description" json:"description" form:"description"`
	Image        *string         `db:"image" json:"image" form:"image"`
	IsBestSeller *bool           `db:"isBestSeller" json:"isBestSeller" form:"isBestSeller"`
	Discount     *float64        `db:"discount" json:"discount" form:"discount"`
	CreatedAt    *time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt    *time.Time      `db:"updatedAt" json:"updatedAt"`
}

func FindAllProducts(category string, search string, orderBy string, limit int, offset int) (services.Info, error) {
	sql := `SELECT
		"p"."id" "id",
		"p"."name" "name",
		array_agg(DISTINCT COALESCE("c"."name", '')) "category",
		"p"."basePrice" "basePrice",
		"p"."description" "description",
		"p"."image" "image",
		"p"."isBestSeller" "isBestSeller",
		"p"."discount" "discount",
		"p"."createdAt" "createdAt",
		"p"."updatedAt" "updatedAt"
	FROM "products" "p"
	LEFT JOIN "productCategories" "pc" ON "pc"."productId"="p"."id"
	LEFT JOIN "categories" "c" ON "c"."id"="pc"."categoryId"
	WHERE "c"."name" ILIKE $1
	GROUP BY "p"."id"
	HAVING "p"."name" ILIKE $2
	ORDER BY "p"."` + orderBy + `" ASC
	LIMIT $3
	OFFSET $4`
	sqlCount := `SELECT COUNT(*) FROM 
	(SELECT
		"p"."id" "id",
		"p"."name" "name",
		array_agg(DISTINCT "c"."name") "category",
		"p"."basePrice" "basePrice",
		"p"."description" "description",
		"p"."image" "image",
		"p"."isBestSeller" "isBestSeller",
		"p"."discount" "discount",
		"p"."createdAt" "createdAt",
		"p"."updatedAt" "updatedAt"
	FROM "products" "p"
	LEFT JOIN "productCategories" "pc" ON "pc"."productId"="p"."id"
	LEFT JOIN "categories" "c" ON "c"."id"="pc"."categoryId"
	WHERE "c"."name" ILIKE $1
	GROUP BY "p"."id"
	HAVING "p"."name" ILIKE $2)`
	fmtSearch := fmt.Sprintf("%%%v%%", search)
	fmtCategory := fmt.Sprintf("%%%v%%", category)
	result := services.Info{}
	data := []ProductWithCategories{}
	db.Select(&data, sql, fmtCategory, fmtSearch, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount, fmtCategory, fmtSearch)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneProduct(id int) (Product, error) {
	sql := `SELECT * FROM "products" WHERE "id"=$1`
	data := Product{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneProductByName(name string) (Product, error) {
	sql := `SELECT * FROM "products" WHERE "name" = $1`
	data := Product{}
	err := db.Get(&data, sql, name)
	return data, err
}

func FindProductNameById(id int) (Product, error) {
	sql := `SELECT "name" FROM "products" WHERE "id" = $1`
	data := Product{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProduct(data Product) (Product, error) {
	sql := `
	INSERT INTO "products" ("name","basePrice","description","image","discount","isBestSeller") VALUES
	(:name, :basePrice, :description, :image, :discount, :isBestSeller)
	RETURNING *`

	result := Product{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateProduct(data Product) (Product, error) {
	sql := `
	UPDATE "products" SET
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"basePrice"=COALESCE(NULLIF(:basePrice,0),"basePrice"),
	"description"=COALESCE(NULLIF(:description,''),"description"),
	"image"=COALESCE(NULLIF(:image,''),"image"),
	"isBestSeller"=COALESCE(:isBestSeller,false),
	"discount"=COALESCE(NULLIF(:discount,0.0),"discount"),
	"updatedAt"=NOW()
	WHERE id = :id
	RETURNING *
	`
	result := Product{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteProduct(id int) (Product, error) {
	sql := `DELETE FROM "products" WHERE "id" = $1 RETURNING *`
	data := Product{}
	err := db.Get(&data, sql, id)
	return data, err
}
