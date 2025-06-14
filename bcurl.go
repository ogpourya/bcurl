package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var firefoxHeaders = map[string]string{
	"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:139.0) Gecko/20100101 Firefox/139.0",
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Accept-Encoding":           "gzip, deflate",
	"Accept-Language":           "en-US,en;q=0.5",
	"DNT":                      "1",
	"Sec-GPC":                  "1",
	"Connection":                "keep-alive",
	"Upgrade-Insecure-Requests": "1",
	"Priority":                  "u=0, i",
}

func headerName(h string) string {
	return strings.ToLower(strings.TrimSpace(strings.SplitN(h, ":", 2)[0]))
}

func main() {
	userArgs := os.Args[1:]
	cleanArgs := []string{"--compressed", "-sS"}

	userHeaderKeys := map[string]bool{}

	for i := 0; i < len(userArgs); i++ {
		if userArgs[i] == "-H" && i+1 < len(userArgs) {
			name := headerName(userArgs[i+1])
			userHeaderKeys[name] = true
			cleanArgs = append(cleanArgs, userArgs[i], userArgs[i+1])
			i++
			continue
		}
		cleanArgs = append(cleanArgs, userArgs[i])
	}

	for k, v := range firefoxHeaders {
		if _, exists := userHeaderKeys[strings.ToLower(k)]; exists {
			continue
		}
		cleanArgs = append(cleanArgs, "-H", fmt.Sprintf("%s: %s", k, v))
	}

	cmd := exec.Command("curl", cleanArgs...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	// Print both outputs BEFORE exiting
	fmt.Print(stdout.String())
	fmt.Fprint(os.Stderr, stderr.String())

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		} else {
			os.Exit(1)
		}
	}
}
