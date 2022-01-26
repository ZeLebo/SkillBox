package main

import (
    "os"
    stor "student/pkg/storage"
)

func main() {

    storage := stor.GetStudent(os.Stdin)
    for i := range storage {
        storage[i].PrintInfo()
    }
}
