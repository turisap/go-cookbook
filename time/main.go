package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	//representZone()
	//representDuration()
	//sleep()
	//formatDates()
	parseToStruct()
}

func parseToStruct() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC822))  // "02 Jan 06 15:04 MST"
	fmt.Println(t.Format(time.RFC822Z)) // "02 Jan 06 15:04 -0700"
	fmt.Println(t.Format(time.RFC3339))
}

func formatDates() {
	t := time.Now()
	fmt.Println(t.Format(time.UnixDate))
	fmt.Println(t.Format("3:04pm"))
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.Kitchen))
	fmt.Println("\n" + strings.Repeat("@", 20))
	fmt.Println(t.Format(time.RFC822))  // "02 Jan 06 15:04 MST"
	fmt.Println(t.Format(time.RFC822Z)) // "02 Jan 06 15:04 -0700"
	fmt.Println(t.Format(time.RFC3339))
}

func sleep() {
	time.Sleep(time.Second * 3)
	t := time.Now()
	fmt.Println(t)

	fmt.Println("\n" + strings.Repeat("@", 20))
	fmt.Println(time.Now())
	fmt.Println(time.Now().Round(3))
}

func representDuration() {
	d := time.Hour * 3

	fmt.Println(d)
}

func representZone() {
	location, err := time.LoadLocation("Asia/Irkutsk")

	if err != nil {
		log.Fatal(err)
	}

	d := time.Now()

	fmt.Println("Local time in Irkutsk", d.In(location))
}

func basis() {

	d := time.Now()
	d = d.Add(40 * time.Minute)

	fmt.Println(d.Format(time.Kitchen))

	t := time.Date(2023, time.December, 21, 11, 23, 0, 0, time.UTC)

	// get month and weekday names
	fmt.Println(t.Month())
	fmt.Println(t.Weekday())
}
