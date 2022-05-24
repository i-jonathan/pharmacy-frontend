package inventory

import (
	"PharmUI/global"
	"fmt"
	"html/template"
	"net/http"
)

func ViewInventory(w http.ResponseWriter, r *http.Request) {
	result := &inventoryResponse{}
	global.Getter("inventory/all", result)
	fmt.Println(result)

	viewTemplate := template.Must(template.ParseFiles("templates/inventory/view.html"))

	viewTemplate.Execute(w, result)
}
