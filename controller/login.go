package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	jsonData, _ := json.Marshal(map[string]string{
		"email":    c.PostForm("email"),
		"password": c.PostForm("password"),
	})

	url := "http://localhost:8000/users/login"
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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

	c.SetCookie("token", "cookievalue", 3600, "/", "localhost", false, true)
	//SetCookie (name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	c.HTML(
		http.StatusOK,
		"views/login.html",
		gin.H{
			"title": "Geeksbeginner",
		},
	)

}
