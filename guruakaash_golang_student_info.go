// Create a struct student:
//properties: id, firstname, lastname, fullname, dateOfBirth, age, CGPA[],
//grades[], avgCGPA, finalGrade

// CRUD on student

// Get all students
// Get students by Id

// camelcase for variable names, private properties, private methods eg. camelCase
// pascal case interfaces and structures, public properties and public methods eg. PascalCase
// snake case folder name, file name, package name, table names, column names eg.snake_case
// kabab case kabab-case Project name, Reopsitory Name eg. kebab-case
// Capitalized snake-case ABCD_EFGH for environment variables eg. CAPITALIZED_SNAKE_CASE

//go env -w GO111MODULE=off

package main

import (
	"fmt"
	"time"
	// age "github.com/bearbin/go-age"
)

const (
	TimeFormat1 = "2006-01-02"
)

type Student struct {
	id                      int
	firstname               string
	lastname                string
	fullname                string
	date_of_birth           time.Time
	age                     int
	cgpa                    []int
	grades                  []string
	final_cgpa              int
	final_grade             string
	year_of_enrollment      int
	year_of_passing         int
	no_of_years_to_graduate int
}

var students = make([]*Student, 0)

func getDOB(dob string) time.Time {
	dob_req, _ := time.Parse(TimeFormat1, dob)
	return dob_req
}
func calcGrades(cgpa []int) []string {
	var grades = make([]string, len(cgpa))
	for i := 0; i < len(cgpa); i++ {
		if cgpa[i] >= 90 {
			grades[i] = "O"
		} else if cgpa[i] >= 80 {
			grades[i] = "A"
		} else if cgpa[i] >= 70 {
			grades[i] = "B"
		} else {
			grades[i] = "C"
		}
	}
	return grades

}
func calcFinalGrade(final_cgpa int) string {
	var grade string
	if final_cgpa >= 90 {
		grade = "O"
	} else if final_cgpa >= 80 {
		grade = "A"
	} else if final_cgpa >= 70 {
		grade = "B"
	} else {
		grade = "C"
	}
	return grade
}
func calcAvgCgpa(cgpa []int) int {
	var avgCgpa int
	var sumCgpa int
	for i := 0; i < len(cgpa); i++ {
		sumCgpa += cgpa[i]
	}
	avgCgpa = sumCgpa / len(cgpa)
	return avgCgpa
}

// func calcYearOfPassing(year_of_enrollment time.Time) time.Time {
// 	// dob1 := getDOB(string (time.Now().Year()))
// 	// return (dob1.Year() - year_of_enrollment.Year())
// 	dob1 := getDOB(string(year_of_enrollment.Year() + 4))
// 	return (dob1)
// }

func calcYearsToGraduate(year_of_passing int, year_of_enrollment int) int {
	// curr_date := getDOB(string(time.Now().Year()))

	return (year_of_passing - year_of_enrollment)
}

func newStudent(firstname, lastname string, dob string, year_of_enrollment int, year_of_passing int, cgpa []int) *Student {
	var fullname string = firstname + " " + lastname
	dob1 := getDOB(dob)
	final_cgpa := calcAvgCgpa(cgpa)
	// year_of_enrollment1 := getDOB(year_of_enrollment)
	// year_of_passing1 := getDOB(year_of_passing)
	var newObjectOfStudent = &Student{
		id:                      len(students) + 1,
		firstname:               firstname,
		lastname:                lastname,
		fullname:                fullname,
		date_of_birth:           dob1,
		age:                     int(time.Now().Sub(dob1).Hours() / 24 / 365),
		cgpa:                    cgpa,
		grades:                  calcGrades(cgpa),
		final_cgpa:              final_cgpa,
		final_grade:             calcFinalGrade(final_cgpa),
		year_of_enrollment:      year_of_enrollment,
		year_of_passing:         year_of_passing,
		no_of_years_to_graduate: calcYearsToGraduate(year_of_passing, year_of_enrollment),
	}
	return newObjectOfStudent
}

