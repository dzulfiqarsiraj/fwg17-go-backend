package models

import (
	"fmt"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type ProductCategory struct {
	Id         int        `db:"id" json:"id"`
	ProductId  *int       `db:"productId" json:"productId" form:"productId"`
	CategoryId *int       `db:"categoryId" json:"categoryId" form:"categoryId"`
	CreatedAt  *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt  *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllProductCategories(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "productCategories" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productCategories"`
	result := services.Info{}
	data := []ProductCategory{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneProductCategory(id int) (ProductCategory, error) {
	sql := `SELECT * FROM "productCategories" WHERE "id"=$1`
	data := ProductCategory{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneProductCategoryByProductId(productId int) (ProductCategory, error) {
	sql := `SELECT * FROM "categories" WHERE "productId"=$1`
	data := ProductCategory{}
	err := db.Get(&data, sql, productId)
	return data, err
}

func CreateProductCategory(data ProductCategory) (ProductCategory, error) {
	sql := `
	INSERT INTO "productCategories" ("productId","categoryId") VALUES
	(:productId,:categoryId)
	RETURNING *`

	result := ProductCategory{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateProductCategory(data ProductCategory) (ProductCategory, error) {
	sql := `
	UPDATE "productCategories" SET
	"productId"=COALESCE(NULLIF(:productId,''),"productId"),
	"categoryId"=COALESCE(NULLIF(:categoryId,''),"categoryId"),
	"updatedAt"=NOW()
	WHERE id = :id
	RETURNING *
	`
	result := ProductCategory{}
	rows, err := db.NamedQuery(sql, data)
	fmt.Println(sql)
	fmt.Println(rows)
	fmt.Println(err)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteProductCategory(id int) (ProductCategory, error) {
	sql := `DELETE FROM "productCategories" WHERE "id" = $1 RETURNING *`
	data := ProductCategory{}
	err := db.Get(&data, sql, id)
	return data, err
}
