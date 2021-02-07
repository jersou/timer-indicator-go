package indicator

import (
	"github.com/getlantern/systray"
	"time"
)

func (indicator *Indicator) getTickDuration() time.Duration {
	switch {
	case indicator.state == WAITING:
		return time.Hour
	case indicator.state == PROCESSING && indicator.EndDate.After(time.Now().Add(2*time.Minute)):
		return ((indicator.EndDate.Sub(time.Now()) - 30*time.Second) % time.Minute) + time.Second
	case indicator.state == PAUSE || indicator.state == DONE:
		return 200 * time.Millisecond
	default:
		return time.Second
	}
}

func (indicator *Indicator) eventLoop() {
	for {
		select {
		case duration := <-indicator.aggregatedDurationsChan:
			indicator.start(duration)
		case <-indicator.stopChan:
			indicator.stop()
		case <-indicator.togglePauseChan:
			indicator.togglePause()
		case <-indicator.customChan:
			indicator.customStart()
		case <-indicator.quitChan:
			systray.Quit()
		case <-time.After(indicator.getTickDuration()):
			// wait for event or TickDuration
		}
		indicator.update()
	}
}
