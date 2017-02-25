package main

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID        int    `storm:"id,increment"`
	Group     string `storm:"index"`
	Email     string `storm:"unique"`
	Name      string
	Age       int
	CreatedAt time.Time `storm:"index"`
}

type Client struct {
	//ID        int       `storm:"id,increment"`
	Name      string    `storm:"id,index"`
	CreatedAt time.Time `storm:"index"`
}

type Project struct {
	ID        int       `storm:"id,increment"`
	Name      string    `storm:"index"`
	CreatedAt time.Time `storm:"index"`
}

type TimeEntry struct {
	ID        int       `storm:"id,increment"`
	Comment   string    `storm:"index"`
	CreatedAt time.Time `storm:"index"`
	StartTime time.Time `storm:"index"`
	EndTime   time.Time `storm:"index"`
}

func insertClient(name string) Client {

	clients := DB.From("goTT", "clients")

	client := getClientByName(name)

	if client == nil {
		client = &Client{
			Name:      name,
			CreatedAt: time.Now(),
		}
	}
	err := clients.Save(client)

	if err != nil {
		log.Fatalln(err)
	}

	return *client
}

func insertProject(client *Client, name string) *Project {
	projects := DB.From("goTT", "projects")
	project := &Project{
		Name:      name,
		CreatedAt: time.Now(),
	}
	err := projects.Save(project)
	if err != nil {
		log.Fatalln(err)
	}
	return project
}

func getClientByName(name string) *Client {
	var client Client
	clients := DB.From("goTT", "clients")
	err = clients.One("Name", name, &client)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

	return &client
}
