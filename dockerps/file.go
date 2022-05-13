// Package dockerps provides a Git commit message template fuzzy search feature.
// Optionally, it provides a fuzzy search for the history of your commit messages.
package dockerps

import (
	"bufio"
	"github.com/atotto/clipboard"
	"github.com/ktr0731/go-fuzzyfinder"
	"strings"
)

var (
	scannerScan       func(scanner *bufio.Scanner) bool
	scannerText       func(scanner *bufio.Scanner) string
	fuzzyfinderFind   func(slice interface{}, itemFunc func(i int) string, opts ...fuzzyfinder.Option) (int, error)
	dockerPs          func() (string, error)
	stringsSplit      func(s, sep string) []string
	clipboardWriteAll func(text string) error
)

func init() {
	scannerScan = func(scanner *bufio.Scanner) bool {
		return scanner.Scan()
	}
	scannerText = func(scanner *bufio.Scanner) string {
		return scanner.Text()
	}
	dockerPs = _dockerPs
	fuzzyfinderFind = fuzzyfinder.Find
	stringsSplit = strings.Split
	clipboardWriteAll = clipboard.WriteAll
}

// Copy wraps Git Copy.
// You can perform a fuzzy search from a message template and commit the result.
func Copy() (err error) {
	psReturn, err := dockerPs()
	if err != nil {
		return err
	}

	var psArray []string
	scanner := bufio.NewScanner(strings.NewReader(psReturn))
	scannerScan(scanner)
	for scannerScan(scanner) {
		psArray = append(psArray, scannerText(scanner))
	}

	id, err := fuzzyfinderFind(
		psArray,
		func(i int) string {
			return psArray[i]
		})
	if err != nil {
		return err
	}

	containerID := stringsSplit(psArray[id], " ")
	if err := clipboardWriteAll(containerID[0]); err != nil {
		return err
	}

	return nil
}
