package models

import (
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type ProductRating struct {
	Id            int        `db:"id" json:"id"`
	ProductId     *int       `db:"productId" json:"productId" form:"productId"`
	Rate          *int       `db:"rate" json:"rate" form:"rate"`
	ReviewMessage *int       `db:"reviewMessage" json:"reviewMessage" form:"reviewMessage"`
	CreatedAt     *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt     *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllProductRatings(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "productRatings"
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productRatings"`
	result := services.Info{}
	data := []ProductRating{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneProductRating(id int) (ProductRating, error) {
	sql := `SELECT * FROM "productRatings" WHERE "id"=$1`
	data := ProductRating{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneProductRatingByProductId(productId int) (ProductRating, error) {
	sql := `SELECT * FROM "productRatings" WHERE "productId" = $1`
	data := ProductRating{}
	err := db.Get(&data, sql, productId)
	return data, err
}

func CreateProductRating(data ProductRating) (ProductRating, error) {
	sql := `
	INSERT INTO "productRatings" ("productId","rate","reviewMessage") VALUES
	(:productId, :rate, :reviewMessage)
	RETURNING *`

	result := ProductRating{}
	rows, err := db.NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateProductRating(data ProductRating) (ProductRating, error) {
	sql := `
	UPDATE "productRatings" SET
	"productId"=COALESCE(NULLIF(:productId,0),"productId"),
	"rate"=COALESCE(NULLIF(:rate,0),"rate"),
	"reviewMessage"=COALESCE(NULLIF(:reviewMessage,''),"reviewMessage"),
	"updatedAt"=NOW()
	WHERE id = :id
	RETURNING *
	`
	result := ProductRating{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteProductRating(id int) (ProductRating, error) {
	sql := `DELETE FROM "productRatings" WHERE "id" = $1 RETURNING *`
	data := ProductRating{}
	err := db.Get(&data, sql, id)
	return data, err
}
