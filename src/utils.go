package main

import (
	"fmt"
	"bufio"
	"os"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func greeting_menu(client User) {
	log("---------------------------------------------------------")
	log("Welcome to amazon!")
	log("We are glad to introduce you our big shop")
	log("---------------------------------------------------------")
	if client.logged_in() {
		fmt.Printf("Your balance: %d\n", client.get_balance())
		log("Choose option")
		log("	1. Edit profile")
		log("	2. Show goods categories")
		log("	3. Increase balance")
		log("	4. Buy item")
		log("	5. Add item (need admin password)")
		log("	6. Add category (need admin password)")
		log("	7. Logout")
	} else {
		log("Choose option")
		log("	1. Register")
		log("	2. Login")
		log("	3. Exit")
	}
	fmt.Print("Your choice: ")
}

func get_option() int {
	var input int
	fmt.Scanf("%d", &input)
	return input
}

func read_string() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[:len(input) - 1]
}

func log(str string) {
	fmt.Println(str)
}

func admin_check() bool {
	log("Do you know admin password?")
	admin_pass := read_string()

	if admin_pass == "123" {
		return true
	} else {
		log("Incorrect admin password!")
		return false
	}
}