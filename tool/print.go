package tool

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrint(data interface{})  {
	fmt.Println()
	var p []byte
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s \n", p)

}
