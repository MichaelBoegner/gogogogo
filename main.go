package main

import (
	"test/email"
	"test/invoice"
)

func main() {

	// first reservation
	email.Contents()
	email.Send()
	invoice.Create()
}
