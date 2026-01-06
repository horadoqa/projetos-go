package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// generate a uuid v4
	uuid := uuid.New()
	fmt.Println(uuid)
}