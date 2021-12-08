package unboot

import (
	"fmt"
	"os/exec"
)

const (
	shutdownEvent = "aevtrsdn"
	restartEvent  = "aevtrrst"
)

func unboot(loginwindowEvent string) error {
	osascript := fmt.Sprintf(
		"tell app \"loginwindow\" to «event %s»",
		loginwindowEvent,
	)

	err := exec.Command("osascript", "-e", osascript).Run()
	if err != nil {
		return err
	}

	return nil
}

// Shutdown shuts down macOS.
func Shutdown() error {
	return unboot(shutdownEvent)
}

// Restart restarts macOS.
func Restart() error {
	return unboot(restartEvent)
}
