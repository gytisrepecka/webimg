/*
 * Copyright © 2020 Gytis Repečka (gytis@repecka.com)
 *
 * This file is part of webimg.
 *
 * webimg is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package webimg

import (
	"fmt"
)

// FormatVersion constructs the version string for the application
func FormatVersion() string {
	// return serverSoftware + " " + softwareVer
	return "gowebimg 0.0.1"
}

// OutputVersion prints out the version of the application.
func OutputVersion() {
	fmt.Println(FormatVersion())
}
