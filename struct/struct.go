package main

import (
	"encoding/json"
	"fmt"
)

type Client struct {
	Nome  string
	Email string
	CPF   int
}

func main() {
	client := Client{
		Nome:  "Wesley",
		Email: "w@w.com",
		CPF:   123456789,
	}

	fmt.Println(client)

	json.Unmarshal([]byte(`{"Nome":"Lucas"}`), &client)

	fmt.Println(client)
}
