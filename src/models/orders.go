package models

import (
	"fmt"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type Order struct {
	Id              int        `db:"id" json:"id"`
	UserId          *int       `db:"userId" json:"userId" form:"userId"`
	OrderNumber     *string    `db:"orderNumber" json:"orderNumber" form:"orderNumber"`
	FullName        *string    `db:"fullName" json:"fullName" form:"fullName"`
	Email           *string    `db:"email" json:"email" form:"email"`
	PromoId         *int       `db:"promoId" json:"promoId" form:"promoId"`
	Tax             *float64   `db:"tax" json:"tax" form:"tax"`
	GrandTotal      *float64   `db:"grandTotal" json:"grandTotal" form:"grandTotal"`
	DeliveryAddress *string    `db:"deliveryAddress" json:"deliveryAddress" form:"deliveryAddress"`
	Status          *string    `db:"status" json:"status" form:"status"`
	Shipping        *string    `db:"shipping" json:"shipping" form:"shipping"`
	CreatedAt       *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt       *time.Time `db:"updatedAt" json:"updatedAt"`
}

type MaxId struct {
	Max *int
}

func FindAllOrders(userId int, status string, limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "orders"
	WHERE "userId" = $1 AND "status" = $2
	ORDER BY "id" DESC
	LIMIT $3
	OFFSET $4`
	sqlCount := `SELECT COUNT(*) FROM "orders" WHERE "userId" = $1  AND "status" = $2`
	result := services.Info{}
	data := []Order{}
	db.Select(&data, sql, userId, status, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount, userId, status)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneOrder(id int) (Order, error) {
	sql := `SELECT * FROM "orders" WHERE "id"=$1`
	data := Order{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindMaxIdOrder() (MaxId, error) {
	sql := `SELECT MAX("id") FROM "orders"`
	result := MaxId{}
	row := db.QueryRow(sql)
	err := row.Scan(&result.Max)
	return result, err
}

func FindOneOrderByOrderNumber(orderNumber string) (Order, error) {
	sql := `SELECT * FROM "orders" WHERE "orderNumber"=$1`
	data := Order{}
	err := db.Get(&data, sql, orderNumber)
	return data, err
}

func CreateOrder(data Order) (Order, error) {
	sql := `
	INSERT INTO "orders" ("userId","orderNumber","fullName","email","promoId","tax","grandTotal","deliveryAddress","status","shipping") VALUES
	(:userId,:orderNumber,:fullName,:email,:promoId,:tax,:grandTotal,:deliveryAddress,COALESCE(:status,'Awaiting Payment'),:shipping)
	RETURNING *`

	result := Order{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateOrder(data Order) (Order, error) {
	sql := `
	UPDATE "orders" SET
	"userId"=COALESCE(NULLIF(:userId,0),"userId"),
	"orderNumber"=COALESCE(NULLIF(:orderNumber,''),"orderNumber"),
	"fullName"=COALESCE(NULLIF(:fullName,''),"fullName"),
	"email"=COALESCE(NULLIF(:email,''),"email"),
	"promoId"=COALESCE(NULLIF(:promoId,0),"promoId"),
	"tax"=COALESCE(NULLIF(:tax,0),"tax"),
	"total"=COALESCE(NULLIF(:total,0),"total"),
	"deliveryAddress"=COALESCE(NULLIF(:deliveryAddress,''),"deliveryAddress"),
	"status"=COALESCE(NULLIF(:status,''),"status"),
	"updatedAt"=NOW()
	WHERE id = :id
	RETURNING *
	`
	result := Order{}
	rows, err := db.NamedQuery(sql, data)
	fmt.Println(sql)
	fmt.Println(rows)
	fmt.Println(err)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateOrderNumber(data Order) (Order, error) {
	sql := `UPDATE "orders" SET 
	"orderNumber" = :orderNumber
	WHERE "id" = :id
	RETURNING *`
	result := Order{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteOrder(id int) (Order, error) {
	sql := `DELETE FROM "orders" WHERE "id" = $1 RETURNING *`
	data := Order{}
	err := db.Get(&data, sql, id)
	return data, err
}
