package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	special   = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

type PasswordOptions struct {
	Lowercase     bool   `json:"lowercase"`
	Uppercase     bool   `json:"uppercase"`
	Numbers       bool   `json:"numbers"`
	Special       bool   `json:"special"`
	CustomSpecial string `json:"customSpecial"`
	Exclude       bool   `json:"exclude"`
	CustomExclude string `json:"customExclude"`
	Length        int    `json:"length"`
	Quantity      int    `json:"quantity"`
}

func GeneratePasswords(options PasswordOptions) ([]string, error) {
	fmt.Printf("\n开始生成安全密码:\n")
	fmt.Printf("使用以下字符集:\n")
	if options.Lowercase {
		fmt.Printf("小写字母: %s\n", lowercase)
	}
	if options.Uppercase {
		fmt.Printf("大写字母: %s\n", uppercase)
	}
	if options.Numbers {
		fmt.Printf("数字: %s\n", numbers)
	}
	if options.Special {
		if options.CustomSpecial != "" {
			fmt.Printf("特殊字符(自定义): %s\n", options.CustomSpecial)
		} else {
			fmt.Printf("特殊字符: %s\n", special)
		}
	}
	fmt.Printf("密码长度: %d, 生成数量: %d\n", options.Length, options.Quantity)

	if options.Length < 4 {
		return nil, errors.New("密码长度必须大于等于4")
	}

	if options.Quantity < 1 || options.Quantity > 20 {
		return nil, errors.New("生成数量必须在1-20之间")
	}

	if options.Length > 100 {
		return nil, errors.New("密码长度不能超过100")
	}

	// 构建字符集
	var charset strings.Builder
	if options.Lowercase {
		charset.WriteString(lowercase)
	}
	if options.Uppercase {
		charset.WriteString(uppercase)
	}
	if options.Numbers {
		charset.WriteString(numbers)
	}
	if options.Special {
		if options.CustomSpecial != "" {
			charset.WriteString(options.CustomSpecial)
		} else {
			charset.WriteString(special)
		}
	}

	if charset.Len() == 0 {
		return nil, errors.New("至少需要选择一种字符类型")
	}

	// 处理排除字符
	charsetStr := charset.String()
	if options.Exclude && options.CustomExclude != "" {
		for _, c := range options.CustomExclude {
			charsetStr = strings.ReplaceAll(charsetStr, string(c), "")
		}
	}

	if len(charsetStr) == 0 {
		return nil, errors.New("排除字符后没有可用字符")
	}

	passwords := make([]string, options.Quantity)
	for i := 0; i < options.Quantity; i++ {
		password, err := generatePassword(charsetStr, options.Length, options)
		if err != nil {
			return nil, err
		}
		passwords[i] = password
	}

	return passwords, nil
}

func generatePassword(charset string, length int, options PasswordOptions) (string, error) {
	if length < 4 {
		return "", errors.New("密码长度太短")
	}

	result := make([]byte, length)
	requiredCount := 0

	// 计算需要包含的必需字符数量
	if options.Lowercase {
		requiredCount++ // 如果需要小写字母，计数加1
	}
	if options.Uppercase {
		requiredCount++ // 如果需要大写字母，计数加1
	}
	if options.Numbers {
		requiredCount++ // 如果需要数字，计数加1
	}
	if options.Special {
		requiredCount++
	}

	if length < requiredCount {
		return "", errors.New("密码长度不足以包含所有必需的字符类型")
	}

	// 先确保每种选中的字符类型都至少出现一次
	pos := 0
	if options.Lowercase {
		// 随机选择一个小写字母
		n, _ := rand.Int(rand.Reader, big.NewInt(26))
		result[pos] = lowercase[n.Int64()]
		pos++
	}
	if options.Uppercase {
		// 随机选择一个大写字母
		n, _ := rand.Int(rand.Reader, big.NewInt(26))
		result[pos] = uppercase[n.Int64()]
		pos++
	}
	if options.Numbers {
		// 随机选择一个数字
		n, _ := rand.Int(rand.Reader, big.NewInt(10))
		result[pos] = numbers[n.Int64()]
		pos++
	}
	if options.Special {
		//
		specialChars := special
		if options.CustomSpecial != "" {
			specialChars = options.CustomSpecial
		}
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(specialChars))))
		result[pos] = specialChars[n.Int64()]
		pos++
	}

	// 处理排除字符
	availableChars := charset
	if options.Exclude && options.CustomExclude != "" {
		for _, c := range options.CustomExclude {
			// 移除排除字符
			availableChars = strings.ReplaceAll(availableChars, string(c), "")
		}
	}

	if len(availableChars) == 0 {
		return "", errors.New("排除字符后没有可用字符")
	}

	charsetLength := big.NewInt(int64(len(availableChars)))

	// 填充剩余位置
	for i := pos; i < length; i++ {
		// 随机选择一个可用字符
		n, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		// 将选中的字符添加到结果中
		result[i] = availableChars[n.Int64()]
	}

	// 打乱密码
	for i := length - 1; i > 0; i-- {
		// 随机选择一个位置
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		// 交换两个���置的字符
		result[i], result[j.Int64()] = result[j.Int64()], result[i]
	}

	// 验证生成的密码是否包含所有必需的字符类型
	passwordStr := string(result)
	fmt.Printf("\n验证密码安全性: '%s'\n", passwordStr)

	if options.Lowercase {
		hasLower := containsAny(passwordStr, lowercase)
		fmt.Printf("小写字母要求: %v\n", hasLower)
		if !hasLower {
			return generatePassword(charset, length, options)
		}
	}
	if options.Uppercase {
		hasUpper := containsAny(passwordStr, uppercase)
		fmt.Printf("大写字母要求: %v\n", hasUpper)
		if !hasUpper {
			return generatePassword(charset, length, options)
		}
	}
	if options.Numbers {
		hasNumber := containsAny(passwordStr, numbers)
		fmt.Printf("数字要求: %v\n", hasNumber)
		if !hasNumber {
			return generatePassword(charset, length, options)
		}
	}
	if options.Special {
		specialChars := special
		if options.CustomSpecial != "" {
			specialChars = options.CustomSpecial
		}
		hasSpecial := containsAny(passwordStr, specialChars)
		fmt.Printf("特殊字符要求: %v\n", hasSpecial)
		if !hasSpecial {
			return generatePassword(charset, length, options)
		}
	}

	fmt.Printf("已生成符合所有安全要求的强密码\n")
	return passwordStr, nil
}

// 验证密码是否包含指定类型的字符
func containsAny(password, charSet string) bool {
	for _, char := range password {
		for _, setChar := range charSet {
			if char == setChar {
				fmt.Printf("成功包含所需字符: %c\n", char)
				return true
			}
		}
	}
	fmt.Printf("需要包含字符集 [%s] 中的字符，重新生成中...\n", charSet)
	return false
}
