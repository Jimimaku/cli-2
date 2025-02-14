//go:build darwin
// +build darwin

package autostart

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ActiveState/cli/internal/assets"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/osutils/user"
	"github.com/ActiveState/cli/internal/strutils"
)

const (
	launchFileSource     = "com.activestate.platform.state.plist.tpl"
	launchFileFormatName = "com.activestate.platform.%s.plist"
)

func (a *app) enable() error {
	enabled, err := a.IsEnabled()
	if err != nil {
		return errs.Wrap(err, "Could not check if app autostart is enabled")
	}

	if enabled {
		return nil
	}

	path, err := a.InstallPath()
	if err != nil {
		return errs.Wrap(err, "Could not get launch file")
	}

	asset, err := assets.ReadFileBytes(launchFileSource)
	if err != nil {
		return errs.Wrap(err, "Could not read asset")
	}

	content, err := strutils.ParseTemplate(
		string(asset),
		map[string]interface{}{"Exec": a.Exec, "Args": strings.Join(a.Args, " ")})
	if err != nil {
		return errs.Wrap(err, "Could not parse %s", fmt.Sprintf(launchFileFormatName, filepath.Base(a.Exec)))
	}

	if err = fileutils.WriteFile(path, []byte(content)); err != nil {
		return errs.Wrap(err, "Could not write launch file")
	}
	return nil
}

func (a *app) disable() error {
	enabled, err := a.IsEnabled()
	if err != nil {
		return errs.Wrap(err, "Could not check if app autostart is enabled")
	}

	if !enabled {
		return nil
	}
	path, err := a.InstallPath()
	if err != nil {
		return errs.Wrap(err, "Could not get launch file")
	}
	return os.Remove(path)
}

func (a *app) IsEnabled() (bool, error) {
	path, err := a.InstallPath()
	if err != nil {
		return false, errs.Wrap(err, "Could not get launch file")
	}
	return fileutils.FileExists(path), nil
}

func (a *app) InstallPath() (string, error) {
	dir, err := user.HomeDir()
	if err != nil {
		return "", errs.Wrap(err, "Could not get home directory")
	}
	path := filepath.Join(dir, "Library/LaunchAgents", fmt.Sprintf(launchFileFormatName, filepath.Base(a.Exec)))
	return path, nil
}
