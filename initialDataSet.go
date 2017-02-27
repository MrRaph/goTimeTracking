package main

import (
	"fmt"
	"log"
	"time"
)

func insertSampleClient() {
	//clientsB := DB.From("goTT", "clients")
	err = DB.Drop("goTT")
	if err != nil {
		//log.Fatal(err)

		fmt.Println(err)
	}

	fmt.Println("> Insertion des clients")
	client1 := insertClient("ACME inc")
	fmt.Println(client1)

	client2 := insertClient("Example")
	fmt.Println(client2)

	client3 := insertClient("ACME inc")
	fmt.Println(client3)

	fmt.Println("> Insertion des projets")
	project1 := insertProject(client1, "Oracle")
	fmt.Println(project1)

	project2 := insertProject(client1, "Linux")
	fmt.Println(project2)

	project3 := insertProject(client2, "Docker")
	fmt.Println(project3)

	project4 := insertProject(client2, "Docker")
	fmt.Println(project4)

	fmt.Println("> Insertion des Task")
	te1 := insertTask(client1, project1, "initialisation")
	fmt.Println(te1)

	te2 := insertTask(client1, project1, "documentation")
	fmt.Println(te2)

	te3 := insertTask(client1, project1, "mise en production")
	fmt.Println(te3)

	fmt.Println("> Sleep")
	fmt.Println("> Mise à jour des Task")
	time.Sleep(10)

	te3.EndTime = time.Now()
	updateTask(te3)
	fmt.Println(te3)

	if te3.isTaskEnded() {
		fmt.Println("Tâche terminée")
	} else {
		fmt.Println("Tâche non terminée !")
	}

	fmt.Println("> Récupération des clients")
	var clients []Client
	err := DB.From("goTT", "clients").All(&clients)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found", len(clients))

	fmt.Println("> Récupération de tous les projets")
	var projects []Project
	err = DB.From("goTT", "projects").All(&projects)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found", len(projects))

	fmt.Println("> Récupération des projets du client 1")
	fmt.Println(getProjectByClient(client1.Name))
}
