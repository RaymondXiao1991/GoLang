package main

import (
	"fmt"
	"regexp"
)

func TestRegExp() {
	text := `瑞尔大厦29栋 1201室`
	matched, _ := regexp.MatchString(`[\PP]+ +[0-9]+`, text)
	if matched {
		reg := regexp.MustCompile(` +`)
		src := []byte(text)
		reduces := reg.FindIndex(src)
		fmt.Printf("%s\n", src[0:reduces[0]])
		fmt.Printf("%s\n", src[reduces[1]:])
	} else {
		fmt.Printf("Do not matched!")
	}
}
