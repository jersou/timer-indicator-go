package indicator

import (
	"github.com/getlantern/systray"
	"github.com/jersou/timer-indicator-go/assets"
	"github.com/jersou/timer-indicator-go/internal/utils"
	"time"
)

func (indicator *Indicator) createDurationItems() {
	indicator.aggregatedDurationsChan = make(chan int)
	var itemsMinutes = [...]int{1, 2, 3, 5, 10, 15, 20, 25, 30, 45, 60, 90, 120, 150, 180, 210, 240, 270, 300}
	for _, nbMin := range itemsMinutes {
		item := systray.AddMenuItem(utils.GetDurationStr(time.Duration(nbMin)*time.Minute), "")
		go func(duration int) {
			for {
				<-item.ClickedCh
				indicator.aggregatedDurationsChan <- duration
			}
		}(nbMin)
	}
}

func (indicator *Indicator) createMenu() {
	systray.SetTemplateIcon(assets.GetWaitingIcon(), assets.GetWaitingIcon())
	indicator.stopChan = systray.AddMenuItem("Stop", "").ClickedCh
	indicator.togglePauseChan = systray.AddMenuItem("Pause/Continue", "").ClickedCh
	indicator.createDurationItems()
	indicator.customChan = systray.AddMenuItem("Custom", "").ClickedCh
	systray.AddSeparator()
	indicator.quitChan = systray.AddMenuItem("Quitter", "").ClickedCh
}
