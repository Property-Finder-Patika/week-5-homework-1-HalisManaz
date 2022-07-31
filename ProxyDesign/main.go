package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

// Ids Lets say we have 10 different license and these license id's from one to ten
var Ids = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Create database struct
type dataBase struct {
	licenseIds []int
}

// Create check methods for database and proxy database
func (d *dataBase) check(id int) {
	if slices.Contains(Ids, id) {
		fmt.Println("ID is OK!")
	} else {
		fmt.Println("ID does not found")
	}
}

func (p *proxyDataBase) check(id int) {
	if slices.Contains(Ids, id) {
		fmt.Println("ID is OK!")
	} else {
		fmt.Println("ID does not found")
	}
}

// Create template for user and proxy database
type user struct {
	id int
}

type proxyDataBase struct {
	proxyLicenseIds *dataBase
}

// Use function for creating proxy database and users
func createProxyDataBase(Ids []int) *proxyDataBase {
	return &proxyDataBase{&dataBase{Ids}}
}

func createUser(id int) *user {
	return &user{id}
}

func main() {
	validUser := createUser(2)
	invalidUser := createUser(11)
	proxy := createProxyDataBase(Ids)
	proxy.check(validUser.id)   // ID is OK!
	proxy.check(invalidUser.id) // ID does not found
}
