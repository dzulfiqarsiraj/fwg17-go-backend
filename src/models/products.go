package models

import (
	"fmt"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/lib/pq"
)

type Product struct {
	Id            int        `db:"id" json:"id"`
	Name          *string    `db:"name" json:"name" form:"name"`
	BasePrice     *int       `db:"basePrice" json:"basePrice" form:"basePrice"`
	Description   *string    `db:"description" json:"description" form:"description"`
	Image         *string    `db:"image" json:"image" form:"image"`
	IsRecommended *bool      `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
	Discount      *float64   `db:"discount" json:"discount" form:"discount"`
	CreatedAt     *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt     *time.Time `db:"updatedAt" json:"updatedAt"`
}

type ProductWithCategories struct {
	Id            int             `db:"id" json:"id"`
	Name          *string         `db:"name" json:"name" form:"name"`
	Category      *pq.StringArray `db:"category" json:"category" form:"category"`
	Description   *string         `db:"description" json:"description" form:"description"`
	BasePrice     *int            `db:"basePrice" json:"basePrice" form:"basePrice"`
	Tag           *string         `db:"tag" json:"tag" form:"tag"`
	Discount      *float64        `db:"discount" json:"discount" form:"discount"`
	Image         *string         `db:"image" json:"image" form:"image"`
	IsRecommended *bool           `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
	CreatedAt     *time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt     *time.Time      `db:"updatedAt" json:"updatedAt"`
}

type Size struct {
	ID              int    `json:"id"`
	Size            string `json:"size"`
	AdditionalPrice int    `json:"additionalPrice"`
}

type ProductDetailed struct {
	Id            int             `db:"id" json:"id"`
	Name          *string         `db:"name" json:"name" form:"name"`
	Category      *string         `db:"category" json:"category" form:"category"`
	Sizes         *pq.StringArray `db:"sizes" json:"sizes" form:"sizes"`
	Variants      *pq.StringArray `db:"variants" json:"variants" form:"variants"`
	Tag           *string         `db:"tag" json:"tag" form:"tag"`
	Description   *string         `db:"description" json:"description" form:"description"`
	BasePrice     *int            `db:"basePrice" json:"basePrice" form:"basePrice"`
	Discount      *float64        `db:"discount" json:"discount" form:"discount"`
	Image         *string         `db:"image" json:"image" form:"image"`
	IsRecommended *bool           `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
	CreatedAt     *time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt     *time.Time      `db:"updatedAt" json:"updatedAt"`
}

func FindAllProducts(category string, keyword string, orderBy string, limit int, offset int) (services.Info, error) {
	sql := `SELECT
		"p"."id" "id",
		"p"."name" "name",
		array_agg(DISTINCT COALESCE("c"."name", '')) "category",
		"p"."description" "description",
		"p"."basePrice" "basePrice",
		"t"."name" "tag",
		"t"."discount" "discount",
		"p"."image" "image",
		"p"."isRecommended" "isRecommended",
		"p"."createdAt" "createdAt",
		"p"."updatedAt" "updatedAt"
		FROM "products" "p"
		FULL JOIN "productCategories" "pc" ON "pc"."productId"="p"."id"
		FULL JOIN "productTags" "pt" ON "pt"."productId"="p"."id"
		FULL JOIN "categories" "c" ON "c"."id"="pc"."categoryId"
		FULL JOIN "tags" "t" ON "t"."id"="pt"."tagId"
		WHERE "p"."name" ILIKE $1
		GROUP BY "p"."id","t"."id"
		HAVING "p"."name" ILIKE $2
		ORDER BY "p"."` + orderBy + `" ASC
		LIMIT $3
		OFFSET $4
		`
	sqlCount := `SELECT COUNT(*) FROM 
	(SELECT
		"p"."id" "id",
		"p"."name" "name",
		array_agg(DISTINCT COALESCE("c"."name", '')) "category",
		"p"."description" "description",
		"p"."basePrice" "basePrice",
		"t"."name" "tag",
		"t"."discount" "discount",
		"p"."image" "image",
		"p"."isRecommended" "isRecommended",
		"p"."createdAt" "createdAt",
		"p"."updatedAt" "updatedAt"
		FROM "products" "p"
		FULL JOIN "productCategories" "pc" ON "pc"."productId"="p"."id"
		FULL JOIN "productTags" "pt" ON "pt"."productId"="p"."id"
		FULL JOIN "categories" "c" ON "c"."id"="pc"."categoryId"
		FULL JOIN "tags" "t" ON "t"."id"="pt"."tagId"
		WHERE "p"."name" ILIKE $1
		GROUP BY "p"."id","t"."id"
		HAVING "p"."name" ILIKE $2) AS "data"`
	fmtSearch := fmt.Sprintf("%%%v%%", keyword)
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

func FindOneProductDetailed(id int) (ProductDetailed, error) {
	sql := `
		SELECT "p"."id", 
		"p"."name",
		"c"."name" "category",
		ARRAY_AGG(DISTINCT JSONB_BUILD_OBJECT('id',"s"."id",'size',"s"."name",'additionalPrice',"s"."additionalPrice")) AS "sizes",
		ARRAY_AGG(DISTINCT JSONB_BUILD_OBJECT('id',"v"."id",'variant',"v"."name",'additionalPrice',"v"."additionalPrice")) AS "variants",
		"t"."name" "tag",
		"p"."description",
		"p"."basePrice",
		"t"."discount",
		"p"."image",
		"p"."isRecommended",
		"p"."createdAt",
		"p"."updatedAt"
		FROM "products" "p"
		FULL JOIN "productVariants" "pv" ON "pv"."productId" = "p"."id"
		FULL JOIN "productSizes" "ps" ON "ps"."productId" = "p"."id"
		FULL JOIN "productCategories" "pc" ON "pc"."productId" = "p"."id"
		FULL JOIN "productTags" "pt" ON "pt"."productId"="p"."id"
		FULL JOIN "variants" "v" ON "v"."id"="pv"."variantId"
		FULL JOIN "sizes" "s" ON "s"."id"="ps"."sizeId"
		FULL JOIN "categories" "c" ON "c"."id"="pc"."categoryId"
		FULL JOIN "tags" "t"ON "t"."id"="pt"."tagId"
		WHERE "p"."id" IS NOT NULL AND "p"."id" = $1
		GROUP BY "p"."id", "p"."name", "c"."name","t"."id"
		ORDER BY "p"."id";
	`
	data := ProductDetailed{}
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
