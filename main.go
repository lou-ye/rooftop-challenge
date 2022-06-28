package main

import (
	"fmt"
	"rooftop-challenge/api"
	"rooftop-challenge/handler"
)

func main() {
	result, err := handler.Handler{RooftopApi: api.RooftopApiImpl{}}.HandleRequest()
	if err != nil {
		fmt.Println("Error executing program: ", err)
	}

	if result {
		fmt.Println("Lo resolviste correctamente")
	} else {
		fmt.Println("Todavía puedes intentarlo")
	}
}
