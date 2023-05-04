package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置HTTP路由
	router := gin.Default()
	router.POST("/", passwordGenerator)
	router.Static("/", "./static")
	err := router.Run(":3005")
	if err != nil {
		fmt.Println("error starting the server:", err)
		return
	}
}

type RequestBody struct {
	Lowercase     bool   `json:"lowercase"`
	Uppercase     bool   `json:"uppercase"`
	Numbers       bool   `json:"number"`
	Special       bool   `json:"special"`
	CustomSpecial string `json:"customSpecial"`
	Exclude       bool   `json:"exclude"`
	CustomExclude string `json:"customExclude"`
	Length        int    `json:"length" binding:"required"`
	Quantity      int    `json:"quantity" binding:"required"`
}

type ResponseBody struct {
	Passwords []string `json:"passwords"`
}

func passwordGenerator(c *gin.Context) {

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 生成密码
	passwords := []string{}
	for i := 0; i < body.Quantity; i++ {
		password := ""
		characters := ""
		if body.Lowercase {
			characters += "abcdefghijklmnopqrstuvwxyz"
		}
		if body.Uppercase {
			characters += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		}
		if body.Numbers {
			characters += "0123456789"
		}
		if body.Special {
			if body.CustomSpecial != "" {
				characters += body.CustomSpecial
			} else {
				characters += "!@#$%^&*"
			}
		}
		if body.Exclude {
			characters = strings.ReplaceAll(characters, "i", "")
			characters = strings.ReplaceAll(characters, "I", "")
			characters = strings.ReplaceAll(characters, "l", "")
			characters = strings.ReplaceAll(characters, "L", "")
			characters = strings.ReplaceAll(characters, "1", "")
			characters = strings.ReplaceAll(characters, "o", "")
			characters = strings.ReplaceAll(characters, "O", "")
			characters = strings.ReplaceAll(characters, "0", "")
			if body.CustomExclude != "" {
				for _, c := range body.CustomExclude {
					characters = strings.ReplaceAll(characters, string(c), "")
				}
			}
		}
		for j := 0; j < body.Length; j++ {
			password += string(characters[rand.Intn(len(characters))])
		}
		passwords = append(passwords, password)
	}
	c.JSON(http.StatusOK, ResponseBody{Passwords: passwords})

}
