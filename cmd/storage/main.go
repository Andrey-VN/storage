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

	fmt.Println("Item", file)
}
