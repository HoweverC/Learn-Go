package main

import (
	"fmt"
	"time"
)

func main() {
	ticker:=time.NewTicker(1*time.Second)
	i:=0
	for{
		<-ticker.C
		i++
		fmt.Println(i)
		if i == 3 {
			ticker.Stop()
			break
		}
	}
}
