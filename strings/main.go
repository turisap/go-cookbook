package main

import (
	"fmt"
	"html"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	// format time
	fmt.Println(time.Now().Format(time.RFC822))
	fmt.Println(time.Now().Format(time.Kitchen))

	var builder strings.Builder

	builder.WriteString("kISkisk")
	builder.WriteByte(66)
	builder.WriteString(fmt.Sprint(23, 56, 77))

	fmt.Println(builder.String())

	b, err := strconv.ParseBool("true")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b)

	s := strconv.FormatFloat(2.3343, 'g', 15, 64)

	fmt.Println(s)

	replacer := strings.NewReplacer("loved", "puffed", "against", "TROLOLO", "I", "she")

	s = `
	I loved her with reason, against promise,
against peace, against hope, against happiness,
against all discouragement that could be.
I loved her with reason, with promise,
against peace, against hope, against happiness,
against all discouragement that could be.
I loved her with reason, with promise,
with peace, with hope, with happiness,
with all discouragement that could be.
`

	replaced := replacer.Replace(s)
	fmt.Println(replaced)

	// taking a substring
	str := "I am traveling down the river"
	q := "am traveling"
	i := strings.Index(str, q)

	fmt.Println(str[i : i+len(q)])

	r := strings.ContainsAny("kirill sh", "a")
	fmt.Println(r)

	sp := strings.Split(s, " ")
	fmt.Printf("%q\n\n", sp)

	sg := strings.Fields(s)
	fmt.Printf("%q\n\n", sg)

	// remove whitespace, newlines, punctuation
	sb := strings.FieldsFunc(s, func(r rune) bool {
		return unicode.IsPunct(r) || !unicode.IsLetter(r)
	})

	fmt.Println(sb)

	str = "<p>Kiskisk</p>"
	estr := html.EscapeString(str)
	fmt.Println(estr)

	ustr := html.UnescapeString(estr)
	fmt.Println(ustr)

}
