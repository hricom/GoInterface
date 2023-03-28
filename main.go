package main

import (
	"fmt"

	"github.com/go_interface/users"
)

func main() {
	fmt.Println("Yanye...")

	var frank users.Cashier = users.NewEmployee("Frank")
	var robert users.Admin = users.NewEmployee("Rober")

	total := frank.CalcTotal(90, 65, 93.6)
	fmt.Println(total)

	robert.DeactivateEmployee(frank)
	fmt.Println(frank.CalcTotal(90, 65, 93.6))
}
