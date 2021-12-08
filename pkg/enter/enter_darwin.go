package enter

// #cgo darwin LDFLAGS: -framework Carbon
// #include <Carbon/Carbon.h>
import "C"

import (
	"fmt"
)

const returnKey = C.kVK_Return

func createKey(key int) (C.CGEventRef, C.CGEventRef, error) {
	down := C.CGEventCreateKeyboardEvent(0, C.CGKeyCode(key), true)
	if down == 0 {
		return 0, 0, fmt.Errorf("failed to create key %d", key)
	}

	up := C.CGEventCreateKeyboardEvent(0, C.CGKeyCode(key), false)
	if up == 0 {
		return 0, 0, fmt.Errorf("failed to create key %d", key)
	}

	return down, up, nil
}

func tap(key int) error {
	down, up, err := createKey(key)
	if err != nil {
		return err
	}

	loc := C.CGEventTapLocation(C.kCGHIDEventTap)
	C.CGEventPost(loc, down)
	C.CGEventPost(loc, up)

	return nil
}

func Tap() error {
	tap(returnKey)
	return nil
}
