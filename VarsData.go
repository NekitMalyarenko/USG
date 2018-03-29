package main


import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	BOT_TOKEN = "BOT_TOKEN"
)


func GetString(varName string) string {
	return os.Getenv(varName)
}


func GetBoolean(varName string) bool {
	return strings.ToLower(os.Getenv(varName)) == "true"
}


func GetInt(varName string) int {
	temp, err := strconv.Atoi(os.Getenv(varName))
	if err != nil {
		log.Fatal(err)
		return -1
	} else {
		return temp
	}
}
