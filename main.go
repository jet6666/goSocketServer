package main

import (
	"fmt"
	"gorock/common"
)

func main() {
	fmt.Println(" this main process ")
	name := common.GetName()
	fmt.Println(" this name " + name)
}
