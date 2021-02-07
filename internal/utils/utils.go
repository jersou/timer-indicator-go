package utils

import (
	"bytes"
	"fmt"
	"github.com/jersou/timer-indicator-go/assets"
	"math"
	"os/exec"
	"time"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetDurationStr(duration time.Duration) string {
	switch {
	case duration < time.Minute:
		return fmt.Sprintf("%d s", int(math.Round(float64(duration)/float64(time.Second))))
	case duration < time.Hour:
		return fmt.Sprintf("%d m", int(math.Round(float64(duration)/float64(time.Minute))))
	case (duration/time.Minute)%60 == 0:
		return fmt.Sprintf("%d h", duration/time.Hour)
	default:
		return fmt.Sprintf("%d h %d", duration/time.Hour, int(math.Round(float64(duration%time.Hour)/float64(time.Minute))))
	}
}

// TODO : remove mplayer dependency
func PlayEndSound() {
	cmd := exec.Command("mplayer", "-")
	cmd.Stdin = bytes.NewReader(assets.TimerIndicatorOga)
	_ = cmd.Start()
}
