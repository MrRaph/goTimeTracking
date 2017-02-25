package main

import (
	"fmt"
	"log"
)

func insertSampleClient() {
	//clientsB := DB.From("goTT", "clients")
	err = DB.Drop("goTT")
	if err != nil {
		//log.Fatal(err)

		fmt.Println(err)
	}

	client1 := insertClient("ACME inc")
	fmt.Println(client1)

	client2 := insertClient("Example")
	fmt.Println(client2)

	var clients []Client
	err := DB.From("goTT", "clients").All(&clients)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found", len(clients))
}
