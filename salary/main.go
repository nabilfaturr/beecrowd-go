package main

import "fmt"

func main () {

	var (
		employeeNumber, employeeWorkedHour int
		employeeSalaryPerHour float64
	)

	fmt.Scanln(&employeeNumber);
	fmt.Scanln(&employeeWorkedHour);
	fmt.Scanln(&employeeSalaryPerHour);

	employeeTotalSalary := float64(employeeWorkedHour) * employeeSalaryPerHour;

	fmt.Printf("NUMBER = %d\n", employeeNumber);
	fmt.Printf("SALARY = U$ %.2f\n", employeeTotalSalary);

}