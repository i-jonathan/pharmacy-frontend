package main

import (
	"PharmUI/account"
	"net/http"
)

func main() {
	//inventory.ViewInventory()
	fs := http.FileServer(http.Dir("templates/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//fileServer := http.FileServer(http.Dir("templates/"))
	//http.Handle("/", http.StripPrefix("/", fileServer))
	//http.HandleFunc("/inventory/view/", inventory.ViewInventory)
	http.HandleFunc("/account/register", account.RegisterAccount)

	//port := os.Getenv("PORT")
	port := "7060"
	_ = http.ListenAndServe(":"+ port, nil)
}
