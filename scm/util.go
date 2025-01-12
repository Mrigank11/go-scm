// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scm

import "strings"

// Split splits the full repository name into segments.
func Split(s string) (owner, name string) {
	parts := strings.Split(s, "/")
	name = parts[len(parts)-1]
	if len(parts) > 1 {
		owner = strings.Join(parts[:len(parts)-1], "--")
	}
	return
}

// Join joins the repository owner and name segments to
// create a fully qualified repository name.
func Join(owner, name string) string {
	return owner + "/" + name
}

// TrimRef returns ref without the path prefix.
func TrimRef(ref string) string {
	ref = strings.TrimPrefix(ref, "refs/heads/")
	ref = strings.TrimPrefix(ref, "refs/tags/")
	return ref
}

// ExpandRef returns name expanded to the fully qualified
// reference path (e.g refs/heads/master).
func ExpandRef(name, prefix string) string {
	prefix = strings.TrimSuffix(prefix, "/")
	if strings.HasPrefix(name, "refs/") {
		return name
	}
	return prefix + "/" + name
}

// IsTag returns true if the reference path points to
// a tag object.
func IsTag(ref string) bool {
	return strings.HasPrefix(ref, "refs/tags/")
}
