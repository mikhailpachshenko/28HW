package main

import (
	"fmt"
	"io"
	ss "students/students"
)

func main() {
	studentMap := ss.NewStorage()

	for {
		s := ss.Student{}
		fmt.Println("Введите имя, возраст и оценку студента:")

		_, err := fmt.Scan(&s.Name, &s.Age, &s.Grade)
		if err == io.EOF {
			fmt.Print("Список студентов:\n")
			break
		}

		if err := studentMap.Put(&s); err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
	}

	count := 1
	for _, student := range studentMap.GetAll() {
		fmt.Print(count, ": ", student.Name, " ", student.Age, " ", student.Grade, "\n")
		count++
	}
}
