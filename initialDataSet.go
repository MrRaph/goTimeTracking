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

	client3 := insertClient("ACME inc")
	fmt.Println(client3)

	project1 := insertProject(client1, "Oracle")
	fmt.Println(project1)

	project2 := insertProject(client1, "Linux")
	fmt.Println(project2)

	project3 := insertProject(client2, "Docker")
	fmt.Println(project3)

	project4 := insertProject(client2, "Docker")
	fmt.Println(project4)

	var clients []Client
	err := DB.From("goTT", "clients").All(&clients)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found", len(clients))

	var projects []Project
	err = DB.From("goTT", "projects").All(&projects)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found", len(projects))
}
