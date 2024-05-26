package models

import (
	"fmt"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type OrderDetail struct {
	Id               int        `db:"id" json:"id"`
	UserId           *int       `db:"userId" json:"userId" form:"userId"`
	OrderId          *int       `db:"orderId" json:"orderId" form:"orderId"`
	ProductId        *int       `db:"productId" json:"productId" form:"productId"`
	ProductSizeId    *int       `db:"productSizeId" json:"productSizeId" form:"productSizeId"`
	ProductVariantId *int       `db:"productVariantId" json:"productVariantId" form:"productVariantId"`
	Quantity         *int       `db:"quantity" json:"quantity" form:"quantity"`
	SubTotal         *float64   `db:"subTotal" json:"subTotal" form:"subTotal"`
	CreatedAt        *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt        *time.Time `db:"updatedAt" json:"updatedAt"`
}

type OrderDetailFull struct {
	Id               *int       `db:"id" json:"id"`
	OrderId          *int       `db:"orderId" json:"orderId"`
	UserId           *int       `db:"userId" json:"userId"`
	OrderNumber      *string    `db:"orderNumber" json:"orderNumber"`
	FullName         *string    `db:"fullName" json:"fullName" `
	Email            *string    `db:"email" json:"email"`
	PromoId          *int       `db:"promoId" json:"promoId"`
	Tax              *float64   `db:"tax" json:"tax"`
	GrandTotal       *float64   `db:"grandTotal" json:"grandTotal"`
	DeliveryAddress  *string    `db:"deliveryAddress" json:"deliveryAddress"`
	Status           *string    `db:"status" json:"status"`
	ProductId        *int       `db:"productId" json:"productId"`
	Product          *string    `db:"product" json:"product"`
	Image            *string    `db:"image" json:"image"`
	ProductSizeId    *int       `db:"productSizeId" json:"productSizeId"`
	Size             *string    `db:"size" json:"size"`
	ProductVariantId *int       `db:"productVariantId" json:"productVariantId"`
	Variant          *string    `db:"variant" json:"variant"`
	Quantity         *int       `db:"quantity" json:"quantity"`
	Tag              *string    `db:"tag" json:"tag"`
	Discount         *float64   `db:"discount" json:"discount"`
	SubTotal         *float64   `db:"subTotal" json:"subTotal"`
	Shipping         *string    `db:"shipping" json:"shipping"`
	Date             *string    `db:"date" json:"date"`
	Time             *string    `db:"time" json:"time"`
	CreatedAt        *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt        *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllOrderDetails(orderId int, limit int, offset int) (services.Info, error) {
	fmt.Println(orderId)
	sql := `SELECT
	"od"."id" "id",
	"o"."id" "orderId",
	"o"."userId" "userId",
	"o"."orderNumber" "orderNumber",
	"o"."fullName" "fullName",
	"o"."email" "email",
	"o"."promoId" "promoId",
	"o"."tax" "tax",
	"o"."grandTotal" "grandTotal",
	"o"."deliveryAddress" "deliveryAddress",
	"o"."status" "status",
	"od"."productId" "productId",
	"p"."name" "product",
	"p"."image" "image",
	"od"."productSizeId" "productSizeId",
	"s"."name" "size",
	"od"."productVariantId" "productVariantId",
	"v"."name" "variant",
	"od"."quantity" "quantity",
	"t"."name" "tag",
	"t"."discount",
	"od"."subTotal" "subTotal",
	"o"."shipping" "shipping",
	to_char(date("o"."createdAt"), 'YYYY-MM-DD') AS "date",
	to_char("o"."createdAt", 'HH12:MI AM') AS "time",
	"od"."createdAt",
	"od"."updatedAt"
	FROM "orderDetails" "od"
	JOIN "orders" "o" ON "o"."id"="od"."orderId"
	JOIN "products" "p" ON "p"."id"="od"."productId"
	JOIN "sizes" "s" ON "s"."id"="od"."productSizeId"
	JOIN "variants" "v" ON "v"."id"="od"."productVariantId"
	FULL JOIN "productTags" "pt" ON "pt"."productId"="p"."id"
	FULL JOIN "tags" "t" ON "t"."id"="pt"."tagId"
	WHERE "orderId" = $1
	ORDER BY "o"."id" ASC
	LIMIT $2
	OFFSET $3`

	sqlCount := `SELECT COUNT(*) as "counts" FROM "orderDetails"
	WHERE "orderId" = $1`

	result := services.Info{}
	data := []OrderDetailFull{}
	db.Select(&data, sql, orderId, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount, orderId)
	err := row.Scan(&result.Count)
	fmt.Println(result)
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
	INSERT INTO "orderDetails"("userId","orderId","productId","productSizeId","productVariantId","quantity","subTotal") VALUES
	(:userId, :orderId, :productId, :productSizeId, :productVariantId, :quantity, :subTotal)
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
