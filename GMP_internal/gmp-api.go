package GMP_internal

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// This filters out bogus or erroneous params
func resolveParams(params map[int]string) {
	kick := []int{}
	for key, _ := range params {
		_, exists := ParamKey(key)
		if !exists {
			kick = append(kick, key)
		}
	}
	// kick'em
	for i := range kick {
		delete(params, kick[i])
	}
}

// Perform the GMP operation
func PerformActionWitnType(action string, params map[int]string, test bool) {
	// filter out invalid params
	resolveParams(params)
	params[KeyHitType] = action
	params[KeyClientID] = ClientIdentifier

	// assemble the query string
	// @see https://play.golang.org/p/GmP3n99lig
	urlValues := make(url.Values)
	for key, value := range params {
		k, _ := ParamKey(key)
		urlValues.Add(k, value)
	}
	log.Println(fmt.Sprintf("%s?%s", ENGA_ENDPOINT_SSL, urlValues.Encode()))
	if !test {
		go func() {
			_, err := http.PostForm(ENGA_ENDPOINT_SSL, urlValues)
			if err != nil {
				log.Println(err.Error())
			}
		}()
	}
}
