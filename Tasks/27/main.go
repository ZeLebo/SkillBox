package main

import "fmt"

type Student struct {
    name string
    age, grade int
}

func newStudent(name string, age, grade int) *Student{
    var chel Student
    chel.name = name
    chel.age = age
    chel.grade = grade
    return &chel
}

func main() {
    var (
        name string;
        age, grade int
    )
    storage := make(map[string] *Student)

    size, err := fmt.Scanf("%s %d %d", &name, &age, &grade)
    for  err == nil && size == 3 {
        storage[name] = newStudent(name, age, grade)
        size, err = fmt.Scanf("%s %d %d", &name, &age, &grade)
    }
    fmt.Println("Students from storage:")
    for i := range storage {
        fmt.Printf("Name: %s, Age: %d, Grade: %d\n", i, storage[i].age, storage[i].grade)
    }
}
