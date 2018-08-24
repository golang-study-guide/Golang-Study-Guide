package main

import (  
    "fmt"
)

// A struck is a user defined datatype
// it's a bit like a class in the Object Orientated Programming sense
type Employee struct {  
    firstName string 
    lastName  string
    age       int
    salary    int
}

func main() {

    // there are 2 ways to initialize, first way is
    var employeeA Employee
    // second way is to use the new builtin function, which works slightly differently as it will create a pointer. 
    employeeB := new(Employee)


    fmt.Println("EmployeeA", employeeA)
    fmt.Println("EmployeeB", employeeB)


    //creating structure using field names
    employee1 := Employee{
        firstName: "Peter",
        age:       18,
        salary:    5000,
        lastName:  "Anderson",
    }

    //creating structure without using field names
    employee2 := Employee{"Richard", "Smith", 29, 8000}

    fmt.Println("Employee 1:", employee1)
    fmt.Println("Employee 1 firstname:", employee1.firstName)
    fmt.Println("Employee 1 age:", employee1.age)
    fmt.Println("Employee 1 salary:", employee1.salary)
    fmt.Println("Employee 1 lastname", employee1.lastName)


    fmt.Println("Employee 2:", employee2)
    fmt.Println("Employee 2 firstname:", employee2.firstName)
    fmt.Println("Employee 2 age:", employee2.age)
    fmt.Println("Employee 2 salary:", employee2.salary)
    fmt.Println("Employee 2 lastname", employee2.lastName)

}
