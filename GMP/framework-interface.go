package GMP

import (
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
}

//	Set the client identifier for requests
func SetClientIdentifier(identifier string) {
	internal.ClientIdentifier = identifier
}

func callInternalRequest(hitType string) {
	internal.PerformActionWitnType(hitType, paramBuffer)
	if clearsContextAfterRequest {
		ClearRequestContext()
	}
}

//	Perform Page View
func PageView() {
	callInternalRequest("pageview")
}

//	Perform Event
func Event() {
	callInternalRequest("event")
}
