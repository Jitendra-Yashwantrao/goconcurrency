package main

import (
	"fmt"
	"goconcurrency/welcome"
	"strings"

	"github.com/pborman/uuid"
)

func main() {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuidWithHyphen, "  ", uuid)
	fmt.Println(welcome.Welcome())

}
