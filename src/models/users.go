package models

import (
	"database/sql"
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/lib"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB = lib.DB

type User struct {
	Id          int          `db:"id" json:"id"`
	Email       string       `db:"email" json:"email" form:"email" binding:"email"`
	Password    string       `db:"password" json:"password" form:"password"`
	FullName    string       `db:"fullName" json:"fullName" form:"fullName"`
	PhoneNumber string       `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Address     string       `db:"address" json:"address" form:"address"`
	Role        string       `db:"role" json:"role" form:"role"`
	Picture     string       `db:"picture" json:"picture" form:"picture"`
	CreatedAt   time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoUser struct {
	Data  []User
	Count int
}

func FindAllUsers(limit int, offset int) (InfoUser, error) {
	sql := `SELECT * FROM "users" LIMIT $1 OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "users"`
	result := InfoUser{}
	dataUser := []User{}
	err := db.Select(&dataUser, sql, limit, offset)

	result.Data = dataUser
	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)
	return result, err
}

func FindOneUser(id int) (User, error) {
	sql := `SELECT * FROM "users" WHERE "id" = $1`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateUser(data User) (User, error) {
	sql :=
		`INSERT INTO "users" ("email", "password") VALUES 
	(:email, :password)
	RETURNING *`

	result := User{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateUser(data User) (User, error) {
	sql :=
		`UPDATE "users" SET 
	email=COALESCE(NULLIF(:email,''),email),
	password=COALESCE(NULLIF(:password,''),password),
	fullName=COALESCE(NULLIF(:fullName,''),fullName),
	phoneNumber=COALESCE(NULLIF(:phoneNumber,''),phoneNumber),
	address=COALESCE(NULLIF(:address,''),address),
	role=COALESCE(NULLIF(:role,''),role),
	picture=COALESCE(NULLIF(:picture,''),picture)
	WHERE id=:id
	RETURNING *`
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
