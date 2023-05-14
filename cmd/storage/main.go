package main

import (
	"fmt"
	"log"

	"github.com/Andrey-VN/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("text.txt", []byte("hello"))

	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetById(file.ID)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(st.)

	fmt.Println("Item", restoredFile.String())
}
