package fernet_test

import (
	"encoding/json"
	"fmt"
	"github.com/kr/fernet"
	"os"
	"time"
)

func ExampleKey_Generate() {
	k := fernet.MustDecodeKey(os.Getenv("MYSECRET"))
	token, err := k.Generate([]byte("hello"))
	if err == nil {
		fmt.Println(string(token))
	}
}

func ExampleKey_Verify() {
	k := fernet.MustDecodeKey(os.Getenv("MYSECRET"))
	token := []byte("…")
	var v struct {
		Username string
	}
	err := json.Unmarshal(k.Verify(token, 60*time.Second), &v)
	if err == nil {
		fmt.Println(v.Username)
	}
}
