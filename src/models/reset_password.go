package models

import (
	"database/sql"
	"time"
)

type FormResetPassword struct {
	Id        int          `db:"id"`
	Email     string       `db:"email" form:"email"`
	Otp       string       `db:"otp" form:"otp"`
	CreatedAt time.Time    `db:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt"`
}

func FindOneRPByOtp(otp string) (FormResetPassword, error) {
	sql := `SELECT * FROM "resetPassword" WHERE "otp" = $1`
	data := FormResetPassword{}
	err := db.Get(&data, sql, otp)
	return data, err
}

func CreateResetPassword(data FormResetPassword) (FormResetPassword, error) {
	sql := `
	INSERT INTO "resetPassword" ("email","otp") VALUES
	(:email, :otp)
	RETURNING *`

	result := FormResetPassword{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteResetPassword(id int) (User, error) {
	sql := `DELETE FROM "resetPassword" WHERE "id" = $1 RETURNING *`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}
