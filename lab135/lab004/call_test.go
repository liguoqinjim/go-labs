package main

import (
	"fmt"
	"github.com/bouk/monkey"
	"os/exec"
	"reflect"
	"testing"
)

func TestCall(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer monkey.UnpatchAll()

		//patch cmd的CombinedOutput方法
		monkey.PatchInstanceMethod(reflect.TypeOf((*exec.Cmd)(nil)), "CombinedOutput", func(_ *exec.Cmd) ([]byte, error) {
			return []byte("results"), nil
		})

		//patch reportExecFailed方法
		monkey.Patch(reportExecFailed, func(msg string) string {
			return msg
		})

		rc, output := Call("anycmd", "anyvalue")
		if rc != 0 {
			t.Errorf("rc got '%d' want 0", rc)
		}
		if output != "results" {
			t.Errorf("output got '%s' want 'results'", output)
		}
	})

	t.Run("failure", func(t *testing.T) {
		defer monkey.UnpatchAll()

		monkey.PatchInstanceMethod(reflect.TypeOf((*exec.Cmd)(nil)), "CombinedOutput", func(_ *exec.Cmd) ([]byte, error) {
			return []byte(""), fmt.Errorf("cmd run error")
		})

		rc, output := Call("anycmd", "anyvalue")
		if rc != 1 {
			t.Errorf("rc got '%d' want 1", rc)
		}
		if output != "cmd run error" {
			t.Errorf("output got '%s' want 'cmd run error'", output)
		}
	})
}
