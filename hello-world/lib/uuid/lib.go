package lib

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func NewV4() string {
	return fmt.Sprintf("Hello world!!! %v", uuid.NewV4())
}
