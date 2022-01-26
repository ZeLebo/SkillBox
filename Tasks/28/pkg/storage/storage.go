package storage

import (
    "student/pkg/student"
    "fmt"
    "os"
)

func GetStudent(fileName *os.File) map[string] *student.Student {
    var (
        name string;
        age, grade int
    )
    storage := make(map[string] *student.Student)

    size, err := fmt.Fscanf(fileName, "%s %d %d", &name, &age, &grade)
    for  err == nil && size == 3 {
        storage[name] = student.NewStudent(name, age, grade)
        size, err = fmt.Fscanf(fileName, "%s %d %d", &name, &age, &grade)
    }

    return storage
}
