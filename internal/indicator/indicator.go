package indicator

import (
	"github.com/getlantern/systray"
	"time"
)

type State int

const (
	WAITING    State = iota
	PROCESSING State = iota
	PAUSE      State = iota
	DONE       State = iota
)

type Indicator struct {
	state     State
	iconState State

	EndDate          time.Time
	remainingAtPause time.Duration

	stopChan                chan struct{}
	togglePauseChan         chan struct{}
	aggregatedDurationsChan chan int
	quitChan                chan struct{}
	customChan              chan struct{}
}

func Run() {
	systray.Run(onReady, nil)
}

func onReady() {
	indicator := Indicator{}
	indicator.createMenu()
	go indicator.eventLoop()
}
