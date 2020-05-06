package person

// Person struct
type Person struct {
	ID      int
	Name    string
	Email   string
	Age     int
	Address string
	memo    string
}

// CreatePerson createPerson
func CreatePerson(id int, name string, email string, age int, address string, memo string) *Person {
	return &Person{
		id,
		name,
		email,
		age,
		address,
		memo,
	}
}
