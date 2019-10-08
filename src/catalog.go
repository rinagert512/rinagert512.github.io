package main

import "fmt"

type Item struct {
	price int
	user_added string
}

type Catalog interface {
	init()
	show_categories()
	add_category() bool
	add_item(username string) bool
}

type CatStruct struct {
	categories map[string]map[string]Item
}

func (cat *CatStruct) init() {
	ok := get_data_from_catdb(cat);
	if !ok {
		log("Db error")
	}
}

func (cat *CatStruct) show_categories() {
	counter := 0

	for category, items := range cat.categories {
		counter = counter + 1
		fmt.Printf("%d. %s\n", counter, category)
		for item, Struct := range items {
			fmt.Printf("  %s, price: %d\n", item, Struct.price)
		}
	}
}
 
func (cat *CatStruct) add_category() bool {
	fmt.Print("What category do you want to add: ")
	category_name := read_string()
	_, ok := cat.categories[category_name]

	if ok {
		log("Category already exists")
		return false
	}
	cat.categories[category_name] = make(map[string]Item)
	add_to_catdb_item(category_name, "", "nil", -1)

	log("Successfully added!")
	return true
}

func (cat *CatStruct) add_item(username string) bool {
	cat.show_categories()
	fmt.Print("What category do you want to add the item to: ")
	category_name := read_string()
	_, ok := cat.categories[category_name]

	if !ok {
		log("No such category.")
		return false
	}

	fmt.Print("What item do you want to add: ")
	item_name := read_string()
	_, ok = cat.categories[category_name][item_name]
	if ok {
		log("Item already exists")
		return false
	}

	fmt.Print("Price of item: ")
	price := get_option()
	cat.categories[category_name][item_name] = Item{price, username}
	add_to_catdb_item(category_name, item_name, username, price)

	log("Successfully added!")
	return true
}