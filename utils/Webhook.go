package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func PrintWebhook(dto interface{}) {
	fmt.Println("====== webhook body =====")
	prettyString, _ := json.MarshalIndent(dto, "", "  ")
	fmt.Printf("%s: %s\n", reflect.TypeOf(dto).Name(), prettyString)
	fmt.Println("=========================")
}
