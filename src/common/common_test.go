package common

import (
	"github.com/nsf/termbox-go"
	"log"
	"testing"
	"time"

	"github.com/RussellLuo/timingwheel"
)

func TestA(t *testing.T) {
	tw := timingwheel.NewTimingWheel(time.Millisecond, 20)
	log.Println("qq")
	tw.Start()
	defer tw.Stop()
	exitC := make(chan time.Time, 1)
	tw.AfterFunc(time.Second, func() {
		log.Println("The timer fires")
	})
	go func() {
		for {
		}
	}()
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	go func() {
		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					log.Println("esc")
					exitC <- time.Now().UTC()
				case termbox.KeyArrowUp:
					log.Println("KeyArrowUp")
				}

			}
		}
	}()
	log.Println(<-exitC)

}
