package gobuffer_helper

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

// Define execCommand as a variable, which by default is exec.Command
var execCommand = exec.Command

func BufferAndPrint(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	Buffer(s)
	fmt.Printf("BUFFERED: %s\n", s)
}

func Buffer(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = execCommand("pbcopy")
	case "linux":
		cmd = execCommand("xclip", "-selection", "clipboard")
	case "windows":
		cmd = execCommand("clip")
	default:
		log.Fatalf("unsupported platform: %s", runtime.GOOS)
	}

	cmd.Stdin = strings.NewReader(s)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to copy to clipboard: %s", err)
	}
}
