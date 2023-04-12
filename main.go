package main

import (
	"io/ioutil"
	"log"
	"protobuf-lesson/pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Suzuki",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	// serialize data
	// from go data to proto data
	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Cant't serialize", err)
	}

	// write serialized data as file
	if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
		log.Fatalln("Can't write", err)
	}
}
