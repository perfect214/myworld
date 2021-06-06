package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"runtime"
)

func main() {
	db, err := sql.Open("mysql","root@tcp(127.0.0.1:3306)/tesst?charset=utf8")
    if err != nil {
    	fmt.Println(errors.Wrap(err, "Cannot operation the db"))
	}
	defer db.close()

	type User struct {
		Id int32
		Name string
		Age int8
	}

	var user User

	rows, err := db.Query(`SELECT id, name, age From user`)

	if err != nil {
		fmt.Println(errors.Wrap(err,"DML operation failed"))
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			fmt.Println(errors.Wrap(err, "DML operation failed"))
			continue
		}
		fmt.Println(user.Id, user.Name, user.Age)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(errors.Wrap(err, "row.Err"))
	}

	err = db.QueryRow(`SELECT id, name, age WHERE id = ?`,2).scan(
		&user.Id, &user.Name, &user.Age,
		)
	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		if _, file, line, ok := runtime.Caller(0); ok {
			fmt.Println(errors.WithMessage(err, "sql.ErrNoRows"), file,line)
		}
	}
	fmt.Println(user.Id, user.Name, user.Age)
}