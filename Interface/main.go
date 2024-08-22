package main

import "fmt"

type payment struct {
	gateway paymenter
}

type paymenter interface {
	pay(amount float32)
}

type paypal struct {
}

func (p paypal) pay(amount float32) {
	fmt.Printf("Paying via PayPal: %.2f\n", amount)
}

func main() {
	paypalGw := paypal{}
	paymentGw := payment{
		gateway: paypalGw,
	}

	paymentGw.gateway.pay(100.00)

}
