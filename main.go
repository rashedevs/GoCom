package main

import (
	"fmt"
	"gocom/util"
)

func main() {
	// cmd.Serve()
	jwt, err := util.CreateJwt("my-secret", util.Payload{
		Sub:         31,
		FirstName:   "Rashed",
		LastName:    "Uz Zaman",
		Email:       "rasheduap2015@gmail.com",
		IsShopOwner: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jwt)
}

/*
### JSON : javascript object notation.
	[] -> List
	{} -> Object
*/
