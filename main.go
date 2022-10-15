package main

import (
	"fmt"
	"os"
)

var clients []Client

func main() {

	clients = []Client {
		{1, "921234561", "exampple1@gmail.com"},
		{2, "921234562", "exampple2@gmail.com"},
		{3, "921234563", "exampple3@gmail.com"},
		{4, "921234564", "exampple4@gmail.com"},
		{5, "921234565", "exampple5@gmail.com"},
	}

	for _, val := range clients {
		fmt.Println("ID: ", val.ID, " Phone: ", val.Phone, " Email: ", val.Email)
	}

	var id int
	var client *Client

	for {
		fmt.Println("Select client ID...")
		fmt.Scan(&id)
		if client := getByID(id); client != nil {
			break
		} else {
			fmt.Println("Undefined client ID")
		}
	}

	var sendType string
	sended := false
	fmt.Println("Send types: sms or email")
	for {
		fmt.Println("Select send type...")
		fmt.Scan(&sendType)
		switch sendType {
		case "email":
			err := client.notifyEmail()
			if err != nil {
				ErrorHandle(err)
			}
			sended = true
		case "sms":
			err := client.notifySMS()
			if err != nil {
				ErrorHandle(err)
			}
			sended = true
		default:
			fmt.Println("Undefined send type error!")
		}
		if sended {
			break
		}
	}
}

type Product struct {
	ID      int
	Product string
	Price   float64
}

type Client struct {
	ID    int
	Phone string
	Email string
}

func getByID(id int) *Client {
	for _, val := range clients {
		if val.ID == id {
			return &val
		}
	}
	return nil
}

func (c *Client) notifyEmail() error {
	//TODO SEND EMAIL
	fmt.Println("EMAIL SENT!")
	return nil
}

func (c *Client) notifySMS() error {
	//TODO SEND SMS
	fmt.Println("SMS SENT!")
	return nil
}

func ErrorHandle(err error) {
	if err != nil {
		if !exists("errors.txt") {
			os.Create("errors.txt")
		}
		f, _ := os.OpenFile("errors.txt", os.O_APPEND|os.O_WRONLY, 0600)
		f.WriteString(err.Error() + "\r\n")
		f.Close()
	}
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
