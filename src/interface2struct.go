package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	Name string `json:"name"`
}

type UniversalDTO struct {
	Data interface{} `json:"data"`
	// more fields with important meta-data about the message...
}

func Interface2Struct() {
	// create a customer, add it to DTO object and marshal it
	customer := Customer{Name: "Ben"}
	dtoToSend := UniversalDTO{customer}
	byteData, _ := json.Marshal(dtoToSend)

	// unmarshal it (usually after receiving bytes from somewhere)
	receivedDTO := UniversalDTO{}
	json.Unmarshal(byteData, &receivedDTO)

	//Attempt to unmarshall our customer
	receivedCustomer := getCustomerFromDTO(receivedDTO.Data)
	fmt.Println(receivedCustomer)
}

func getCustomerFromDTO(data interface{}) Customer {
	customer := Customer{}
	bodyBytes, _ := json.Marshal(data)
	json.Unmarshal(bodyBytes, &customer)
	return customer
}
