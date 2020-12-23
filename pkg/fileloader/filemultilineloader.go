package fileloader

import (
	"bufio"
	"os"
	"strings"
)

type Filereader interface {
	Open(filepath string) (bool, error)
	SetDelimeter(string)
	Close() error
}

// FileMlutilineLoader
// mlutiline file loader so as to gather objects of multiple lines
type MultilineFilereader struct {
	filepath    string
	filepointer *os.File
	scanner     *bufio.Scanner
}

// Open
// Open the filepath set int he module if possible.
func (fml *MultilineFilereader) Open() error {

	fp, err := os.Open(fml.filepath)
	//defer fml.filepointer.Close()
	if err != nil {
		return err
	}
	fml.filepointer = fp
	fml.scanner = bufio.NewScanner(fml.filepointer)
	fml.scanner.Split(bufio.ScanLines)

	return nil
}

// ReadMultiLine
// Split blocks of lines by delimieter
// Returns the text block string, eof boolean and any error raised
func (fml *MultilineFilereader) ReadMultiLineBlock() (string, bool, error) {
	lines, eof, err := fml.ReadMultiLineSlices()
	textblock := strings.Join(lines, " ")
	return textblock, eof, err

}

// ReadMultiLIneSlices
func (fml *MultilineFilereader) ReadMultiLineSlices() (lines []string, eof bool, err error) {
	for {
		if fml.scanner.Scan() {
			l := fml.scanner.Text()
			if len(l) == 0 {
				break
			} else {
				lines = append(lines, l)
			}
		} else {
			err := fml.scanner.Err()
			if err != nil {
				return lines, false, err
			}
			//is eof
			return lines, true, nil
		}
	}

	return lines, false, nil

}

// Close
// Close filepath pointer
func (fml *MultilineFilereader) Close() error {
	err := fml.filepointer.Close()
	if err != nil {
		return err
	}
	return nil
}

// NewMultilineFilereader
// Generate a new
func NewMultilineFilereader(filepath string) (MultilineFilereader, error) {
	fml := MultilineFilereader{filepath: filepath}
	err := fml.Open()
	if err != nil {
		return fml, err
	}
	return fml, nil
}
