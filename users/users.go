package users

import (
	"fmt"
	"math/rand"
)

type Users struct {
	Id   int
	Name string
}

type Employee struct {
	Users
	Active bool
}

type Cashier interface {
	CalcTotal(item ...float32) float32
	deactivate()
}

type Admin interface {
	DeactivateEmployee(c Cashier)
}

func NewEmployee(name string) *Employee {
	return &Employee{
		Users: Users{
			Id:   rand.Intn(1000),
			Name: name},
		Active: true,
	}
}

func (e *Employee) CalcTotal(item ...float32) float32 {
	if !e.Active {
		fmt.Println("Cashier deactivate")
		return 0
	}

	var sum float32

	for _, i := range item {
		sum += i
	}

	return sum * 1.15
}

func (e *Employee) deactivate() {
	e.Active = false
}

func (e *Employee) DeactivateEmployee(c Cashier) {
	c.deactivate()
}
