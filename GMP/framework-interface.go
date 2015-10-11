package GMP

import (
	"fmt"
	"log"

	internal "github.com/evilnode/cposc-go/GMP_internal"
)

//	A buffer of parameters for the next request
var paramBuffer map[int]string = make(map[int]string)
var defaultParamBuffer map[int]string = make(map[int]string)
var clearsContextAfterRequest bool = true

//  Wipe param buffer and reset defaults
func ClearRequestContext() {
	for key, _ := range paramBuffer {
		delete(paramBuffer, key)
	}
	for key, value := range defaultParamBuffer {
		paramBuffer[key] = value
	}
}

//
func SetClearsContextAfterRequest(clears bool) {
	clearsContextAfterRequest = clears
}

// 	Set a request parameter
func SetParam(key int, value string) {
	paramBuffer[key] = value
}

func SetDefaultParam(key int, value string) {
	defaultParamBuffer[key] = value
	log.Println(fmt.Sprintf("SETTING PARAM: %d => %s", key, value))
}

//	Set the client identifier for requests
func SetClientIdentifier(identifier string) {
	internal.ClientIdentifier = identifier
	log.Println(fmt.Sprintf("CLIENT IDENTIFIER: %s", identifier))
}

func callInternalRequest(hitType string, test bool) {
	internal.PerformActionWitnType(hitType, paramBuffer, test)
	if clearsContextAfterRequest {
		ClearRequestContext()
	}
}

//	Perform Page View
func PageView(test bool) {
	callInternalRequest("pageview", test)
}

//	Perform Event
func Event(test bool) {
	callInternalRequest("event", test)
}
