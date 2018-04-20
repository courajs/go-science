package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type UserVerifier struct {
	id                int
	email             string
	username          string
	encryptedPassword string
}

func main() {
	db, err := sql.Open("postgres", "postgres://focker:@localhost:5432/mlnptv_www_production?sslmode=disable")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT id, email, username, encrypted_password FROM users LIMIT 10")
	if err != nil {
		panic(err)
	}

	colTypes, err := rows.ColumnTypes()
	if err != nil {
		panic(err)
	}
	typeTypes := make([]string, len(colTypes))
	for i, coltype := range colTypes {
		typeTypes[i] = coltype.Name() + ":" + coltype.DatabaseTypeName()
	}
	fmt.Println(typeTypes)
	var uv UserVerifier
	for rows.Next() {
		rows.Scan(&uv.id, &uv.email, &uv.username, &uv.encryptedPassword)
		fmt.Println(uv)
	}
}
