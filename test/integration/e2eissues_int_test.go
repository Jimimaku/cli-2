package integration

import (
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/testhelpers/e2e"
)

/*
Tests in this file are failing due to bugs in the test framework. They are skipped so as not to break our CI.
The intend is to collect e2e bugs here so that we can test that they are fixed once we refactor the framework.
*/

func TestMultipleSends(t *testing.T) {
	t.Skip("")

	if runtime.GOOS != "windows" {
		t.Skip("This test is only relevant on Windows")
	}

	ts := e2e.New(t, true)
	defer ts.Close()

	promptFile := filepath.Join(ts.Dirs.Work, "prompt.bat")
	err := fileutils.WriteFile(promptFile, []byte(`
@echo off
set /p "prompt1=Prompt 1: "
echo "Prompt 1: %prompt1%"
set /p "prompt2=Prompt 2: "
echo "Prompt 2: %prompt1%"
`))
	if err != nil {
		t.Fatal(err)
	}

	tp := ts.SpawnCmd(promptFile)
	tp.Expect("Prompt 1: ", 5*time.Second)
	tp.Send("Answer 1")
	tp.Expect("Prompt 1: Answer 1", 5*time.Second)
	tp.Expect("Prompt 2: ", 5*time.Second)
	tp.Send("Answer 2")
	tp.Expect("Prompt 2: Answer 2", 5*time.Second)
	tp.ExpectExitCode(0, 5*time.Second)
}
