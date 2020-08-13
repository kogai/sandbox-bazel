package main

import (
	"fmt"

	lib "lib/uuid"

	uuid "github.com/satori/go.uuid"
)

func main() {
	fmt.Println("Hello world!!!", uuid.NewV4(), lib.NewV4())
}
