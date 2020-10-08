package lasttime

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ExpBufio(count int, usebufio bool) {
	var writer io.Writer
	const buffsize = 65536

	if usebufio {
		bufWriter := bufio.NewWriterSize(os.Stdout, buffsize)
		defer bufWriter.Flush()
		writer = bufWriter
	} else {
		writer = os.Stdout
	}
	for i := 0; i < count; i++ {
		fmt.Fprintln(writer, strings.Repeat("x", i+1))
	}
}
