package main

import (
	"fmt"
	"github.com/dr-leee/go/source/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("Hello, world!", st)
}