func createAndAddRecord() {

	var firstname string
	var lastname string
	var date_of_birth string
	var y_o_e int
	var y_o_p int
	var n int
	fmt.Println("Enter First Name: ")
	fmt.Scan(&firstname)
	fmt.Println("Enter Last Name: ")
	fmt.Scan(&lastname)
	fmt.Println("Enter Date of birth: ")
	fmt.Scan(&date_of_birth)
	fmt.Println("Enter year of enrollment: ")
	fmt.Scan(&y_o_e)
	fmt.Println("Enter year of passing: ")
	fmt.Scan(&y_o_p)
	fmt.Println("Enter number of subjects: ")
	fmt.Scan(&n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// student1 := newStudent("Rahul", "Kumar", "2001-10-10", 2019, 2023, arr)

	student1 := newStudent(firstname, lastname, date_of_birth, y_o_e, y_o_p, arr)
	fmt.Println(student1.fullname)
	students = append(students, student1)
	fmt.Println(len(students))

	// fmt.Println()

}

func readStudentById() {
	var id int
	fmt.Println("Enter student id: ")
	fmt.Scan(&id)
	for i := 0; i < len(students); i++ {
		if students[i].id == id {
			fmt.Println("Id: ", students[i].id)
			fmt.Println("First Name: ", students[i].firstname)
			fmt.Println("Last Name: ", students[i].lastname)
			fmt.Println("Full Name: ", students[i].fullname)
			fmt.Println("Date of Birth: ", students[i].date_of_birth)
			fmt.Println("Age: ", students[i].age)
			fmt.Println("CGPA: ", students[i].cgpa)
			fmt.Println("Grades: ", students[i].grades)
			fmt.Println("Final CGPA: ", students[i].final_cgpa)
			fmt.Println("Final Grade: ", students[i].final_grade)
			fmt.Println("Year of enrollment: ", students[i].year_of_enrollment)
			fmt.Println("Year of passing: ", students[i].year_of_passing)
			fmt.Println("Number of years to Graduate: ", students[i].no_of_years_to_graduate)
			break
		}
	}
}

func readAllStudents() {
	for i := 0; i < len(students); i++ {

		fmt.Println()
		fmt.Println("Id: ", students[i].id)
		fmt.Println("First Name: ", students[i].firstname)
		fmt.Println("Last Name: ", students[i].lastname)
		fmt.Println("Full Name: ", students[i].fullname)
		fmt.Println("Date of Birth: ", students[i].date_of_birth)
		fmt.Println("Age: ", students[i].age)
		fmt.Println("CGPA: ", students[i].cgpa)
		fmt.Println("Grades: ", students[i].grades)
		fmt.Println("Final CGPA: ", students[i].final_cgpa)
		fmt.Println("Final Grade: ", students[i].final_grade)
		fmt.Println("Year of enrollment: ", students[i].year_of_enrollment)
		fmt.Println("Year of passing: ", students[i].year_of_passing)
		fmt.Println("Number of years to Graduate: ", students[i].no_of_years_to_graduate)
		fmt.Println()

	}
}
func deleteStudentById() {
	var id int
	fmt.Println("Enter student id: ")
	fmt.Scan(&id)

	for i := 0; i < len(students); i++ {
		if students[i].id == id {

			students[i].id = students[len(students)-1].id
			students[i].firstname = students[len(students)-1].firstname
			students[i].lastname = students[len(students)-1].lastname
			students[i].fullname = students[len(students)-1].fullname
			students[i].date_of_birth = students[len(students)-1].date_of_birth
			students[i].age = students[len(students)-1].age
			students[i].cgpa = students[len(students)-1].cgpa
			students[i].grades = students[len(students)-1].grades
			students[i].final_cgpa = students[len(students)-1].final_cgpa
			students[i].final_grade = students[len(students)-1].final_grade
			students[i].year_of_enrollment = students[len(students)-1].year_of_enrollment
			students[i].year_of_passing = students[len(students)-1].year_of_passing
			students[i].no_of_years_to_graduate = students[len(students)-1].no_of_years_to_graduate
			students = students[:len(students)-1]
			break
		}
	}
}
func updateStudentByIdInner(id int, field_name string, input_value interface{}) {
	fmt.Println("DOING")
	fmt.Println(input_value)

	for i := 0; i < len(students); i++ {
		if students[i].id == id {
			switch field_name {
			case "firstname":
				input_value_temp := input_value.(string)
				students[i].firstname = input_value_temp
				students[i].fullname = students[i].firstname + " " + students[i].lastname
			case "lastname":
				input_value_temp := input_value.(string)
				students[i].lastname = input_value_temp
				students[i].fullname = students[i].firstname + " " + students[i].lastname

			case "dob":
				input_value_temp := input_value.(string)
				students[i].date_of_birth = getDOB(input_value_temp)
				students[i].age = int(time.Now().Sub(students[i].date_of_birth).Hours() / 24 / 365)
			case "enrollment_year":
				input_value_temp := input_value.(int)
				students[i].year_of_enrollment = input_value_temp
				students[i].no_of_years_to_graduate = calcYearsToGraduate(students[i].year_of_passing, students[i].year_of_enrollment)
			case "passing_year":
				input_value_temp := input_value.(int)
				students[i].year_of_passing = input_value_temp
				students[i].no_of_years_to_graduate = calcYearsToGraduate(students[i].year_of_passing, students[i].year_of_enrollment)
			}
		}
	}

}

func updateStudentById() {

	var id int
	var field_name string

	fmt.Println("Enter student id: ")
	fmt.Scan(&id)
	fmt.Println("Enter field to update: ")
	fmt.Scan(&field_name)
	if field_name == "firstname" || field_name == "lastname" || field_name == "dob" {
		var input_value string
		fmt.Println("Enter value to update: ")
		fmt.Scan(&input_value)

		updateStudentByIdInner(id, field_name, input_value)
	} else {
		var input_value int
		fmt.Println("Enter value to update: ")
		fmt.Scan(&input_value)

		updateStudentByIdInner(id, field_name, input_value)
	}
}
func main() {

	for i := 0; i < 1; {
		var choice int
		fmt.Println("Options Menu : ")
		fmt.Println("1. Create a record")
		fmt.Println("2. Read student by Id")
		fmt.Println("3. Read all students")
		fmt.Println("4. Update a record")
		fmt.Println("5. Delete a record")
		fmt.Println("6. Exit")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			createAndAddRecord()
		case 2:
			readStudentById()
		case 3:
			readAllStudents()
		case 4:
			updateStudentById()
		case 5:
			deleteStudentById()
		case 6:
			i++

		}
	}

}
