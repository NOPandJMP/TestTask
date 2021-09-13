package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/user"
	"strings"
)

var client = &http.Client{}

//Since interf.Addrs() returns multiple addresses, I decided to use this method
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()

	index := strings.Index(localAddr, ":")
	return localAddr[:index]
}

func main() {
	//user is only needed to get the current user name
	if len(os.Args) < 1 {
		fmt.Println("use go run main.go (new , delete, update) , you need to select one of the options without brackets")
		os.Exit(1)
	}
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	//getting a computer name
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	userInfo := new(UserInfo)
	userInfo.UserName = user.Username
	userInfo.PcName = hostname
	userInfo.NetworkAddr = GetOutboundIP()

	//Depending on what the user transmits through the instrument, we will perform certain actions
	switch os.Args[1] {
	case "new":
		postRequest(userInfo)
	case "update":
		updateRequest(userInfo)
	case "delete":
		deleteRequest(userInfo)
	default:
		fmt.Println("use go run main.go (new , delete, update) , you need to select one of the options without brackets")
		os.Exit(1)
	}

}

//Sending a post request to add a new user
func postRequest(user *UserInfo) {
	bytesRepresentation, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/workspace", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	response, err := client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer req.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)

	//handling Errors
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bodyBytes))
}

//Sending a put request to update info about user
func updateRequest(user *UserInfo) {
	bytesRepresentation, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/workspace", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	response, err := client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer req.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)

	//handling Errors
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bodyBytes))
}

//Sending a delete request to delete user
func deleteRequest(user *UserInfo) {
	bytesRepresentation, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/workspace", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	response, err := client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer req.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)

	//handling Errors
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bodyBytes))
}
