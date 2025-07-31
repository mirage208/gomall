package jwtx

import (
	"fmt"
	"log"
	"time"
)

func ExampleGetToken() {
	// Example usage of GetToken function
	now := time.Now().Unix()
	secret := "uOvKLmVfztaXGpNYd4Z0I1SiT7MweJhl"
	token, err := GetToken(secret, now, 86400, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)
	// Output: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTM5Nzk3MTcsImlhdCI6MTc1Mzg5MzMxNywidWlkIjoxfQ.mbPU6gjjG9GptMdGyikeblOBWJ6HmAXexoiJUpAcjAQ
}
