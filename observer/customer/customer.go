package customer

import "fmt"

//Customer struct
type Customer struct {
	ID string
}

//Update sends the news to the customer
func (c *Customer) Update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.ID, itemName)
}

//GetID returns the customer ID
func (c *Customer) GetID() string {
	return c.ID
}
