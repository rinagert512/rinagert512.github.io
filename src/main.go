package main

func main() {
	var client User = &UsrStruct{false, "", "", 0}
	var catalog Catalog = &CatStruct{make(map[string]map[string]Item)}

	prepareDb()
	catalog.init()
	defer db.Close()

	for true {
		greeting_menu(client)
		choice := get_option()
		if client.logged_in() {
			if choice == 1 {
				client.edit_profile()
			} else if choice == 2 {
				catalog.show_categories()
			} else if choice == 3 {
				client.increase_balance()
			} else if choice == 4 {
				client.buy()
			} else if choice == 5 {
				if admin_check() {
					catalog.add_item(client.get_login())
				}
			} else if choice == 6 {
				if admin_check() {
					catalog.add_category()
				}
			} else if choice == 7 {
				client.logout()
			} else {
				log("Incorrect choice")
			}
		} else {
			if choice == 1 {
				register()
			} else if choice == 2 {
				login(&client)
			} else if choice == 3 {
				break
			} else {
				log("Incorrect choice")
			}
		}
	}
}