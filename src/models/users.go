package models

import (
	"fmt"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/lib"
	"github.com/DzulfiqarSiraj/go-backend/src/services"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB = lib.DB

type User struct {
	Id          int        `db:"id" json:"id"`
	Email       string     `db:"email" json:"email" form:"email"`
	Password    string     `db:"password" json:"password" form:"password"`
	FullName    *string    `db:"fullName" json:"fullName" form:"fullName"`
	PhoneNumber *string    `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Address     *string    `db:"address" json:"address" form:"address"`
	Role        string     `db:"role" json:"role" form:"role"`
	Picture     *string    `db:"picture" json:"picture"`
	CreatedAt   *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllUsers(search string, orderBy string, limit int, offset int) (services.Info, error) {
	var sql string
	var sqlCount string
	if search == "" {
		sql = `SELECT * FROM "users"
		ORDER BY "` + orderBy + `" ASC
		LIMIT $1
		OFFSET $2`
		sqlCount = `SELECT COUNT(*) FROM "users"`
		result := services.Info{}
		data := []User{}
		db.Select(&data, sql, limit, offset)
		result.Data = data

		row := db.QueryRow(sqlCount)
		err := row.Scan(&result.Count)

		return result, err
	} else {
		sql = `SELECT * FROM "users"
		WHERE "fullName" ILIKE $1
		ORDER BY "` + orderBy + `" ASC
		LIMIT $2
		OFFSET $3`
		sqlCount = `SELECT COUNT(*) FROM "users" WHERE "fullName" ILIKE $1`
		fmtSearch := fmt.Sprintf("%%%v%%", search)
		result := services.Info{}
		data := []User{}
		db.Select(&data, sql, fmtSearch, limit, offset)
		result.Data = data

		row := db.QueryRow(sqlCount, fmtSearch)
		err := row.Scan(&result.Count)

		return result, err
	}
}

func FindOneUser(id int) (User, error) {
	sql := `SELECT * FROM "users" WHERE "id" = $1`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneUserByEmail(email string) (User, error) {
	sql := `SELECT * FROM "users" WHERE "email" = $1`
	data := User{}
	err := db.Get(&data, sql, email)
	return data, err
}

func CreateUser(data User) (User, error) {
	sql := `
	INSERT INTO "users" ("email","password","fullName","phoneNumber","address","role","picture") VALUES
	(:email, :password, :fullName, :phoneNumber, :address, COALESCE(:role,'Customer'), :picture)
	RETURNING *`

	result := User{}
	rows, err := db.NamedQuery(sql, data)
	fmt.Println(rows)
	fmt.Println(err)

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateUser(data User) (User, error) {
	sql := `
	UPDATE "users" SET
	"email"=COALESCE(NULLIF(:email,''), "email"),
	"password"=COALESCE(NULLIF(:password,''), "password"),
	"fullName"=COALESCE(NULLIF(:fullName,''), "fullName"),
	"phoneNumber"=COALESCE(NULLIF(:phoneNumber,''), "phoneNumber"),
	"address"=COALESCE(NULLIF(:address,''), "address"),
	"role"=COALESCE(NULLIF(:role,''),"role"),
	"picture"=COALESCE(NULLIF(:picture,''), "picture"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := User{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteUser(id int) (User, error) {
	sql := `DELETE FROM "users" WHERE "id" = $1 RETURNING *`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}
