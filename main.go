package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	// 设置HTTP路由
	http.HandleFunc("/", passwordGenerator)
	// 设置静态文件路由
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	// 启动HTTP服务器
	listener, err := net.Listen("tcp4", ":3005")
	if err != nil {
		fmt.Println("error starting the server:", err)
		return
	}
	err = http.Serve(listener, nil)
	if err != nil {
		fmt.Println("error serving the request:", err)
	}
}

func passwordGenerator(w http.ResponseWriter, r *http.Request) {
	// 设置响应头的Content-Type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// 获取HTTP表单中的参数
	lowercase := r.FormValue("lowercase") == "on"                     // 是否包含小写字母
	uppercase := r.FormValue("uppercase") == "on"                     // 是否包含大写字母
	numbers := r.FormValue("numbers") == "on"                         // 是否包含数字
	special := r.FormValue("special") == "on"                         // 是否包含特殊字符
	custom := r.FormValue("custom")                                   // 自定义的特殊字符
	exclude := r.FormValue("exclude") == "on"                         // 是否排除相似字符
	customExclude := r.FormValue("custom-exclude")                    // 自定义要排除的字符
	length, _ := strconv.Atoi(r.FormValue("length"))                  // 密码长度
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))              // 生成密码的数量
	lastLength, _ := strconv.Atoi(r.PostFormValue("last-length"))     // 最后生成密码时的密码长度
	lastQuantity, _ := strconv.Atoi(r.PostFormValue("last-quantity")) // 最后生成密码时的生成密码的数量

	// 如果用户没有进行生成密码操作，那么上面两个变量的值都为0
	hasGenerated := lastLength != 0 && lastQuantity != 0

	// 如果有生成过密码，那么将参数的值替换为最后的值
	if hasGenerated {
		length = lastLength
		quantity = lastQuantity
	}

	// 生成密码
	passwords := []string{}
	for i := 0; i < quantity; i++ {
		password := ""
		characters := ""
		if lowercase {
			characters += "abcdefghijklmnopqrstuvwxyz"
		}
		if uppercase {
			characters += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		}
		if numbers {
			characters += "0123456789"
		}
		if special {
			if custom != "" {
				characters += custom
			} else {
				characters += "!@#$%^&*"
			}
		}
		if exclude {
			characters = strings.ReplaceAll(characters, "i", "")
			characters = strings.ReplaceAll(characters, "I", "")
			characters = strings.ReplaceAll(characters, "l", "")
			characters = strings.ReplaceAll(characters, "L", "")
			characters = strings.ReplaceAll(characters, "1", "")
			characters = strings.ReplaceAll(characters, "o", "")
			characters = strings.ReplaceAll(characters, "O", "")
			characters = strings.ReplaceAll(characters, "0", "")
			if customExclude != "" {
				for _, c := range customExclude {
					characters = strings.ReplaceAll(characters, string(c), "")
				}
			}
		}
		for j := 0; j < length; j++ {
			password += string(characters[rand.Intn(len(characters))])
		}
		passwords = append(passwords, password)
	}

	// 生成HTML响应
	html := "<!DOCTYPE html><html><head><title>Password Generator</title><link rel=\"stylesheet\" href=\"/static/style.css\"></head><body><h1>Password Generator</h1><form method=\"post\"><p><label><input type=\"checkbox\" name=\"lowercase\" checked> 小写字母 (a-z)</label></p><p><label><input type=\"checkbox\" name=\"uppercase\" checked> 大写字母 (A-Z)</label></p><p><label><input type=\"checkbox\" name=\"numbers\" checked> 数字 (0-9)</label></p><p><label><input type=\"checkbox\" name=\"special\" checked> 特殊字符 (!@#$%^&amp;*)</label> <input type=\"text\" name=\"custom\" placeholder=\"自定义特殊字符\"></p><p><label><input type=\"checkbox\" name=\"exclude\"> 排除相似字符 (i, I, l, L, 1, o, O, 0)</label> <input type=\"text\" name=\"custom-exclude\" placeholder=\"自定义要排除的字符\"></p><p><label>密码长度: <input type=\"number\" name=\"length\" value=\"" + strconv.Itoa(length) + "\" min=\"1\" max=\"128\"></label></p><p><label>生成密码的数量: <input type=\"number\" name=\"quantity\" value=\"" + strconv.Itoa(quantity) + "\" min=\"1\" max=\"100\"></label></p><input type=\"hidden\" name=\"last-length\" value=\"" + strconv.Itoa(length) + "\"><input type=\"hidden\" name=\"last-quantity\" value=\"" + strconv.Itoa(quantity) + "\"><p><button type=\"submit\">生成密码</button></p></form>"

	if len(passwords) > 0 {
		html += "<h2>生成的密码</h2>"
		for _, password := range passwords {
			html += "<p>" + password + "</p>"
		}
	}
	html += "<script src=\"/static/script.js\"></script></body></html>"
	fmt.Fprint(w, html)
}
