package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	VK_LMENU  = 0xA4 // Left Alt key
	VK_RMENU  = 0xA5 // Right Alt key
	VK_Q      = 0x51 // Q key
)

var (
	user32                = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

func GetAsyncKeyState(key int) bool {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(key))
	return (ret & 0x8000) != 0
}

func main() {
	fmt.Println("Alt+Q Key Monitor")
	fmt.Println("Press Alt and then press 'q' multiple times")
	fmt.Println("Release Alt to see the count of 'q' presses")
	fmt.Println("Press Ctrl+C to exit")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	altPressed := false
	qPressed := false
	qCount := 0

	go func() {
		for {
			altDown := GetAsyncKeyState(VK_LMENU) || GetAsyncKeyState(VK_RMENU)
			
			qDown := GetAsyncKeyState(VK_Q)
			
			if altDown && !altPressed {
				altPressed = true
				qCount = 0
				fmt.Println("Alt pressed - tracking 'q' presses")
			} else if !altDown && altPressed {
				fmt.Printf("Alt released - 'q' was pressed %d times while Alt was held\n", qCount)
				altPressed = false
			}
			
			if altPressed {
				if qDown && !qPressed {
					qCount++
					fmt.Printf("'q' pressed while Alt is held (count: %d)\n", qCount)
				}
			}
			
			qPressed = qDown
			
			time.Sleep(20 * time.Millisecond)
		}
	}()

	<-sig
	fmt.Println("\nExiting program...")
}