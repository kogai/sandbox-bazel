package main

import (
	"fmt"

	lib "app/uuid"

	uuid "github.com/satori/go.uuid"
)

func main() {
	fmt.Println("Hello world!!!", uuid.NewV4(), lib.NewV4())
}
