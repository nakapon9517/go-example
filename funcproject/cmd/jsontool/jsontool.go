package jsontool

import (
	"encoding/json"
	"fmt"
	data "funcproject/datastruct"
	errControl "funcproject/error"
)

// Test test
func Test() {
	// person1
	fmt.Println("\n>>>person1")
	person := data.CreatePerson(1, "Gopher", "gopher@example.org", 5, "", "golang lover")
	byteData := ConvertPersonToJSON(person)
	fmt.Println(string(byteData))

	// person2
	fmt.Println("\n>>>person2")
	jsondata := []byte(`{"id":1,"name":"Gopher","age":5}`)
	person2 := ConvertJSONToPerson(jsondata)
	fmt.Println(person2)
}

// ConvertPersonToJSON getPersonToJson
func ConvertPersonToJSON(person *data.Person) []byte {
	b, err := json.Marshal(person)
	errControl.ErrorHandler(err)
	return b
}

// ConvertJSONToPerson getJsonToPerson
func ConvertJSONToPerson(byteData []byte) data.Person {
	var persons data.Person
	err := json.Unmarshal(byteData, &persons)
	errControl.ErrorHandler(err)
	return persons
}
