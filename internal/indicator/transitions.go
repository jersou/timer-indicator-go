package indicator

import (
	"github.com/getlantern/systray"
	"github.com/jersou/timer-indicator-go/assets"
	"github.com/jersou/timer-indicator-go/internal/utils"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func (indicator *Indicator) start(duration int) {
	indicator.state = PROCESSING
	systray.SetTemplateIcon(assets.GetProcessingIcon(), assets.GetProcessingIcon())
	indicator.EndDate = time.Now().Add(time.Duration(duration) * time.Minute)
}

// TODO : remove zenity dependency
func (indicator *Indicator) customStart() {
	entry, err := exec.Command("zenity", "--entry", "--text=Nb min =").Output()
	if err == nil {
		if nbMin, err := strconv.Atoi(strings.TrimSpace(string(entry))); err == nil {
			indicator.start(nbMin)
		}
	}
}

func (indicator *Indicator) stop() {
	indicator.state = WAITING
	systray.SetTemplateIcon(assets.GetWaitingIcon(), assets.GetWaitingIcon())
	systray.SetTitle("")
	systray.SetTooltip("")
}

func (indicator *Indicator) togglePause() {
	if indicator.state == PAUSE {
		indicator.state = PROCESSING
		systray.SetTemplateIcon(assets.GetProcessingIcon(), assets.GetProcessingIcon())
		indicator.EndDate = time.Now().Add(indicator.remainingAtPause)
	} else if indicator.state == PROCESSING {
		indicator.state = PAUSE
		systray.SetTemplateIcon(assets.GetPauseIcon(), assets.GetPauseIcon())
		indicator.remainingAtPause = indicator.EndDate.Sub(time.Now())
	}
}

func (indicator *Indicator) done() {
	utils.PlayEndSound()
	systray.SetTitle("DONE")
	systray.SetTooltip("DONE")
	indicator.state = DONE
}
