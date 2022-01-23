package main

import (
    stor "student/pkg/storage"
)

func main() {
        
    storage := stor.GetStudent()
    for i := range storage {
        storage[i].PrintInfo()
    }
}
