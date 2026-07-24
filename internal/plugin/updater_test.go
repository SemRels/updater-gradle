package plugin

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUpdaterUpdateBuildGradle(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "updater-gradle-")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.RemoveAll(dir) }()

	file := filepath.Join(dir, "build.gradle")
	if err := os.WriteFile(file, []byte("version = '1.2.3'\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := NewUpdater().Update(file, "1.3.0"); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	got, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), "version = '1.3.0'") {
		t.Fatalf("updated file = %s", got)
	}
}

func TestUpdaterUpdateGradleProperties(t *testing.T) {
	t.Parallel()

	updated, changed := replace("gradle.properties", "version=1.2.3\n", "1.3.0")
	if !changed || !strings.Contains(updated, "version=1.3.0") {
		t.Fatalf("replace() = %q changed=%v", updated, changed)
	}
}

func TestUpdaterMissingFile(t *testing.T) {
	t.Parallel()

	err := NewUpdater().Update(filepath.Join(t.TempDir(), "build.gradle"), "1.3.0")
	if err == nil || !strings.Contains(err.Error(), "read") {
		t.Fatalf("expected read error, got %v", err)
	}
}

func TestUpdaterMissingVersion(t *testing.T) {
	t.Parallel()

	_, changed := replace("build.gradle", "plugins {}\n", "1.3.0")
	if changed {
		t.Fatal("expected no replacement")
	}
}
