package magazine

type Car struct {
	Name  string
	Speed int
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

type Employee struct {
	Name    string
	Salary  float64
	Address Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
