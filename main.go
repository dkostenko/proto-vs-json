package main

import (
	"encoding/json"
	"fmt"

	"github.com/gogo/protobuf/proto"
)

func main() {
	qwe := &Qwe{
		Age:  24,
		Name: "ivan",
	}

	dataJSON, _ := json.Marshal(qwe)
	dataProto, _ := proto.Marshal(qwe)

	fmt.Println("len(json) =", len(dataJSON))
	fmt.Println("len(dataProto) =", len(dataProto))
	fmt.Println(dataJSON)
	fmt.Println(dataProto)
}
