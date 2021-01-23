package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	u1 := uuid.Must(uuid.NewV4(), nil)
	fmt.Printf("UUIDv4: %s\n", u1)
}
