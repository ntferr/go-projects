package main

import (
	"os"

	"github.com/gocarina/gocsv"
)

type Person struct {
	ID   string `csv:"id"`
	Name string `csv:"name"`
	Job  Job    `csv:"job"`
}

type Job struct {
	Area   string  `csv:"area"`
	Salary float32 `csv:"salary"`
}

func main() {
	jose := Person{
		ID:   "1",
		Name: "Jose ",
		Job: Job{
			Area:   "IT",
			Salary: 3200.70,
		},
	}

	people := []*Person{}

	people = append(people, &jose)

	personFile, err := os.OpenFile("person.csv", os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	defer personFile.Close()

	err = gocsv.MarshalFile(people, personFile)
	if err != nil {
		panic(err)
	}
}
