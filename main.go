package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"unsafe"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	messageBox       = user32.NewProc("MessageBoxW")
	getForegroundWin = user32.NewProc("GetForegroundWindow")
)

const (
	MB_OK                = 0x00000000
	MB_ICONINFORMATION   = 0x00000040
	MB_DEFBUTTON1        = 0x00000000
	MB_APPLMODAL         = 0x00000000
	SW_HIDE              = 0
	SW_SHOWNORMAL        = 1
	SW_SHOWMINIMIZED     = 2
	SW_SHOWMAXIMIZED     = 3
	SW_SHOWNOACTIVATE    = 4
	SW_SHOW              = 5
	SW_MINIMIZE          = 6
	SW_SHOWMINNOACTIVE   = 7
	SW_SHOWNA            = 8
	SW_RESTORE           = 9
	SW_SHOWDEFAULT       = 10
	SW_FORCEMINIMIZE     = 11
	SW_MAX               = 11
	MF_BYCOMMAND         = 0x00000000
	SC_CLOSE             = 0xF060
	WM_SYSCOMMAND        = 0x0112
	SC_MINIMIZE          = 0xF020
)

func makepopup(title, text string) {
	hWnd := getForegroundWindow()
	showWindow(hWnd, SW_MINIMIZE)

	// Call MessageBoxW to display the pop-up window
	messageBox.Call(0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(MB_OK|MB_ICONINFORMATION|MB_DEFBUTTON1|MB_APPLMODAL))
}

func main() {
	// Check if enough arguments are provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run script.go <obfuscated_hex_pattern> <file_path>")
		makepopup("No arguments error!", "Incorrect syntax.\nUsage: go run script.go <obfuscated_hex_pattern> <file_path>")
		return
	}

	// Get the obfuscated hex pattern and file path from command-line arguments
	obfuscatedHexPattern := os.Args[1]
	filePath := os.Args[2]

	// Convert the obfuscated hex pattern to bytes
	hexBytes, err := hex.DecodeString(obfuscatedHexPattern)
	if err != nil {
		fmt.Printf("Invalid obfuscated hex pattern: %s\n", err)
		makepopup("Argument error!", fmt.Sprintf("Invalid obfuscated hex pattern: %s\n", err))
		return
	}

	// Read the file content
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading the file: %s\n", err)
		makepopup("File read error!", fmt.Sprintf("Error reading the file: %s\n", err))
		return
	}

	// Find the index of the hex pattern in the file content
	index := bytesIndex(data, hexBytes)
	if index == -1 {
		fmt.Println("Hex pattern not found in the file.")
		makepopup("No pattern found!", "Hex pattern not found in the file.")
		return
	}

	// Extract everything after the hex pattern into a variable
	result := data[index+len(hexBytes):]

	// Print the extracted variable
	fmt.Println(string(result))
	fmt.Println("Success!")

	// Show pop-up with the extracted result
	makepopup("Extraction Successful!", string(result))
}

// bytesIndex finds the index of a byte slice (needle) in another byte slice (haystack)
func bytesIndex(haystack []byte, needle []byte) int {
	hLen, nLen := len(haystack), len(needle)
	if nLen == 0 {
		return 0
	}
	for i := 0; i < hLen-nLen+1; i++ {
		if haystack[i] == needle[0] && bytesEqual(haystack[i:i+nLen], needle) {
			return i
		}
	}
	return -1
}

// bytesEqual checks if two byte slices are equal
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func getForegroundWindow() syscall.Handle {
	ret, _, _ := getForegroundWin.Call()
	return syscall.Handle(ret)
}

func showWindow(hWnd syscall.Handle, nCmdShow int) bool {
	ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow").Call(uintptr(hWnd), uintptr(nCmdShow))
	return ret != 0
}
