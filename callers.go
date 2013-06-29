package trace

import (
	"fmt"
	"runtime"
	"strings"
)

// Callers is a helper function to get a stack trace from within a
// trace listener function.  The result is a list of strings, each
// giving a function name and line number.  The first string
// corresponds to the call of trace.T(), the last string corresponds
// to the program's main function.  If Callers is called from outside
// a trace listener, nil is returned.
func Callers() []string {
	res := []string{}

	callToTSeen := false
	for i := 2; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		} else if !callToTSeen {
			if strings.HasSuffix(file, "github.com/seehuhn/trace/trace.go") {
				callToTSeen = true
			}
			continue
		} else if strings.HasSuffix(file, "src/pkg/runtime/proc.c") {
			break
		}
		res = append(res, fmt.Sprintf("%s:%d", file, line))
	}
	return res
}