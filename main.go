package main

import (
	"fmt"
	"io/ioutil"
	"log"

	complexpb "github.com/emskaplann/protobuf-example-go/src/complex"
	enumpb "github.com/emskaplann/protobuf-example-go/src/enum_example"
	simplepb "github.com/emskaplann/protobuf-example-go/src/simple"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()
	readAndWriteDemo(sm)

	smAsString := protoToJSON(sm)
	smFromString := &simplepb.SimpleMessage{}
	protoFromJSON(smAsString, smFromString)
	fmt.Println(smFromString)

	em := doEnum()
	cm := doComplex()
	fmt.Println(em, cm)
}

func doComplex() proto.Message {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "first message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   3,
				Name: "second message",
			},
			&complexpb.DummyMessage{
				Id:   10,
				Name: "third message",
			},
		},
	}

	return &cm
}

func doEnum() proto.Message {
	em := enumpb.EnumMessage{
		Id:           3152,
		DayOfTheWeek: enumpb.DayOfTheWeek_SATURDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_SUNDAY

	return &em
}

func protoToJSON(pb proto.Message) string {
	marshalOptions := protojson.MarshalOptions{
		Indent:          "  ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}

	out, err := marshalOptions.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
	}

	return string(out)
}

func protoFromJSON(json string, pb proto.Message) error {
	unmarshalOptions := protojson.UnmarshalOptions{}
	b := []byte(json)
	err := unmarshalOptions.Unmarshal(b, pb)
	if err != nil {
		log.Fatalln("Can't read proto from json", err)
		return err
	}
	return nil
}

func readAndWriteDemo(pb proto.Message) {
	writeToFile("my_message.bin", pb)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("my_message.bin", sm2)
	fmt.Println("Read the content: ", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into to the given protocol buffer struct", err2)
		return err2
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 2, 3, 7, 4},
	}

	sm.Name = "Renamed Message"

	fmt.Println("The ID is:", sm.GetId())
	fmt.Println("The Name is:", sm.GetName())
	return &sm
}
