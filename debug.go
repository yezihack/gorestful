package gorestful

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
)
var DefaultWriter io.Writer = os.Stdout

// set mode
func SetMode(mode string) {
	Mode = mode
}

func debugPrintMessage(httpMethod, absolutePath string, h Handle) {
	handlerName := nameOfFunction(h)
	debugPrint("%-6s %-25s --> %s\n", httpMethod, absolutePath, handlerName)
}

func debugPrint(format string, values ...interface{}) {
	if IsDebugging() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(DefaultWriter, "[Gorestful-debug] "+format, values...)
	}
}
func debugPrintWARNINGNew() {
	debugPrint(`[WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using code:	gorestful.SetMode(gorestful.ReleaseMode)

`)
}
func IsDebugging() bool {
	return Mode == ModeDebug
}

func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
