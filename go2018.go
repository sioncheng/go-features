package main

import "time"

func main()  {
	year2018 := time.Date(2018, 1, 1, 0,0,0,0, time.Local)
	duration := year2018.Sub(time.Now())
	for {
		select {
		case <- time.After(duration):
			println("hello 2018!")
		}
	}
}
