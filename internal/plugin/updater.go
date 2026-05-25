// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

// Package plugin updates Gradle version files in-place.
package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var buildPattern = regexp.MustCompile(`(?m)^(\s*version\s*=\s*['"])[^'"]*(['"].*)$`)
var propertiesPattern = regexp.MustCompile(`(?m)^(\s*version\s*=\s*).*$`)

// Updater updates Gradle build files.
type Updater struct{}

// NewUpdater creates an updater.
func NewUpdater() *Updater {
	return &Updater{}
}

// Update rewrites the version declaration in a Gradle file.
func (u *Updater) Update(path, version string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	updated, changed := replace(path, string(data), version)
	if !changed {
		return fmt.Errorf("version declaration not found in %s", path)
	}

	if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

func replace(path, content, version string) (string, bool) {
	if filepath.Base(path) == "gradle.properties" {
		if !propertiesPattern.MatchString(content) {
			return "", false
		}
		return propertiesPattern.ReplaceAllString(content, `${1}`+version), true
	}
	if !buildPattern.MatchString(content) {
		return "", false
	}
	return buildPattern.ReplaceAllString(content, `${1}`+version+`${2}`), true
}
