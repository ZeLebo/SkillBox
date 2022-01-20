package student

import (
	"fmt"
	"os"
)

type Student struct {
	Name string
	Age, Grade int
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s *Student) SetAge(age int) {
	s.Age = age
}

func (s *Student) SetGrade(grade int) {
	s.Grade = grade
}

func (s Student) GetName() string {
	return s.Name
}

func (s Student) GetAge() int {
	return s.Age
}

func (s Student) GetGrade() int {
	return s.Grade
}


func (s Student) PrintInfo() {
	fmt.Fprintf(os.Stdout, "Student: %s, Age: %d, Grade: %d\n", s.Name, s.Age, s.Grade)
}

func NewStudent(name string, age, grade int) *Student {
	var chel Student
	chel.Name = name
	chel.Age = age
	chel.Grade = grade
	return &chel
}

func GetStudent() map[string] *Student {
    var (
        name string;
        age, grade int
    )
    storage := make(map[string] *Student)

    size, err := fmt.Fscanf(os.Stdin, "%s %d %d", &name, &age, &grade)
    for  err == nil && size == 3 {
        storage[name] = NewStudent(name, age, grade)
        size, err = fmt.Fscanf(os.Stdin, "%s %d %d", &name, &age, &grade)
    }

    return storage
}
