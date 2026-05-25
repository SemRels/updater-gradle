// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

package main

import (
	"log"

	plugin "github.com/SemRels/updater-gradle/internal/plugin"
)

func main() {
	publisher := plugin.NewPublisher(plugin.Config{})
	log.Printf("updater-gradle plugin ready: updates Gradle project versions and publishes artifacts (%T)", publisher)
}
