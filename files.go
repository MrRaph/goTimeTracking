package main

import "os"

func getDataDir() string {
	return os.Getenv("APPDATA") + "\\" + appName
}

func createDataDir() {
	path := getDataDir()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}
