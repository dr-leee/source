package main

import (
	"fmt"
	"github.com/dr-leee/source/source/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("newfile.txt", []byte("Hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(st.files)

	fmt.Println("File is restored", restoredFile)
}
