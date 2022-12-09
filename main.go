package main

import (
	"flag"
	"log"
	"runtime"
	"time"
)

func main() {

	interval := flag.Int("i", 60, "set time interval")
	core := flag.Int("c", 1, "set core")
	flag.Parse()

	if *core > runtime.NumCPU() {
		log.Println("please enter less than the number of cpu cores")
		return
	}

	sig := make(chan bool)

	go Timer(interval, sig)
	for i := 0; i < *core; i++ {
		go Actuator(sig)
	}

	for {
		select {
		case <-sig:
			log.Println("end of stress test")
			return
		}
	}
}

func Timer(interval *int, sig chan bool) {
	for i := 0; i < *interval; i++ {
		time.Sleep(1 * time.Second)
		log.Println(*interval-i, "remaining...")
	}

	sig <- true
}

func Actuator(sig chan bool) {
	runtime.LockOSThread()

	for {
	}
}
