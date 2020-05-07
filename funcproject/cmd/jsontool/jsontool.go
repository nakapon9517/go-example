package jsontool

import (
	"encoding/json"
	data "funcproject/datastruct"
	errControl "funcproject/error"
)

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
