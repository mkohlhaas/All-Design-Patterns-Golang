package main

import "fmt"

/////////////// Interface ////////////////////////

type department interface {
	getNumberOfProfessors() int
	getName() string
}

//////////// Departments /////////////////////////

///////// Computer Departement ///////////////////

type computerscience struct {
	numberOfProfessors int
}

func (c *computerscience) getNumberOfProfessors() int {
	return c.numberOfProfessors
}

func (c *computerscience) getName() string {
	return "computerscience"
}

///////// Mechanical Departement /////////////////

type mechanical struct {
	numberOfProfessors int
}

func (c *mechanical) getNumberOfProfessors() int {
	return c.numberOfProfessors
}

func (c *mechanical) getName() string {
	return "mechanical"
}

///////// Null Departement ///////////////////////

type nullDepartment struct {
	numberOfProfessors int
}

func (c *nullDepartment) getNumberOfProfessors() int {
	return c.numberOfProfessors
}

func (c *nullDepartment) getName() string {
	return "nullDepartment"
}

///////// Collection of Departements /////////////

type college struct {
	departments []department
}

func (c *college) addDepartment(department string, numProf int) {
	switch department {
	case "computerscience":
		computerScienceDepartment := &computerscience{numberOfProfessors: numProf}
		c.departments = append(c.departments, computerScienceDepartment)
	case "mechanical":
		mechanicalDepartment := &mechanical{numberOfProfessors: numProf}
		c.departments = append(c.departments, mechanicalDepartment)
	}
	return
}

func (c *college) getDepartment(departmentName string) department {
	for _, department := range c.departments {
		if department.getName() == departmentName {
			return department
		}
	}
	//Return a null department if the department doesn't exist
	return &nullDepartment{}
}

func main() {
	// all departments
	departmentArray := []string{"computerscience", "mechanical", "civil", "electronics"}

	// local helper functions
	createCollege1 := func() *college {
		college := &college{}
		college.addDepartment("computerscience", 4)
		college.addDepartment("mechanical", 5)
		return college
	}

	createCollege2 := func() *college {
		college := &college{}
		college.addDepartment("computerscience", 2)
		return college
	}

	totalProfs := func(c *college) int {
		totalProfessors := 0
		for _, deparmentName := range departmentArray {
			d := c.getDepartment(deparmentName)
			totalProfessors += d.getNumberOfProfessors()
		}
		return totalProfessors
	}

	// College 1
	college1 := createCollege1()
	fmt.Printf("Total number of professors in College 1 is %d.\n", totalProfs(college1))

	// College 2
	college2 := createCollege2()
	fmt.Printf("Total number of professors in College 2 is %d.\n", totalProfs(college2))
}
