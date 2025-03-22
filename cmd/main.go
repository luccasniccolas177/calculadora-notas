package main

import (
	"github.com/luccasniccolas177/calculadora-notas/internal/services"
	"github.com/luccasniccolas177/calculadora-notas/models"
)

func main() {
	lucas := &models.Student{
		FirstName: "Lucas",
		LastName:  "Araya",
		Courses:   make(map[string]models.Course),
	}

	services.AddCourse(lucas, "Evaluación de Proyectos TICS")
	services.AddCourse(lucas, "Data Science")

	lucas.Courses["Evaluación de Proyectos TICS"].Grades["Controles"] = []float64{70.0, 0.0, 0.0, 0.0}
	lucas.Courses["Evaluación de Proyectos TICS"].Grades["Solemnes"] = []float64{0.0, 0.0, 0.0, 0.0}
	lucas.Courses["Data Science"].Grades["Solemnes"] = []float64{0.0, 0.0, 0.0, 0.0}
	//lucas.Courses["Procesamiento de Imagenes"].Grades["Tareas"] = []float32{0.0, 0.0, 0.0, 0.0}

	//services.PrintCourses(lucas)
	//services.Addgrade(lucas, "Data Science")
	services.PrintCourses(lucas)
}
