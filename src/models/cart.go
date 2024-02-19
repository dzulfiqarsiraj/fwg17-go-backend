package models

import (
	"fmt"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type Cart struct {
	Id              int        `db:"id" json:"id"`
	OrderDetailId   int        `db:"orderDetailId" json:"orderDetailId" form:"orderDetailId"`
	ProductName     *string    `db:"productName" json:"productName" form:"productName"`
	ProductSize     *string    `db:"productSize" json:"productSize" form:"productSize"`
	ProductVariant  *string    `db:"productVariant" json:"productVariant" form:"productVariant"`
	Quantity        *int       `db:"quantity" json:"quantity" form:"quantity"`
	AdditionalPrice *int       `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
	Total           *int       `db:"total" json:"total" form:"total"`
	CreatedAt       *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt       *time.Time `db:"updatedAt" json:"updatedAt"`
}

type CartInfo struct {
	TotalPrice int
}

func FindAllCarts(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "cart" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "cart"`
	result := services.Info{}
	data := []Cart{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneCart(id int) (Cart, error) {
	sql := `SELECT * FROM "cart" WHERE "id"=$1`
	data := Cart{}
	err := db.Get(&data, sql, id)
	return data, err
}

func TotalPrice() (CartInfo, error) {
	sql := `SELECT SUM("total") FROM "cart"`
	result := CartInfo{}
	row := db.QueryRow(sql)
	err := row.Scan(&result.TotalPrice)
	return result, err
}

func CreateCart(data Cart) (Cart, error) {
	sql := `
	INSERT INTO "cart" ("orderDetailId","productName","productSize","productVariant","quantity","additionalPrice","total") VALUES
	(:orderDetailId, :productName,:productSize,:productVariant,:quantity,:additionalPrice,:total)
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
	"orderDetailId"=:orderDetailId
	"quantity"=COALESCE(NULLIF(:quantity,0),"quantity"),
	"updatedAt"=NOW()
	WHERE id = :id
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

func DeleteCart(id int) (Cart, error) {
	sql := `DELETE FROM "cart" WHERE "id" = $1 RETURNING *`
	data := Cart{}
	err := db.Get(&data, sql, id)
	return data, err
}

func DeleteAllCart() (Cart, error) {
	sql := `DELETE FROM "cart" RETURNING *`
	data := Cart{}
	err := db.Get(&data, sql)
	return data, err
}
