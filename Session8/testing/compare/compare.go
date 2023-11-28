package compare

import "time"

type Person struct {
	Name       string
	Age        int
	InsertDate time.Time
}

func CreatePerson(name string, age int) Person {

	return Person{
		Name:       name,
		Age:        age,
		InsertDate: time.Now(),
	}
}
