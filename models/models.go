package models

type Student struct {
	FirstName string
	LastName  string
	Courses   map[string]Course
}

type Course struct {
	Grades map[string][]float64
}
