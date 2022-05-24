package account

import (
	"PharmUI/global"
	"fmt"
)

func ViewAccounts() {
	result := new(accountResponse)
	err := global.Getter("account/all", result)

	if err != nil {
		fmt.Println("do something with the error")
	}

	fmt.Println(result.Data[0].Email)
}
