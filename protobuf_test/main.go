package main

import (
	pb "example"
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func main() {
	test := &pb.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
		Union: &pb.Test_Name{Name: "fred"},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("marshaling error: ", err)
	}
	fmt.Println("gen byte is ", data)
	for k, v := range data {
		if v == 104 {
			data[k] = 72
		}
		if v == 102 {
			data[k] = 70
		}
	}
	if err := ioutil.WriteFile(".\\pblog", data, 0644); err != nil {
		fmt.Println("some err happend in write pb log file")
	}
	readbyte := []byte{}
	if readbyte, err = ioutil.ReadFile(".\\pblog"); err != nil {
		fmt.Println("some err happend in read pb log file")
	}
	fmt.Println("read byte is ", readbyte)

	newTest := &pb.Test{}
	err = proto.Unmarshal(readbyte, newTest)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		fmt.Printf("data mismatch %s != %s\n", test.GetLabel(), newTest.GetLabel())
	}
	// Use a type switch to determine which oneof was set.
	switch u := newTest.Union.(type) {
	case *pb.Test_Number: // u.Number contains the number.
		fmt.Println("Union is type of int", u)
	case *pb.Test_Name: // u.Name contains the string.
		fmt.Println("Union is type of string", u)
	}
}
