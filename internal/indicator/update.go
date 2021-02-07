package indicator

import (
	"github.com/getlantern/systray"
	"github.com/jersou/timer-indicator-go/assets"
	"github.com/jersou/timer-indicator-go/internal/utils"
	"time"
)

func (indicator *Indicator) update() {
	switch indicator.state {

	case PROCESSING:
		if lastDuration := indicator.EndDate.Sub(time.Now()); lastDuration < 0 {
			indicator.done()
		} else {
			title := utils.GetDurationStr(lastDuration)
			systray.SetTitle(title)
			systray.SetTooltip(title)
		}

	case PAUSE:
		if indicator.iconState == PAUSE {
			systray.SetIcon(assets.GetWarningIcon())
			indicator.iconState = PROCESSING
		} else {
			systray.SetIcon(assets.GetPauseIcon())
			indicator.iconState = PAUSE
		}

	case DONE:
		if indicator.iconState == DONE {
			systray.SetIcon(assets.GetWarningIcon())
			indicator.iconState = PAUSE
		} else {
			systray.SetIcon(assets.GetProcessingIcon())
			indicator.iconState = DONE
		}
	}
}
