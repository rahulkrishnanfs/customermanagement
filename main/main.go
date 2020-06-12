package main

import (
	"fmt"

	"github.com/rahulkrishnanfs/customermanagement/domain"
	"github.com/rahulkrishnanfs/customermanagement/mapstore"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (controller CustomerController) Add(customer domain.Customer) {

	err := controller.store.Create(customer)

	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func (controller CustomerController) Updating(customer domain.Customer) {
	err := controller.store.Update(customer.ID, customer)

	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func (controller CustomerController) Delete(id string) {
	err := controller.store.Delete(id)
	if err != nil {
		fmt.Printf("%v", err)
	}

}

func (controller CustomerController) GetById(id string) {
	customer, err := controller.store.GetById(id)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("The customer details are ", customer)

}

func (controller CustomerController) GetAll() {
	customers, err := controller.store.GetAll()
	if err != nil {
		fmt.Println("hey something error")
	}
	for _, customer := range customers {
		fmt.Println(customer)
	}
}

func main() {

	customer1 := domain.Customer{
		ID:    "cust101",
		Name:  "rahulkrishnan R A",
		Email: "rahulkrishnanfs@gmail.com",
	}
	customer2 := domain.Customer{
		ID:    "cust102",
		Name:  "raSIL R A",
		Email: "rasil@gmail.com",
	}

	store := CustomerController{
		store: mapstore.NewMapStore(),
	}

	store.Add(customer1)
	store.Add(customer2)
	store.Add(customer1)

	customer2.Name = "Rasil R A"
	store.Updating(customer2)
	store.Delete("cus")
	store.Delete("cust102")
	store.GetAll()
	store.GetById("cust101")
}
