package models

import (
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/services"
)

type Tag struct {
	Id        int        `db:"id" json:"id"`
	Name      *string    `db:"name" json:"name" form:"name"`
	CreatedAt *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllTags(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "tags"
	ORDER BY "id" ASC
	LIMIT $1
	OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "tags"`
	result := services.Info{}
	data := []Tag{}
	db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err := row.Scan(&result.Count)
	return result, err
}

func FindOneTag(id int) (Tag, error) {
	sql := `SELECT * FROM "tags" WHERE "id"=$1`
	data := Tag{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneTagByName(name string) (Tag, error) {
	sql := `SELECT * FROM "tags" WHERE "name" = $1`
	data := Tag{}
	err := db.Get(&data, sql, name)
	return data, err
}

func CreateTag(data Tag) (Tag, error) {
	sql := `
	INSERT INTO "tags" ("name") VALUES
	(:name)
	RETURNING *`

	result := Tag{}
	rows, err := db.NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateTag(data Tag) (Tag, error) {
	sql := `
	UPDATE "tags" SET
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"updatedAt"=NOW()
	WHERE id = :id
	RETURNING *
	`
	result := Tag{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteTag(id int) (Tag, error) {
	sql := `DELETE FROM "tags" WHERE "id" = $1 RETURNING *`
	data := Tag{}
	err := db.Get(&data, sql, id)
	return data, err
}
