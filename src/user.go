package main

import "fmt"

type User interface {
	logged_in() bool
	login(login, password string, balance int)
	edit_profile()
	increase_balance()
	buy()
	get_balance() int
	logout()
	get_login() string
}

type UsrStruct struct {
	logged bool
	username, password string
	balance int
}

func (client *UsrStruct) logged_in() bool {
	if client.logged {
		return true
	}
	return false
}

func (client *UsrStruct) login(login, password string, balance int) {
	client.username = login
	client.password = password
	client.logged = true
	client.balance = balance
}

func (client *UsrStruct) edit_profile() {
	for true {
		log("---------------------------------------------------------")
		log("What do you want to do?")
		log("	1. Change username")
		log("	2. Change password")
		log("	3. Exit")
		fmt.Print("Input your choice: ")
		choice := get_option()
		if choice == 1 {
			change_username(client)
		} else if choice == 2 {
			change_password(client)
		} else if choice == 3 {
			return
		} else {
			log("Invalid choice!")
		}
	}
}	

func (client *UsrStruct) increase_balance() {
	fmt.Print("How much money you want to add: ")
	money := get_option()

	if money < 0 {
		log("Something went wrong")
		return
	}

	client.balance += money
	add_usr_balance(client.username, client.balance)
	log("Successfully increased balance!")
}

func (client *UsrStruct) buy() {
	fmt.Print("Which item do you want to buy: ")
	item := read_string()
	
	fmt.Print("What category is the item in: ")
	category := read_string()

	ok := buy_item(item, category, client)
	if !ok {
		log("Error occured while buying.")
		return
	}
	log("Successfully bought!")
	return
}

func (client *UsrStruct) get_balance() int {
	return client.balance
}

func (client *UsrStruct) logout() {
	client.logged = false
	client.username = ""
	client.password = ""
	client.balance = 0
}

func (client *UsrStruct) get_login() string {
	return client.username
}