package main

import (
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)

// to create modules go mod init github.com/roosvelandia/test
// to use other users code  go get github.com/donvito/hellomod
// go build main.go
// delete dependencies go mod tidy
// ver el cache donde están las dependencias go mod download -json
func main() {
	hellomod.SayHello()
	hellomod2.SayHello("Roosvell")
}
