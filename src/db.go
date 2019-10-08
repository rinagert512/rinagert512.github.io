package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

const (
	DbInfo = "user=postgres password=1234 dbname=amazon sslmode=disable"
	CreateUsersTable = "CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, username TEXT, pass TEXT, balance INT)"
	CreateCatalogTable = "CREATE TABLE IF NOT EXISTS catalog(id SERIAL PRIMARY KEY, category TEXT, item TEXT, creator TEXT, price INT)"
)

var db *sql.DB

func check_password(login, pass string) bool {
	var db_pass string
	row := db.QueryRow("SELECT pass FROM users WHERE username = $1", login)
	row.Scan(&db_pass)

	if db_pass != pass {
		log("Incorrect password!")
		return false
	}
	return true
}

func check_login(login string) bool {
	var amount int
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", login)
	err := row.Scan(&amount)

	if err != nil {
		panic("Error accessing database")
		return false
	} 
	if amount > 0 {
		return false
	}
	return true
}

func register() {
	log("---------------------------------------------------------")
	fmt.Print("Input login: ")
	login := read_string()
	fmt.Print("Input password: ")
	password := read_string()

	if !check_login(login) {
		log("User already exists")
		return
	}

	var id int
	db.QueryRow("INSERT INTO users (username, pass, balance) values($1, $2, 0)", login, password).Scan(&id)
	log("Successfully registered!")
}

func login(client *User) {
	log("---------------------------------------------------------")
	fmt.Print("Input login: ")
	login := read_string()
	fmt.Print("Input password: ")
	password := read_string()

	var amount int
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", login)
	err := row.Scan(&amount)

	if err != nil {
		panic("Error accessing database")
		return
	} 
	if amount == 0 {
		log("No such user")
		return
	} 
	if !check_password(login, password) {
		log("Invalid password")
		return
	}

	var balance int
	db.QueryRow("SELECT balance FROM users WHERE username = $1", login).Scan(&balance)

	log("Successfully logged in!")
	(*client).login(login, password, balance)
}	

func change_username(client *UsrStruct) {
	log("---------------------------------------------------------")
	fmt.Print("Input password: ")
	password := read_string()

	if !check_password((*client).username, password) {
		log("Invalid password")
		return
	}

	fmt.Print("Input new login: ")
	new_login := read_string()

	if !check_login(new_login) {
		log("User already exists")
		return
	}

	_, err := db.Exec("UPDATE users SET username = $1 WHERE username = $2", new_login, (*client).username)
	if err != nil {
		panic(err)
	}

	log("Successfuly changed login!")
	(*client).username = new_login
	return
}

func change_password(client *UsrStruct) {
	log("---------------------------------------------------------")
	fmt.Print("Input old password: ")
	password := read_string()

	if !check_password((*client).username, password) {
		log("Invalid password")
		return
	}

	fmt.Print("Input new password: ")
	new_password := read_string()

	_, err := db.Exec("UPDATE users SET pass = $1 WHERE username = $2", new_password, (*client).username)
	if err != nil {
		panic(err)
	}

	log("Successfuly changed password!")
	(*client).password = new_password
}

func get_data_from_catdb(cat *CatStruct) bool {
	rows, err := db.Query("SELECT * FROM catalog")

	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var id, price int
			var category, item, creator string

			rows.Scan(&id, &category, &item, &creator, &price)
			cat.categories[category] = make(map[string]Item)
			cat.categories[category][item] = Item{price, creator}
		}

		return true
	} else {
		return false
	}
}

func add_to_catdb_item(category string, item string, creator string, price int) {
	if item == "" {
		item = "nil"
	} else {
		db.Exec("DELETE FROM catalog WHERE category = $1 AND item = $2", category, "nil")
	}
	var id int
	db.QueryRow("INSERT INTO catalog (category, item, creator, price) values($1, $2, $3, $4)", category, item, creator, price).Scan(&id)
}

func add_usr_balance(login string, money int) {
	_, err := db.Exec("UPDATE users SET balance = $1 WHERE username = $2", money, login)
	if err != nil {
		panic(err)
	}
}

func buy_item(item string, category string, client *UsrStruct) bool {
	var amount int
	row := db.QueryRow("SELECT COUNT(*) FROM catalog WHERE category = $1 AND item = $2", category, item)
	err := row.Scan(&amount)

	if err != nil {
		panic(err)
	} else if amount == 0 {
		return false
	} else {
		var balance int
		row := db.QueryRow("SELECT price FROM catalog WHERE category = $1 AND item = $2", category, item)
		row.Scan(&balance)

		if balance > client.balance {
			return false
		}
		client.balance -= balance
		add_usr_balance(client.username, client.balance)
		return true
	}
 }

func prepareDb() {
	var err error
	db, err = sql.Open("postgres", DbInfo)
	checkErr(err)

	_, err = db.Exec(CreateUsersTable)
	checkErr(err)

	_, err = db.Exec(CreateCatalogTable)
	checkErr(err)
}