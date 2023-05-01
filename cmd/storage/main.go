package main

import (
	"fmt"

	"github.com/Andrey-VN/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("Item", st)
}
