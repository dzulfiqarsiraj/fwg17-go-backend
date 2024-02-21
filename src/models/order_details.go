package models

import (
	"fmt"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type OrderDetail struct {
	Id               int        `db:"id" json:"id"`
	ProductId        *int       `db:"productId" json:"productId" form:"productId"`
	ProductSizeId    *int       `db:"productSizeId" json:"productSizeId" form:"productSizeId"`
	ProductVariantId *int       `db:"productVariantId" json:"productVariantId" form:"productVariantId"`
	Quantity         *int       `db:"quantity" json:"quantity" form:"quantity"`
	OrderId          *int       `db:"orderId" json:"orderId" form:"orderId"`
	UserId           *int       `db:"userId" json:"userId" form:"userId"`
	CreatedAt        *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt        *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllOrderDetails(userId int, limit int, offset int) (services.Info, error) {
	fmtUserId := fmt.Sprintf(`%v`, userId)
	sql := `SELECT * 
	FROM "orderDetails"
	WHERE "userId" = ` + fmtUserId + ` AND "orderId" IS NOT NULL
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) 
	FROM "orderDetails"
	WHERE "userId" = ` + fmtUserId + ` AND "orderId" IS NOT NULL`

	result := services.Info{}
	data := []OrderDetail{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneOrderDetail(id int, userId int) (OrderDetail, error) {
	sql := `SELECT * FROM "orderDetails" WHERE "id"=$1 AND "userId"=$2`
	data := OrderDetail{}
	err := db.Get(&data, sql, id, userId)
	return data, err
}

func CreateOrderDetail(data OrderDetail) (OrderDetail, error) {
	sql := `
	INSERT INTO "orderDetails" ("productId","productSizeId","productVariantId","quantity","orderId","userId") VALUES
	(:productId, :productSizeId, :productVariantId, :quantity, :orderId, :userId)
	RETURNING *`

	result := OrderDetail{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateOrderDetail(data OrderDetail) (OrderDetail, error) {
	sql := `
	UPDATE "orderDetails" SET
	"productId"=COALESCE(NULLIF(:productId,0),"productId"),
	"productSizeId"=COALESCE(NULLIF(:productSizeId,0),"productSizeId"),
	"productVariantId"=COALESCE(NULLIF(:productVariantId,0),"productVariantId"),
	"quantity"=COALESCE(NULLIF(:quantity,0),"quantity"),
	"orderId"=COALESCE(NULLIF(:orderId,0),"orderId"),
	"updatedAt"=NOW()
	WHERE "productId" = :productId AND "userId" = :userId
	RETURNING *
	`
	result := OrderDetail{}
	rows, err := db.NamedQuery(sql, data)
	fmt.Println(sql)
	fmt.Println(rows)
	fmt.Println(err)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateOrderDetailByOrderId(userId int, data OrderDetail) (OrderDetail, error) {
	fmtUserId := fmt.Sprintf(`%v`, userId)
	sql := `
	UPDATE "orderDetails" SET
	"orderId"=COALESCE(NULLIF(:orderId,0),"orderId"),
	"updatedAt"=NOW()
	WHERE "orderId" IS NULL AND "userId" = ` + fmtUserId + `
	RETURNING *
	`
	result := OrderDetail{}
	rows, err := db.NamedQuery(sql, data)
	fmt.Println(sql)
	fmt.Println(rows)
	fmt.Println(err)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteOrderDetail(productId int, userId int) (OrderDetail, error) {
	sql := `DELETE FROM "orderDetails" 
	WHERE "productId" = $1 AND "userId" = $2
	RETURNING *`
	data := OrderDetail{}
	err := db.Get(&data, sql, productId, userId)
	return data, err
}
