16.1 Problem 
One of your colleagues has created a program composed of a unique main.go file:

// package-imports/application-refactor/problem/main.go
package main

import "fmt"

func main() {

    // first reservation
    customerName := "Doe"
    customerEmail := "john.doe@example.com"
    var nights uint = 12
    emailContents := getEmailContents("M", customerName, nights)
    sendEmail(emailContents, customerEmail)
    createAndSaveInvoice(customerName, nights, 145.32)
}

// send an email
func sendEmail(contents string, to string) {
    // ...
    // ...
}

// prepare email template
func getEmailContents(title string, name string, nights uint) string {
    text := "Dear %s %s,\n your room reservation for %d night(s) is confirmed. Have a nice day !"
    return fmt.Sprintf(text,
        title,
        name,
        nights)
}

// create the invoice for the reservation
func createAndSaveInvoice(name string, nights uint, price float32) {
    // ...
}

You are asked to refactor the code to improve its maintainability. Propose a new code organization :

Which packages should you create?

Should you create new directories?