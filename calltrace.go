package calltrace

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime/debug"
	"time"
)

var traceFile *os.File

func createTraceFile() (err error) {
	err = nil
	if traceFile == nil {
		//t.Format("20060102150405"))
		traceFile, err = ioutil.TempFile("/tmp", time.Now().Format("15:04:05-"))
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
	return err
}

func Log(s string) {
	if err := createTraceFile(); err == nil {
		fmt.Fprintf(traceFile, "%s\n", s)
	}
}

func CallTrace(s string) {
	if err := createTraceFile(); err == nil {
		fmt.Fprintf(traceFile, "%s\n%s\n", s, string(debug.Stack()))
	}
}
