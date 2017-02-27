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

type Task struct {
	ID          int       `storm:"id,increment"`
	ClientName  string    `storm:"id,index"`
	ProjectName string    `storm:"id,index"`
	Comment     string    `storm:"index"`
	CreatedAt   time.Time `storm:"index"`
	StartTime   time.Time `storm:"index"`
	EndTime     time.Time `storm:"index"`
}

/******************************************************************************
*** Insert
******************************************************************************/

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

func insertTask(client *Client, project *Project, comment string) *Task {
	timeEntriesBucket := DB.From("goTT", "timeentries")
	timeEntry := &Task{
		ClientName:  client.Name,
		ProjectName: project.Name,
		Comment:     comment,
		CreatedAt:   time.Now(),
		StartTime:   time.Now(),
		EndTime:     time.Now().AddDate(0, -1, 0),
	}

	err := timeEntriesBucket.Save(timeEntry)
	if err != nil {
		log.Fatalln(err)
	}
	return timeEntry
}

func updateTask(timeEntry *Task) *Task {
	timeEntriesBucket := DB.From("goTT", "timeentries")
	err := timeEntriesBucket.Update(timeEntry)
	if err != nil {
		log.Fatalln(err)
	}
	return timeEntry
}

/******************************************************************************
*** Select
******************************************************************************/

func getTaskByProject(clientName, projectName string) *[]Task {
	var timeEntries []Task
	timeEntryBucket := DB.From("goTT", "timeentries")

	err = timeEntryBucket.Select(q.Eq("ClientName", clientName), q.Eq("ProjectName", projectName)).Find(&timeEntries)

	if err != nil {
		fmt.Println(err)
		timeEntries[0] = Task{
			ID:          -1,
			ClientName:  "",
			ProjectName: "",
		}
		return &timeEntries
	}

	return &timeEntries
}

func getTaskFromComment(comment string) (client *Client, project *Project, timeentry *Task) {

	timeEntryBucket := DB.From("goTT", "timeentries")
	err = timeEntryBucket.Select(q.Eq("Comment", comment)).Find(timeentry)

	client = getClientByName(timeentry.ClientName)
	project = getProjectByName(timeentry.ClientName, timeentry.ProjectName)

	return
}

func getProjectByName(clientName, name string) *Project {
	var project []Project
	projects := DB.From("goTT", "projects")

	err = projects.Select(q.Eq("ClientName", clientName), q.Eq("Name", name)).Find(&project)

	if err != nil {
		fmt.Println(err)

		return &Project{
			ClientName: "",
			Name:       "",
		}
	} else {
		if len(project) > 1 {
			fmt.Println("More than one Projects")
		}
	}

	return &project[0]
}

func getProjectByClient(clientName string) *[]Project {
	var project []Project
	projects := DB.From("goTT", "projects")

	err = projects.Select(q.Eq("ClientName", clientName)).Find(&project)

	if err != nil {
		fmt.Println(err)

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
