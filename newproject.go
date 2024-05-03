package main

import (
	"encoding/json"
	"fmt"

	"strconv"

	"github.com/davecgh/go-spew/spew"
)

type Character struct {
	Str int
	Dex int
	Sta int
}

type ActionResult struct {
	success bool
	msg     string
}

func doif(b bool, f1, f2 func()) string {
	switch {
	case b:
		return "asdasd"
	case !b:
		return "buuuu"
	}
	return `no match`
}

func main() {
	C1 := Character{Str: 10, Dex: 10, Sta: 10}
	PrintfChar(C1)
	MarshallChar2json(C1)
	SpewChar(C1)

	actionResult := DoAttack(C1, 12)
	fmt.Printf("%+v\n", actionResult)
}

func DoAttack(C1 Character, i int) ActionResult {
	_AR := ActionResult{
		success: (C1.Dex >= i),
	}
	_AR.msg = strconv.FormatBool(_AR.success) + "_rajraj"
	return _AR
}

func PrintfChar(C Character) {
	//fmt.Println(quote.Go())
	fmt.Printf("%+v\n", C)
}

func MarshallChar2json(C Character) {
	str, _ := json.MarshalIndent(C, "", "\t")
	fmt.Println(string(str))
}

func SpewChar(C Character) {
	spew.Dump(C)
}
