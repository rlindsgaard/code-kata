package input

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

type JSONFlag struct {
	// The JSON string to be parsed
	Target any
}

func (f *JSONFlag) String() string {
	b, err := json.Marshal(f.Target)
	if err != nil {
		return "failed to marshal JSON"
	}
	return string(b)
}

func (f *JSONFlag) Set(value string) error {
	return json.Unmarshal([]byte(value), f.Target)
}

func (f *JSONFlag) Type() string {
	return "json"
}

func HandleStdIn(args []string) []string {
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		// do nothing
	} else {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		// remove surrounding whitespace
		jsonBlob := strings.Trim(string(stdin), "\n")
		jsonBlob = strings.Trim(jsonBlob, "\r")
		jsonBlob = strings.TrimSpace(jsonBlob)

		if jsonBlob != "" {
			args = append(args, "--inputJSON", jsonBlob)
		}
	}
	return args
}
