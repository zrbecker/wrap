package wos

import (
	"os"

	"github.com/zrbecker/wrap"
)

func Open(name string) wrap.Result[*os.File] {
	file, err := os.Open(name)
	if err != nil {
		return wrap.Error[*os.File](err)
	}
	return wrap.OK(file)
}

func Read(file *os.File, buf []byte) wrap.Result[int] {
	n, err := file.Read(buf)
	if err != nil {
		return wrap.Error[int](err)
	}
	return wrap.OK(n)
}
