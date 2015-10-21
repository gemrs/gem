package log_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	_ "gem"
	"gem/log"
	_ "gem/python"
)

func TestLogger(t *testing.T) {
	fmt.Printf("starting\n")

	var stdoutBuf bytes.Buffer

	getLastOutput := func() string {
		defer stdoutBuf.Reset()
		return stdoutBuf.String()
	}

	log.SetBackend(&stdoutBuf)

	checkLogString := func(strs []string) bool {
		output := getLastOutput()
		for _, s := range strs {
			if !strings.Contains(output, s) {
				t.Errorf("Mismatch looking for [%v], got [%v]", s, output)
				return false
			}
		}
		return true
	}

	fmt.Printf("ok\n")

	logger := log.Sys.Module("TestLogger1")

	fmt.Printf("ok2\n")

	i := 0

	var message string

	message = fmt.Sprintf("TEST_%v", i)
	logger.Criticalf(message)
	checkLogString([]string{"CRITICAL", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	logger.Debugf(message)
	checkLogString([]string{"DEBUG", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	logger.Errorf(message)
	checkLogString([]string{"ERROR", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	logger.Fatalf(message)
	checkLogString([]string{"FATAL", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	logger.Infof(message)
	checkLogString([]string{"INFO", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	logger.Noticef(message)
	checkLogString([]string{"NOTICE", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	logger.Warningf(message)
	checkLogString([]string{"WARNING", "TestLogger1", message})
	i++
}
