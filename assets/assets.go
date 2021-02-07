package assets

import (
	_ "embed"
	"runtime"
)

//go:embed waiting.ico
var waitingIco []byte

//go:embed waiting.png
var waitingPng []byte

func GetWaitingIcon() []byte {
	if runtime.GOOS == "windows" {
		return waitingIco
	} else {
		return waitingPng
	}
}

//go:embed processing.ico
var processingIco []byte

//go:embed processing.png
var processingPng []byte

func GetProcessingIcon() []byte {
	if runtime.GOOS == "windows" {
		return processingIco
	} else {
		return processingPng
	}
}

//go:embed pause.ico
var pauseIco []byte

//go:embed pause.png
var pausePng []byte

func GetPauseIcon() []byte {
	if runtime.GOOS == "windows" {
		return pauseIco
	} else {
		return pausePng
	}
}

//go:embed warning.ico
var warningIco []byte

//go:embed warning.png
var warningPng []byte

func GetWarningIcon() []byte {
	if runtime.GOOS == "windows" {
		return warningIco
	} else {
		return warningPng
	}
}

//go:embed timer-indicator.oga
var TimerIndicatorOga []byte
