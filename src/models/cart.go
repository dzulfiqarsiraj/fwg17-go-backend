package models

import (
	"fmt"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type Cart struct {
	Id               int        `db:"id" json:"id"`
	ProductId        *int       `db:"productId" json:"productId" form:"productId"`
	ProductSizeId    *int       `db:"productSizeId" json:"productSizeId" form:"productSizeId"`
	ProductVariantId *int       `db:"productVariantId" json:"productVariantId" form:"productVariantId"`
	Quantity         *int       `db:"quantity" json:"quantity" form:"quantity"`
	Total            *float64   `db:"total" json:"total" form:"total"`
	UserId           *int       `db:"userId" json:"userId" form:"userId"`
	CreatedAt        *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt        *time.Time `db:"updatedAt" json:"updatedAt"`
}

type CartInfo struct {
	TotalPrice int
}

func FindAllCarts(userId int, limit int, offset int) (services.Info, error) {
	fmtUserId := fmt.Sprintf(`%v`, userId)
	sql := `SELECT * FROM "cart" 
	WHERE "userId" = ` + fmtUserId + `
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "cart" WHERE "userId" = ` + fmtUserId
	result := services.Info{}
	data := []Cart{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneCart(id int, userId int) (Cart, error) {
	sql := `SELECT * FROM "cart" WHERE "id" = $1 AND "userId" = $2`
	data := Cart{}
	err := db.Get(&data, sql, id, userId)
	return data, err
}

func TotalPrice(userId int) (CartInfo, error) {
	sql := `SELECT SUM("total") FROM "cart" 
	WHERE "userId" = $1`
	result := CartInfo{}
	row := db.QueryRow(sql, userId)
	err := row.Scan(&result.TotalPrice)
	return result, err
}

func CreateCart(data Cart) (Cart, error) {
	sql := `
	INSERT INTO "cart" ("productId","productSizeId","productVariantId","quantity","total","userId") VALUES
	(:productId,:productSizeId,:productVariantId,:quantity,:total,:userId)
	RETURNING *`

	result := Cart{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateCart(data Cart) (Cart, error) {
	sql := `
	UPDATE "cart" SET
	"productSizeId" = COALESCE(NULLIF(:productSizeId,0), "productSizeId"),
	"productVariantId" = COALESCE(NULLIF(:productVariantId,0), "productVariantId"),
	"quantity" = COALESCE(NULLIF(:quantity,0), "quantity"),
	"total" = :total,
	"updatedAt"=NOW()
	WHERE "id" = :id AND "userId" = :userId
	RETURNING *
	`
	result := Cart{}
	rows, err := db.NamedQuery(sql, data)
	fmt.Println(sql)
	fmt.Println(rows)
	fmt.Println(err)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteCart(id int, userId int) (Cart, error) {
	sql := `DELETE FROM "cart" 
	WHERE "id" = $1 AND "userId" = $2
	RETURNING *`
	data := Cart{}
	err := db.Get(&data, sql, id, userId)
	return data, err
}

func DeleteAllCart(userId int) (Cart, error) {
	sql := `DELETE FROM "cart" WHERE "userId" = $1 RETURNING *`
	data := Cart{}
	err := db.Get(&data, sql, userId)
	return data, err
}
