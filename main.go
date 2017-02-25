package main

import (
	"github.com/asdine/storm"
)

const (
	appName string = "goTT"
)

var DB, err = storm.Open(getDataDir() + "\\" + appName + ".db")

func main() {

	defer DB.Close()
	createDataDir()
	insertSampleClient()
}
