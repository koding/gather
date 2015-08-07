package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	EarlyExit = "klient works"
	EnvKey    = "GATHER"

	Abuse     = "abuse"
	Analytics = "analytics"

	CommonScriptPath      = "gatherers/common"
	AbuseScriptPrefix     = "gatherers/ab/"
	AnalyticsScriptPrefix = "gatherers/an/"

	DefaultKeys = []string{Abuse, Analytics}
)

type Result struct {
	Error string      `json:"error,omitempty"`
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type Results []*Result

func main() {
	envKey := os.Getenv(EnvKey)
	if envKey == "" || !contains(envKey, DefaultKeys) {
		fmt.Println(EarlyExit)
		os.Exit(0)
	}

	defer os.Unsetenv(EnvKey)

	var results = Results{}

	commonBytes, err := Asset(CommonScriptPath)
	if err != nil {
		log.Fatal(err)
	}

	var scriptPrefix string = AbuseScriptPrefix
	if envKey == Analytics {
		scriptPrefix = AnalyticsScriptPrefix
	}

	for scriptPath, _ := range _bindata {
		if runnable(scriptPath, scriptPrefix) {
			result, err := runScript(commonBytes, scriptPath)
			if err != nil {
				continue
			}

			results = append(results, result)
		}
	}

	resultBytes, _ := json.MarshalIndent(results, "", "\t")
	fmt.Println(string(resultBytes))
}

func runnable(scriptPath, scriptPrefix string) bool {
	return strings.Contains(scriptPath, scriptPrefix)
}

func runScript(commonBytes []byte, scriptPath string) (*Result, error) {
	scriptBytes, err := Asset(scriptPath)
	if err != nil {
		return nil, err
	}

	combinedBytes := fmt.Sprintf("%s\n%s", commonBytes, scriptBytes)
	resultBytes, err := exec.Command("bash", "-c", combinedBytes).Output()
	if err != nil {
		return nil, err
	}

	return encodeResult(resultBytes)
}

func encodeResult(resultBytes []byte) (*Result, error) {
	outputBuffer := bytes.NewBuffer(resultBytes)

	var result Result
	if err := json.NewDecoder(outputBuffer).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func contains(key string, collection []string) bool {
	for _, c := range collection {
		if c == key {
			return true
		}
	}

	return false
}
