package anythingllm

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {

	client := NewClient()
	data, _ := client.ValidToken()

	fmt.Println("数据：", data)
}
