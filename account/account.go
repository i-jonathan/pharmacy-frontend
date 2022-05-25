package account

import (
	"PharmUI/global"
	"PharmUI/models"
	"html/template"
	"log"
	"net/http"
)

func ViewAccounts() {
	//result := new(accountResponse)
	//err := global.Getter("account/all", result)
	//
	//if err != nil {
	//	fmt.Println("do something with the error")
	//}
	//
	//fmt.Println(result.Data[0].Email)
}

func RegisterAccount(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/account/register.gohtml"))

	if r.Method != http.MethodPost {
		err := tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
		return
	}

	if r.FormValue("password") != r.FormValue("retype-password") {
		log.Println("return with error message")
		err := tmpl.Execute(w, struct {
			Message string
			Error bool
		}{"Please check your password and retry", true})
		log.Println(err)
		return
	}

	newUser := models.Account{
		FirstName:   r.FormValue("first-name"),
		LastName:    r.FormValue("last-name"),
		PhoneNumber: r.FormValue("phone"),
		Email:       r.FormValue("email"),
		Password:    r.FormValue("password"),
	}

	resp, err := global.Poster("account", newUser)
	if err != nil {
		log.Println(err)
	}

	if resp.Status != http.StatusOK {
		// TODO display error
	}
	loginTemplate := template.Must(template.ParseFiles("templates/account/login.html"))
	err = loginTemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}