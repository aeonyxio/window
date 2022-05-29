package window

import (
	"syscall"
	"unsafe"

	"github.com/aeonyxio/math2d"
	"github.com/lxn/win"
)

func GetWindowByName(name string) uintptr {
	return uintptr(win.FindWindow(nil, syscall.StringToUTF16Ptr(name)))
}

func EnsureWindowActive(id uintptr) {
	hwnd := win.HWND(id)
	win.ShowWindow(hwnd, win.SW_RESTORE)
	win.SetForegroundWindow(hwnd)
	win.SetActiveWindow(hwnd)
}

func GetWindowRect(id uintptr) math2d.Rectangle {
	rect := win.RECT{}
	hwnd := win.HWND(id)
	win.GetWindowRect(hwnd, &rect)
	return math2d.Rectangle{
		X: int(rect.Left),
		Y: int(rect.Top),
		Width: int(rect.Right - rect.Left),
		Height: int(rect.Bottom - rect.Top),
	}
}

// GetHWND get foreground window hwnd
func GetHWND() win.HWND {
	hwnd := win.GetForegroundWindow()
	return hwnd
}

// SendInput send n input event
func SendInput(nInputs uint32, pInputs unsafe.Pointer, cbSize int32) uint32 {
	return win.SendInput(nInputs, pInputs, cbSize)
}

// SendMsg send a message with hwnd
func SendMsg(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	return win.SendMessage(hwnd, msg, wParam, lParam)
}

// SetActiveWindow set window active with hwnd
func SetActiveWindow(hwnd win.HWND) win.HWND {
	return win.SetActiveWindow(hwnd)
}

// SetFocus set window focus with hwnd
func SetFocus(hwnd win.HWND) win.HWND {
	return win.SetFocus(hwnd)
}

// GetMian get the main display hwnd
func GetMain() win.HWND {
	return win.GetActiveWindow()
}

// GetMianId get the main display id
func GetMainId() int {
	return int(GetMain())
}
