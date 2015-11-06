package log_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/log2"
	_ "github.com/sinusoids/gem/gem/python/api"
)

func TestLogger(t *testing.T) {
	_ = py.NewLock()

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
				t.Error("Mismatch looking for [%v], got [%v]", s, output)
				return false
			}
		}
		return true
	}

	logger := log.Sys.Module("TestLogger1")

	i := 0

	var message string

	message = fmt.Sprintf("TEST_%v", i)
	logger.Error(message)
	checkLogString([]string{"CRITICAL", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	logger.Debug(message)
	checkLogString([]string{"DEBUG", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	logger.Error(message)
	checkLogString([]string{"ERROR", "TestLogger1", message})
	i++

	/* Can't test fatal - calls os.Exit(1)
	message = fmt.Sprintf("TEST_%v", i)
	logger.Fatalf(message)
	checkLogString([]string{"FATAL", "TestLogger1", message})
	i++
	*/

	message = fmt.Sprintf("TEST_%v", i)
	logger.Info(message)
	checkLogString([]string{"INFO", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	logger.Info(message)
	checkLogString([]string{"NOTICE", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	logger.Error(message)
	checkLogString([]string{"WARNING", "TestLogger1", message})
	i++

	subLogger := logger.SubModule("TestLogger2")

	message = fmt.Sprintf("TEST_%v", i)
	subLogger.Error(message)
	checkLogString([]string{"CRITICAL", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	subLogger.Debug(message)
	checkLogString([]string{"DEBUG", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	subLogger.Error(message)
	checkLogString([]string{"ERROR", "TestLogger1", message})
	i++

	/* Can't test fatal - calls os.Exit(1)
	message = fmt.Sprintf("TEST_%v", i)
	subLogger.Fatalf(message)
	checkLogString([]string{"FATAL", "TestLogger1", message})
	i++
	*/

	message = fmt.Sprintf("TEST_%v", i)
	subLogger.Info(message)
	checkLogString([]string{"INFO", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	subLogger.Info(message)
	checkLogString([]string{"NOTICE", "TestLogger1", message})
	i++

	message = fmt.Sprintf("TEST_%v", i)
	subLogger.Error(message)
	checkLogString([]string{"WARNING", "TestLogger1", message})
	i++

}
