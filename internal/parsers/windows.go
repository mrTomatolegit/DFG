//go:build windows
// +build windows

//go:generate go build -ldflags "-s -w -extldflags '-static'" $GOFILE
package parsers

import (
	"syscall"
	"unsafe"
)

// Detect if windows golang executable file is running via double click or from cmd/shell terminator
// https://stackoverflow.com/questions/8610489/distinguish-if-program-runs-by-clicking-on-the-icon-typing-its-name-in-the-cons?rq=1
// https://github.com/shirou/w32/blob/master/kernel32.go
// https://github.com/kbinani/win/blob/master/kernel32.go#L3268
// win.GetConsoleProcessList(new(uint32), win.DWORD(2))
func IsDoubleClickRun() bool {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	lp := kernel32.NewProc("GetConsoleProcessList")
	if lp != nil {
		var pids [2]uint32
		var maxCount uint32 = 2
		ret, _, _ := lp.Call(uintptr(unsafe.Pointer(&pids)), uintptr(maxCount))
		if ret > 1 {
			return false
		}
	}
	return true
}
