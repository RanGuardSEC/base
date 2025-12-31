package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		showMenu()

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			encodeMode(reader)
		case "2":
			decodeMode(reader)
		case "3":
			fmt.Println("程序退出")
			return
		default:
			fmt.Println("无效选择，请输入 1、2 或 3")
		}
	}
}

func showMenu() {
	fmt.Println("1. 编码模式")
	fmt.Println("2. 解码模式")
	fmt.Println("3. 退出程序")
	fmt.Print("请选择 (1/2/3): ")
}

func encodeMode(reader *bufio.Reader) {
	fmt.Println("\n--- 编码模式 ---")
	fmt.Println("输入 'back' 返回主菜单")

	for {
		fmt.Print("请输入要编码的文本: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "back" {
			fmt.Println("返回主菜单...")
			return
		}

		if text == "" {
			fmt.Println("输入不能为空")
			continue
		}

		encoded := base64.StdEncoding.EncodeToString([]byte(text))
		fmt.Printf("编码结果: %s\n\n", encoded)
	}
}

func decodeMode(reader *bufio.Reader) {
	fmt.Println("\n--- 解码模式 ---")
	fmt.Println("输入 'back' 返回主菜单")

	for {
		fmt.Print("请输入要解码的Base64字符串: ")
		encodedText, _ := reader.ReadString('\n')
		encodedText = strings.TrimSpace(encodedText)

		if encodedText == "back" {
			fmt.Println("返回主菜单...")
			return
		}

		if encodedText == "" {
			fmt.Println("输入不能为空")
			continue
		}

		decoded, err := base64.StdEncoding.DecodeString(encodedText)
		if err != nil {
			fmt.Println("解码失败: 请输入有效的Base64字符串")
			continue
		}

		fmt.Printf("解码结果: %s\n\n", string(decoded))
	}
}
