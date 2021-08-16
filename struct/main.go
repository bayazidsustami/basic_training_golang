package main

import "fmt"

//declare struct
type student struct {
	name  string
	grade int
}

type person struct {
	name string
	age  int
}

type foreignStudent struct {
	country string
	grade   int
	person  // emmbed struct
}

type People = person //type alias struct

/*
nested struct
type student struct{
	person struct{
		name string
		age int
	}
	grade int
	hobbies []string
}
*/

/*
 declare struct horizontal

 type person struct{name string; age int; hobbies []string }

 use semicolon(;) for separate property
*/

/*
struct with property tag

type person struct{
	name string "tag1"
	age int "tag2"
}
*/

func main() {
	var s1 student
	s1.name = "jhon wick"
	s1.grade = 2

	var s2 = student{"ethan", 3}
	var s3 = student{name: "jason"}

	fmt.Println("Name : ", s1.name)
	fmt.Println("grade : ", s1.grade)
	fmt.Println("Name s2: ", s2.name)
	fmt.Println("Name s3 :", s3.name)

	var s4 *student = &s2 //struct with pointer
	fmt.Println("Name : ", s2.name)
	fmt.Println("Name : ", s4.name)

	s4.name = "frank"
	fmt.Println("Name : ", s2.name)
	fmt.Println("Name : ", s4.name)

	var fStudent = foreignStudent{}
	fStudent.name = "Smith"
	fStudent.age = 18
	fStudent.country = "Canada"
	fStudent.grade = 3

	fmt.Println("Foreign Student Data :")
	fmt.Println("Name \t:", fStudent.name)
	fmt.Println("Age \t:", fStudent.age)
	/*
		we can use this way to getting value of embedding struct
		or if the struct have the same property name
		so we should explicitly access property
	*/
	//fmt.Println("Ages \t:", fStudent.person.age)
	fmt.Println("Country \t:", fStudent.country)
	fmt.Println("Grade \t:", fStudent.grade)

	//assign sub struct
	var struct1 = person{name: "Jhon", age: 18}
	var struct2 = foreignStudent{person: struct1, country: "USA", grade: 3}
	fmt.Println("Foreign Student Data :")
	fmt.Println("Name \t:", struct2.name)
	fmt.Println("Age \t:", struct2.age)
	fmt.Println("Country \t:", struct2.country)
	fmt.Println("Grade \t:", struct2.grade)

	//anonymous struct

	/*
		declare a anonymous struct
		var struct{
			person
			grade int
		}
	*/

	var aStruct = struct {
		person
		grade int
	}{
		person: person{name: "kyle", age: 19},
		grade:  3,
	}
	// or wa can use this way
	/* aStruct.person = person{name: "kyle", age: 19}
	aStruct.grade = 3 */

	fmt.Println("name : ", aStruct.person.name)
	fmt.Println("age : ", aStruct.age)
	fmt.Println("grade : ", aStruct.grade)

	//combine slice & struct
	/* var allStudents = []student{
		{name: "Andreas", grade: 3},
		{name: "Bourne", grade: 2},
		{name: "Charlie", grade: 3},
	} */

	//or anonymous struct with slice
	var allStudents = []struct {
		name string
		age  int
	}{
		{name: "Andreas", age: 19},
		{name: "Bourne", age: 20},
		{name: "Charlie", age: 18},
	}
	for _, studentItem := range allStudents {
		fmt.Println(studentItem.name, "age is", studentItem.age)
	}

}
