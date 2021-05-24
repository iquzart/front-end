package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterUser struct {
}

func Register(c *gin.Context) {
	jsonData, _ := json.Marshal(map[string]string{
		"first_name": c.PostForm("firstname"),
		"last_name":  c.PostForm("lastname"),
		"email":      c.PostForm("email"),
		"password":   c.PostForm("password"),
		"phone":      c.PostForm("phone"),
	})

	url := "http://localhost:8000/users/signup"
	fmt.Println("URL:>", url)

	fmt.Println(string(jsonData))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	c.HTML(
		http.StatusOK,
		"views/register.html",
		gin.H{
			"title": "Register",
		},
	)

}
