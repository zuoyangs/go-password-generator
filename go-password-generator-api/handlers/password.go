package handlers

import (
	"fmt"
	"net/http"
	"password-generator-api/utils"

	"github.com/gin-gonic/gin"
)

func HandleGeneratePassword(c *gin.Context) {
	// 获取请求的 Origin
	origin := c.Request.Header.Get("Origin")
	// 允许的域名列表
	allowedOrigins := []string{
		"https://password.zuoyang.tech",
		"http://password.zuoyang.tech",
		"http://localhost:5173",  // Vue.js 开发服务器默认端口
		"http://localhost:8080",
	}

	// 检查是否是允许的域名
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			c.Header("Access-Control-Allow-Origin", origin)
			break
		}
	}
	
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	var options utils.PasswordOptions
	if err := c.ShouldBindJSON(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	fmt.Printf("收到前端请求参数:\n")
	fmt.Printf("字符选项: 小写字母=%v, 大写字母=%v, 数字=%v, 特殊字符=%v\n",
		options.Lowercase, options.Uppercase, options.Numbers, options.Special)
	fmt.Printf("密码长度: %d\n", options.Length)
	fmt.Printf("生成数量: %d\n", options.Quantity)
	if options.Special && options.CustomSpecial != "" {
		fmt.Printf("自定义特殊字符: %s\n", options.CustomSpecial)
	}
	if options.Exclude && options.CustomExclude != "" {
		fmt.Printf("排除字符: %s\n", options.CustomExclude)
	}

	passwords, err := utils.GeneratePasswords(options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"passwords": passwords})
}

