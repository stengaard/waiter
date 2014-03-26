package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dustin/go-humanize"
)

var termTime int
var intTime int

func init() {
	flag.IntVar(&termTime, "term", 1, "how long to wait before exiting after receiving SIGTERM")
	flag.IntVar(&intTime, "int", 1, "how long to wait before exiting after receiving SIGINT")
}

func main() {

	flag.Parse()

	waits := make(map[os.Signal]time.Duration)
	waits[os.Interrupt] = time.Duration(intTime) * time.Second
	waits[syscall.SIGTERM] = time.Duration(termTime) * time.Second

	waitLoop(waits)

}

func waitLoop(waits map[os.Signal]time.Duration) {

	c := make(chan os.Signal, 0)
	for sig := range waits {
		signal.Notify(c, sig)
	}

	var exit <-chan time.Time
	tick := time.NewTicker(6 * time.Second)
	defer tick.Stop()
	start := time.Now()

	for {
		select {
		case sig := <-c:
			fmt.Printf("recevied %s - exiting in %s\n", sig, waits[sig])
			if exit == nil {
				exit = time.After(waits[sig])
			} else {
				fmt.Println("double recv signal - bum out")
				return
			}

		case <-exit:
			fmt.Println("dead")
			return

		case <-tick.C:
			s := humanize.Time(start)
			fmt.Printf("started %s - still haven't crashed\n", s)
		}
	}
}
