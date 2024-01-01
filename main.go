package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"golang.design/x/clipboard"
)

func main() {
	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	jwtdata := clipboard.Read(clipboard.FmtText)

	token, err := jwt.Parse(string(jwtdata), nil)
	if err != nil && !errors.Is(err, jwt.ErrTokenUnverifiable) {
		panic(err)
	}

	if token != nil {
		data, err := json.MarshalIndent(token.Claims, "", "    ")
		if err != nil {
			panic("error marshalling")
		}

		fmt.Printf("%s\n", string(data))
	}

}
