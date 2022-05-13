package dockerps

import (
	"fmt"
	"os/exec"
	"testing"
)

func Test__dockerPs(t *testing.T) {
	tests := []struct {
		name          string
		execCommand   func(name string, arg ...string) *exec.Cmd
		commandOutput func(c *exec.Cmd) ([]byte, error)
		want          string
		wantErr       bool
	}{
		{
			name: "Normal",
			execCommand: func(name string, arg ...string) *exec.Cmd {
				return &exec.Cmd{}
			},
			commandOutput: func(c *exec.Cmd) ([]byte, error) {
				return []byte("hoge"), nil
			},
			want:    "hoge",
			wantErr: false,
		},
		{
			name: "ErrorBecauseCommandReturnError",
			execCommand: func(name string, arg ...string) *exec.Cmd {
				return &exec.Cmd{}
			},
			commandOutput: func(c *exec.Cmd) ([]byte, error) {
				return []byte(""), fmt.Errorf("error")
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			execCommand = tt.execCommand
			commandOutput = tt.commandOutput
			got, err := _dockerPs()
			if (err != nil) != tt.wantErr {
				t.Errorf("_dockerPs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("_dockerPs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
