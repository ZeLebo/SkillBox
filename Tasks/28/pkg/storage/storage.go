package main

import (
    "student/pkg/student"
)

func main() {
    storage := student.GetStudent()
    for i := range storage {
        storage[i].PrintInfo()
    }
}
