package models

import (
	"time"

	"github.com/DzulfiqarSiraj/go-backend/src/lib"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB = lib.DB

type User struct {
	Id          int         `db:"id" json:"id"`
	Email       string      `db:"email" json:"email" form:"email" binding:"email"`
	Password    string      `db:"password" json:"password" form:"password"`
	FullName    interface{} `db:"fullName" json:"fullName" form:"default:null"`
	PhoneNumber interface{} `db:"phoneNumber" json:"phoneNumber" form:"default:null"`
	Address     interface{} `db:"address" json:"address" form:"default:null"`
	Role        *string     `db:"role" json:"role" form:"default:Customer"`
	Picture     interface{} `db:"picture" json:"picture" form:"default:null"`
	CreatedAt   *time.Time  `db:"createdAt" json:"createdAt"`
	UpdatedAt   *time.Time  `db:"updatedAt" json:"updatedAt"`
}

func FindAllUsers() ([]User, error) {
	sql := `SELECT * FROM "users"`
	data := []User{}
	err := db.Select(&data, sql)
	return data, err
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
	INSERT INTO "users" ("email","password") VALUES
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
