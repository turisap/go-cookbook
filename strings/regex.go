package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	re := regexp.MustCompile(`\w{5} \d{2}`)

	str := "kik mir 45 new cloud 95 4 k"

	fmt.Println(re.MatchString(str))
	fmt.Println(re.FindString(str))
	fmt.Println(re.FindAllString(str, -1))

	idxs := re.FindStringIndex(str)
	fmt.Println(idxs)
	fmt.Println(str[idxs[0]:idxs[1]])

	strT := re.ReplaceAllStringFunc(str, strings.ToUpper)
	fmt.Println(strT)
}
