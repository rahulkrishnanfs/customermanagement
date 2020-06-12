package mapstore

import (
	"errors"
	"fmt"

	"github.com/rahulkrishnanfs/customermanagement/domain"
)

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

func (mapstore MapStore) Create(customer domain.Customer) error {

	if _, ok := mapstore.store[customer.ID]; ok {
		return errors.New("Employee with ID is already Present ")
	}
	mapstore.store[customer.ID] = customer
	fmt.Println("Employee is added to the database")
	return nil
}

func (mapstore MapStore) Update(id string, customer domain.Customer) error {
	if _, ok := mapstore.store[id]; ok {
		mapstore.store[id] = customer
		for k, v := range mapstore.store {
			fmt.Println(k, v)
		}
		return nil
	}
	return errors.New("There is no employee with the id present")
}

func (mapstore MapStore) Delete(id string) error {
	if _, ok := mapstore.store[id]; ok {
		delete(mapstore.store, id)
		for k, v := range mapstore.store {
			fmt.Println(k, v)
		}
		return nil
	}
	return errors.New("Specified Id does not exist to delete")
}

func (mapstore MapStore) GetById(id string) (domain.Customer, error) {
	if _, ok := mapstore.store[id]; ok {
		return mapstore.store[id], nil
	}
	return domain.Customer{}, errors.New("The requested ID is not present")

}

func (mapstore MapStore) GetAll() ([]domain.Customer, error) {
	customers := make([]domain.Customer, 0)
	for _, v := range mapstore.store {
		customers = append(customers, v)
	}
	return customers, nil
}
