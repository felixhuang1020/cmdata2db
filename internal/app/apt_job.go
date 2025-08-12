package app

import (
	"fmt"
	"io"
	"net/http"
)

func CallExternalAPI() {
	resp, err := http.Get("")
	if err != nil {
		fmt.Println("调用 API 失败:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("API 响应:", string(body))
}
