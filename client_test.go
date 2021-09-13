package main

import (
	"log"
	"net"
	"os"
	"os/user"
	"strings"
	"testing"
)

func TestUserInfo(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	//getting a computer name
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	if user.Username != "zkk" {
		t.Errorf("Get username error ,  %s not to equal %s", user.Username, "zkk")
	}

	if hostname != "ubuntu" {
		t.Errorf("Get hostname error ,  %s not to equal %s", hostname, "ubuntu")
	}

}

func TestGetIp(t *testing.T) {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()

	index := strings.Index(localAddr, ":")
	if localAddr[:index] != "192.168.187.129" {
		t.Errorf("192.168.187.129 not equal to %s", localAddr)
	}

}

func TestUserArgs(t *testing.T) {
	useArgs := []string{"new", "update", "delete"}

	for _, arg := range useArgs {
		switch arg {
		case "new":
			t.Logf("good")
		case "update":
			t.Logf("good")
		case "delete":
			t.Logf("good")
		default:
			t.Logf("bad arg")
		}
	}
}
