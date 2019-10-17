package codegen

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func DumpCode(dst io.Writer, src io.Reader) {
	scanner := bufio.NewScanner(src)
	lineno := 1
	for scanner.Scan() {
		fmt.Fprintf(dst, "%5d: %s\n", lineno, scanner.Text())
		lineno++
	}
}

func WriteGoCodeToFile(fn string, data []byte) error {
	formatted, err := format.Source(data)
	if err != nil {
		DumpCode(os.Stderr, bytes.NewReader(data))
		return errors.Wrap(err, `failed to format source code`)
	}
	return WriteToFile(fn, formatted)
}

func WriteToFile(fn string, data []byte) error {
	dir := filepath.Dir(fn)
	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return errors.Wrapf(err, `failed to create directory %s`, dir)
		}
	}

	f, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, `failed to open file for writing`)
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return errors.Wrap(err, `failed to write to file`)
	}

	return nil
}
