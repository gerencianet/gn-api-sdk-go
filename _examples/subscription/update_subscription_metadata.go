package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/_examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	body := map[string]interface{} {
		"custom_id": "Product 0001",
		"notification_url": "http://domain.com/notification",
	}

	res, err := gn.UpdateSubscriptionMetadata(1, body) // no lugar do 1 coloque o subscription_id correto

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}