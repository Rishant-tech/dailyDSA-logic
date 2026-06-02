package main 

/*
Order
1 100
2 200
3 150
 
 
Payments
1 100
2 50
2 150
4 80
 
Expected Output
 
1 MATCHED
2 MATCHED
3 MISSING_PAYMENT
4 MISSING_ORDER9
*/

import (
	"fmt"
)

type Orders struct{
	id int
	amount int
}

type Payments struct {
	id int
	amount int
}

func main() {

	orders := []Orders{
		{1,50}, //t - 1
		{2,200},
		{3,150},
	}

	payments := []Payments{
		{1,100},
		{2,50},
		{2,150},
		{4,80},
	}


	//handle missing payment and matched payment vlaues for order id
	for _, order := range orders {
		totalAmount := 0

		for _, payment := range payments{
			if payment.id == order.id{
				totalAmount += payment.amount
			}
		}

		if totalAmount == order.amount {
			fmt.Println(order.id,"MATCHED")
		} else {
			fmt.Println(order.id,"MISSING_PAYMENT")
		}
	}

	//handle missing order id, id payment exists for that
	for  _, payment := range payments {

		found := false //flag to keep track of payment correspoinding to order id

		for  _, order := range orders {
			if order.id == payment.id {
				found = true
				break
			} 
		}
		
		if !found {
			fmt.Println(payment.id, "MISSING_ORDER")
		}
	}

}