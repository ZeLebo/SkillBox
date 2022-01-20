package main

import (
    "student/pkg/student"
    "fmt"
    "os"
)

func GetStudent() map[string] *student.Student {
    var (
        name string;
        age, grade int
    )
    storage := make(map[string] *student.Student)

    size, err := fmt.Fscanf(os.Stdin, "%s %d %d", &name, &age, &grade)
    for  err == nil && size == 3 {
        storage[name] = student.NewStudent(name, age, grade)
        size, err = fmt.Fscanf(os.Stdin, "%s %d %d", &name, &age, &grade)
    }

    return storage
}

func main() {
    storage := GetStudent()
    for i := range storage {
        storage[i].PrintInfo()
    }
}
