package models

import (
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type ProductTag struct {
	Id        int        `db:"id" json:"id"`
	TagId     *int       `db:"tagId" json:"tagId" form:"tagId"`
	ProductId *int       `db:"productId" json:"productId" form:"productId"`
	CreatedAt *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllProductTags(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "productTags"
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productTags"`
	result := services.Info{}
	data := []ProductTag{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneProductTag(id int) (ProductTag, error) {
	sql := `SELECT * FROM "productTags" WHERE "id"=$1`
	data := ProductTag{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneTagByTagId(tagId int) (ProductTag, error) {
	sql := `SELECT * FROM "productTags" WHERE "tagId" = $1`
	data := ProductTag{}
	err := db.Get(&data, sql, tagId)
	return data, err
}

func CreateProductTag(data ProductTag) (ProductTag, error) {
	sql := `
	INSERT INTO "productTags" ("name","tagId","productId") VALUES
	(:name, :tagId, :productId)
	RETURNING *`

	result := ProductTag{}
	rows, err := db.NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateProductTag(data ProductTag) (ProductTag, error) {
	sql := `
	UPDATE "productTags" SET
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"tagId"=COALESCE(NULLIF(:tagId,0),"tagId"),
	"productId"=COALESCE(NULLIF(:productId,0),"productId"),
	"updatedAt"=NOW()
	WHERE id = :id
	RETURNING *
	`
	result := ProductTag{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteProductTag(id int) (ProductTag, error) {
	sql := `DELETE FROM "productTags" WHERE "id" = $1 RETURNING *`
	data := ProductTag{}
	err := db.Get(&data, sql, id)
	return data, err
}
