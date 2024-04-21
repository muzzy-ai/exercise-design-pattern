package main

import "fmt"

// User interface
type User interface {
    doStuff() string
}

// UserBase struct
type UserBase struct {
    name string
}

func (u *UserBase) doStuff() string {
    return "hello, i am " + u.name
}

// Admin struct
type Admin struct {
    UserBase
}

// NewAdmin creates a new Admin instance
func NewAdmin() User {
    return &Admin{
        UserBase: UserBase{
            name: "Admin",
        },
    }
}

// Customer struct
type Customer struct {
    UserBase
}

// NewCustomer creates a new Customer instance
func NewCustomer() User {
    return &Customer{
        UserBase: UserBase{
            name: "Customer",
        },
    }
}

// CreateUser creates user based on the given name
func CreateUser(name string) User {
    if name == "Admin" {
        return NewAdmin()
    } else if name == "Customer" {
        return NewCustomer()
    }
    return nil
}

func main() {
    admin := CreateUser("Admin")
    fmt.Println(admin.doStuff())
    customer := CreateUser("Customer")
    fmt.Println(customer.doStuff())
}
