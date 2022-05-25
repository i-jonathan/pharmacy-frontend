package inventory

import (
	"PharmUI/global"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func ViewInventory(w http.ResponseWriter, r *http.Request) {
	result := &inventoryResponse{}
	err := global.Getter("inventory/all", result)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)

	viewTemplate := template.Must(template.ParseFiles("templates/inventory/view.html"))

	viewTemplate.Execute(w, result)
}
