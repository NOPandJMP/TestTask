package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	_ "github.com/lib/pq"
)

var (
	//line to connect to our database
	connStr = "user=postgres password=0000 dbname=users sslmode=disable"
	mutex   sync.Mutex
	db      *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	//Getting our structure object from json
	user := DeserializeRequest(w, r)
	_, err := db.Exec("Update Info Set pc_name = $1 , username = $2 , network_address=$3 Where pc_name = $1", user.PcName, user.UserName, user.NetworkAddr)
	if err != nil {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Problem with database operation ", 424)
	}
	fmt.Fprint(w, "Successfully update user information")

}

func newUser(w http.ResponseWriter, r *http.Request) {
	//Getting our structure object from json
	user := DeserializeRequest(w, r)
	_, err := db.Exec("insert into Info (pc_name,username ,network_address) values ($1, $2 , $3)", user.PcName, user.UserName, user.NetworkAddr)
	if err != nil {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Problem with database operation ", 424)
	}
	fmt.Fprint(w, "Successfully adding user information")

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	//Getting our structure object from json
	user := DeserializeRequest(w, r)
	_, err := db.Exec("Delete from Info Where pc_name = $1", user.PcName)
	if err != nil {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Problem with database operation ", 424)
	}
	fmt.Fprint(w, "Successfully delte user information")

}

func WorkSpace(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	//We look up which method was received, and then we call the corresponding functions
	switch r.Method {
	case http.MethodPost:
		newUser(w, r)
	case http.MethodPut:
		updateUser(w, r)

	case http.MethodDelete:
		deleteUser(w, r)
	default:
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Please send a correct request!", 405)
	}
	mutex.Unlock()
}

func DeserializeRequest(w http.ResponseWriter, r *http.Request) *UserInfo {
	var req = new(UserInfo)
	body, err := ioutil.ReadAll(r.Body)
	//Check for an error
	if err != nil {
		fmt.Fprintf(w, "err %q\n", err.Error())
	} else {
		//If everything is OK, we write by pointer in the structure
		err = json.Unmarshal(body, &req)
		if err != nil {
			fmt.Println(w, "can't unmarshal: ", err.Error())
		}
	}
	return req
}
