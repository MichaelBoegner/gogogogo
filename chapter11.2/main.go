package main

import {
	"cmboegner.com/packages/emailer/email"
	"cmboegner.com/packages/emailer/invoice"	
}

func main() {

	// first reservation
	customerName := "Doe"
	customerEmail := "john.doe@example.com"
	var nights uint = 12
	emailContents := email.GetContents("M", customerName, nights)
	email.Send(emailContents, customerEmail)
	invoice.CreateAndSave(customerName, nights, 145.32)
}

