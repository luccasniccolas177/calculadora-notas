package services

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/luccasniccolas177/calculadora-notas/models"
)

func AddCourse(s *models.Student, name string) error {
	if _, ok := s.Courses[name]; ok {
		return errors.New("el nombre del curso ingresado ya existe")
	}

	s.Courses[name] = models.Course{
		Grades: make(map[string][]float64),
	}

	return nil
}

func Addgrade(s *models.Student, course string) error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Nombre: ")
	var name string
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Print("Nota(s): ")
	grades := []float64{}
	if scanner.Scan() {
		gradesContainer := strings.Fields(scanner.Text())

		for _, gradeString := range gradesContainer {
			grade, err := strconv.ParseFloat(gradeString, 32)
			if err != nil {
				return err
			}

			grades = append(grades, grade)
		}
	}

	s.Courses[course].Grades[name] = grades

	return nil
}

func PrintCourses(s *models.Student) {
	fmt.Printf("Estudiante: %s %s\n", s.FirstName, s.LastName)
	fmt.Println("Cursos y sus evaluaciones:")

	for name, course := range s.Courses {
		fmt.Printf("\nCurso: %s\n", name)

		// Si el curso no tiene evaluaciones, indicarlo
		if len(course.Grades) == 0 {
			fmt.Println("  No hay evaluaciones registradas.")
			continue
		}

		for gradeName, grades := range course.Grades {
			fmt.Printf("  Evaluación: %s\n", gradeName)

			if len(grades) == 0 {
				fmt.Println("    No hay notas registradas.")
				continue
			}

			// Mostrar las notas con formato claro
			fmt.Print("    Notas: ")
			for _, grade := range grades {
				fmt.Printf("%.2f ", grade)
			}
			fmt.Println()
		}
	}
}
