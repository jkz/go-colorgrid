package keydown

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

// Allow any kind of operation
type Op interface{}

type Key uint16

type Controller struct {
	Input chan Key
	Stop  chan bool
}

func (c *Controller) getInput() {
	//timer := time.After(timeout * time.Second)
	for {
		select {
		case <-c.Stop:
			fmt.Println("STOP")
			return
			/*
				case <-timer:
					fmt.Println("TIMEOUT")
					return
			*/
		default:
			fmt.Print("!")
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				if ev.Ch != 0 {
					c.Input <- Key(ev.Ch)
				} else {
					c.Input <- Key(ev.Key)
				}
			case termbox.EventResize:
				fmt.Printf("?")
			case termbox.EventError:
				fmt.Printf("E")
			}
		}
	}
}

var timeout time.Duration = 15

func (c *Controller) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	/*
		defer func() {
			if r := recover(); r != nil {
				c.Stop <- true
			}
		}()
	*/
	defer termbox.Close()
	c.getInput()
}

func NewController() Controller {
	return Controller{make(chan Key, 10), make(chan bool, 10)}
}
