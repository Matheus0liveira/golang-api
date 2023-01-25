package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	resp, err := http.Get("http://localhost:8080/user")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println("Not Success", resp.StatusCode)

		return
	}

	var response []User

	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Data recovery Error", err.Error())
		return
	}

	fmt.Println(response)

}
