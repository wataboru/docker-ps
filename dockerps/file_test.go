package dockerps

import (
	"bufio"
	"fmt"
	"github.com/ktr0731/go-fuzzyfinder"
	"testing"
)

func TestCopy(t *testing.T) {
	count := 0
	tests := []struct {
		name              string
		fuzzyfinderFind   func(slice interface{}, itemFunc func(i int) string, opts ...fuzzyfinder.Option) (int, error)
		dockerPs          func() (string, error)
		stringsSplit      func(s, sep string) []string
		clipboardWriteAll func(text string) error
		scannerScan       func(scanner *bufio.Scanner) bool
		scannerText       func(scanner *bufio.Scanner) string
		wantErr           bool
	}{
		{
			name: "Normal",
			fuzzyfinderFind: func(slice interface{}, itemFunc func(i int) string, opts ...fuzzyfinder.Option) (int, error) {
				return 0, nil
			},
			dockerPs: func() (string, error) {
				return "hoge", nil
			},
			stringsSplit: func(s, sep string) []string {
				return []string{"ho ge"}
			},
			clipboardWriteAll: func(text string) error {
				return nil
			},
			scannerScan: func(scanner *bufio.Scanner) bool {
				count++
				return count <= 2
			},
			scannerText: func(scanner *bufio.Scanner) string {
				return "ho ge"
			},
			wantErr: false,
		},
		{
			name: "ErrorBecauseFuzzyFinderFindReturnError",
			dockerPs: func() (string, error) {
				return "", fmt.Errorf("error")
			},
			fuzzyfinderFind: func(slice interface{}, itemFunc func(i int) string, opts ...fuzzyfinder.Option) (int, error) {
				return 0, fmt.Errorf("error")
			},
			wantErr: true,
		},
		{
			name: "ErrorBecauseDockerPsReturnError",
			fuzzyfinderFind: func(slice interface{}, itemFunc func(i int) string, opts ...fuzzyfinder.Option) (int, error) {
				return 0, nil
			},
			dockerPs: func() (string, error) {
				return "", fmt.Errorf("error")
			},
			wantErr: true,
		},
		{
			name: "ErrorBecauseClipboardWriteAllReturnError",
			fuzzyfinderFind: func(slice interface{}, itemFunc func(i int) string, opts ...fuzzyfinder.Option) (int, error) {
				return 0, nil
			},
			dockerPs: func() (string, error) {
				return "hoge", nil
			},
			stringsSplit: func(s, sep string) []string {
				return []string{"ho ge"}
			},
			clipboardWriteAll: func(text string) error {
				return fmt.Errorf("error")
			},
			scannerScan: func(scanner *bufio.Scanner) bool {
				count++
				return count <= 5
			},
			scannerText: func(scanner *bufio.Scanner) string {
				return "ho ge"
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dockerPs = tt.dockerPs
			fuzzyfinderFind = tt.fuzzyfinderFind
			stringsSplit = tt.stringsSplit
			scannerText = tt.scannerText
			scannerScan = tt.scannerScan
			clipboardWriteAll = tt.clipboardWriteAll
			if err := Copy(); (err != nil) != tt.wantErr {
				t.Errorf("copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
