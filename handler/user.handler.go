package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/abetobing/go-eventlistener-example/ext"
	"github.com/go-chi/render"
	"github.com/go-resty/resty/v2"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Job  string `json:"job"`
}

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	client := resty.New()

	url := "https://reqres.in/api/users"
	resp, err := client.R().
		EnableTrace().
		Get(url)
	if err != nil {
		fmt.Println("Error ", err)
		return
	}

	var user User
	err = json.Unmarshal(resp.Body(), &user)
	if err != nil {
		fmt.Println("Error ", err)
	}

	render.DefaultResponder(w, r, user)

	ext.SubmitElastic(user)

	elapsed := time.Since(start)
	fmt.Println("Elapsed time ", elapsed)

}
