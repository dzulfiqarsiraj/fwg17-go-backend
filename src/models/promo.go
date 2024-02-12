package models

import (
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type Promo struct {
	Id          int        `db:"id" json:"id"`
	Name        *string    `db:"name" json:"name" form:"name"`
	Code        *string    `db:"code" json:"code" form:"code"`
	Description *string    `db:"description" json:"description" form:"description"`
	Percentage  *float64   `db:"percentage" json:"percentage" form:"percentage"`
	MaxPromo    *int       `db:"maxPromo" json:"maxPromo" form:"maxPromo"`
	MinPurchase *int       `db:"minPurchase" json:"minPurchase" form:"minPurchase"`
	IsExpired   *bool      `db:"isExpired" json:"isExpired" form:"isExpired"`
	CreatedAt   *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllPromo(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "promo"
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "promo"`
	result := services.Info{}
	data := []Promo{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOnePromo(id int) (Promo, error) {
	sql := `SELECT * FROM "promo" WHERE "id"=$1`
	data := Promo{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOnePromoByName(name string) (Promo, error) {
	sql := `SELECT * FROM "promo" WHERE "name" = $1`
	data := Promo{}
	err := db.Get(&data, sql, name)
	return data, err
}

func CreatePromo(data Promo) (Promo, error) {
	sql := `
	INSERT INTO "promo" ("name","code","description","percentage","maxPromo","minPurchase","isExpired") VALUES
	(:name, :code, :description, :percentage, :maxPromo, :minPurchase, :isExpired)
	RETURNING *`

	result := Promo{}
	rows, err := db.NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdatePromo(data Promo) (Promo, error) {
	sql := `
	UPDATE "promo" SET
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"code"=COALESCE(NULLIF(:code,''),"code"),
	"description"=COALESCE(NULLIF(:description,''),"description"),
	"percentage"=COALESCE(NULLIF(:percentage,0.0),"percentage"),
	"maxPromo"=COALESCE(NULLIF(:maxPromo,0),"maxPromo"),
	"minPurchase"=COALESCE(NULLIF(:minPurchase,0),"minPurchase"),
	"isExpired"=COALESCE(NULLIF(:isExpired,false),"isExpired"),
	"updatedAt"=NOW()
	WHERE id = :id
	RETURNING *
	`
	result := Promo{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeletePromo(id int) (Promo, error) {
	sql := `DELETE FROM "promo" WHERE "id" = $1 RETURNING *`
	data := Promo{}
	err := db.Get(&data, sql, id)
	return data, err
}
