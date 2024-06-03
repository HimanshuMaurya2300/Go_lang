package main

import "fmt"

func main() {

	studentGrades := make(map[string]int)
	studentGrades["Himanshu"] = 34
	studentGrades["Maurya"] = 56
	studentGrades["Charlie"] = 6

	fmt.Println(studentGrades)

	fmt.Println("Marks of Himanshu is : ", studentGrades["Himanshu"])
	studentGrades["Himanshu"] = 99
	fmt.Println("New marks of Himanshu is : ", studentGrades["Himanshu"])

	delete(studentGrades, "Himanshu")

	fmt.Println("Marks of Himanshu is : ", studentGrades["Himanshu"])
	fmt.Println(studentGrades)

	grades, exist := studentGrades["David"]

	fmt.Println("Grade of David is : ", grades)
	fmt.Println("David exist : ", exist)
	fmt.Println("Marks of David is : ", studentGrades["David"])

	Grades, Exist := studentGrades["Maurya"]

	fmt.Println("Grade of Maurya is : ", Grades)
	fmt.Println("Maurya exist : ", Exist)
	fmt.Println("Marks of Maurya is : ", studentGrades["Maurya"])

	for key, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", key, value)
	}

	person := map[string]int{
		"Himanshu": 34,
		"Maurya":   56,
		"Charlie":  6,
	}
	fmt.Println(person)
}
