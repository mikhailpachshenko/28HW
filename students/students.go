package students

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

type Storage struct {
	database map[string]*Student
}

func (s *Storage) Get(input string) *Student {
	return s.database[input]
}

func (s *Storage) GetAll() []*Student {
	res := []*Student{}
	for _, v := range s.database {
		res = append(res, v)
	}
	return res
}

func (s *Storage) Put(input *Student) error {
	if input == nil {
		return fmt.Errorf("Ошибка при добавлении в базу данных")
	}
	if input.Name == "" || input.Age == 0 || input.Grade >= 101 {
		return fmt.Errorf("Ошибка при добавлении в базу данных")
	}

	s.database[input.Name] = input
	return nil
}

func NewStorage() *Storage {
	return &Storage{
		database: make(map[string]*Student),
	}
}
