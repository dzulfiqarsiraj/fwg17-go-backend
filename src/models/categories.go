package models

import (
	"fmt"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type Category struct {
	Id        int        `db:"id" json:"id"`
	Name      *string    `db:"name" json:"name" form:"name"`
	CreatedAt *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllCategories(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "categories" 
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "categories"`
	result := services.Info{}
	data := []Category{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneCategory(id int) (Category, error) {
	sql := `SELECT * FROM "categories" WHERE "id"=$1`
	data := Category{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneCategoryByName(name string) (Category, error) {
	sql := `SELECT * FROM "categories" WHERE "name"=$1`
	data := Category{}
	err := db.Get(&data, sql, name)
	return data, err
}

func CreateCategory(data Category) (Category, error) {
	sql := `
	INSERT INTO "categories" ("name") VALUES
	(:name)
	RETURNING *`

	result := Category{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateCategory(data Category) (Category, error) {
	sql := `
	UPDATE "categories" SET
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"updatedAt"=NOW()
	WHERE id = :id
	RETURNING *
	`
	result := Category{}
	rows, err := db.NamedQuery(sql, data)
	fmt.Println(sql)
	fmt.Println(rows)
	fmt.Println(err)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteCategory(id int) (Category, error) {
	sql := `DELETE FROM "categories" WHERE "id" = $1 RETURNING *`
	data := Category{}
	err := db.Get(&data, sql, id)
	return data, err
}
