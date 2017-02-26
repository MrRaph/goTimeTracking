package main

import (
	"fmt"
	"log"
	"time"

	"github.com/asdine/storm/q"
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
	Name      string    `storm:"id,index"`
	CreatedAt time.Time `storm:"index"`
}

type Project struct {
	Name       string    `storm:"id,index"`
	ClientName string    `storm:"id,index"`
	CreatedAt  time.Time `storm:"index"`
}

type TimeEntry struct {
	ID        int       `storm:"id,increment"`
	Comment   string    `storm:"index"`
	CreatedAt time.Time `storm:"index"`
	StartTime time.Time `storm:"index"`
	EndTime   time.Time `storm:"index"`
}

func insertClient(name string) *Client {

	clients := DB.From("goTT", "clients")

	client := getClientByName(name)

	if client.Name == "" {
		client = &Client{
			Name:      name,
			CreatedAt: time.Now(),
		}

		err := clients.Save(client)

		if err != nil {
			log.Fatalln(err)
		}
	}

	return client
}

func insertProject(client *Client, name string) *Project {
	projects := DB.From("goTT", "projects")
	project := getProjectByName(client.Name, name)

	if project.Name == "" {
		project = &Project{
			Name:       name,
			ClientName: client.Name,
			CreatedAt:  time.Now(),
		}

		err := projects.Save(project)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return project
}

func getProjectByName(clientName, name string) *Project {
	var project Project
	projects := DB.From("goTT", "projects")

	err = projects.Select(q.Gte("ClientName", clientName), q.Lte("Name", name)).Find(&project)

	if err != nil {
		fmt.Println(err)
		return &Project{
			ClientName: "",
			Name:       "",
		}
	}
	return &project
}

func getClientByName(name string) *Client {
	var client Client
	clients := DB.From("goTT", "clients")
	err = clients.One("Name", name, &client)
	if err != nil {
		fmt.Println(err)
		return &Client{
			Name: "",
		}
	}
	return &client
}
