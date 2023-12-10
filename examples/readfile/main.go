package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/zrbecker/wrap"
	"github.com/zrbecker/wrap/examples/readfile/wos"
)

const bufferLength = 1024

func SourceDir() wrap.Result[string] {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return wrap.Error[string](errors.New("failed to recover source directory"))
	}
	return wrap.OK(filepath.Dir(file))
}

func ReadFile(name string) (res wrap.Result[[]byte]) {
	// We require a deferred call to UpwrapHandler to safely use Unwrap
	defer wrap.UpwrapHandler(&res)

	file := wos.Open(name).Unwrap()

	buf := make([]byte, bufferLength)
	length := wos.Read(file, buf).Unwrap()

	return wrap.OK(buf[:length])
}

func main() {
	dir, err := SourceDir().UnwrapOrError()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %s\n", err)
		os.Exit(1)
	}

	data, err := ReadFile(filepath.Join(dir, "./data/message.txt")).UnwrapOrError()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("contents: %s", string(data))
}
