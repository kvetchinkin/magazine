package magazine

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Employee struct {
	Name    string
	Salary  float64
	Address Address
}
