package workshop

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

// assert should print out nicely information on assertion failure which line and file
func assert(o bool) {
	if !o {
		fmt.Printf("\n%c[35m%s%c[0m\n\n", 27, getRecentLine(), 27)
		os.Exit(1)
	}
}

// getRecentLine should print out Information on Assertion failure
func getRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
