// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/drone/go-scm/scm"
)

func encode(s string) string {
	result := strings.Replace(s, "/", "%2F", -1)
	result = strings.Replace(result, "--", "%2F", -1)
	fmt.Printf("%v => %v\n", s, result)
	return result
}

func encodeListOptions(opts scm.ListOptions) string {
	params := url.Values{}
	if opts.Page != 0 {
		params.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.Size != 0 {
		params.Set("per_page", strconv.Itoa(opts.Size))
	}
	return params.Encode()
}

func encodeMemberListOptions(opts scm.ListOptions) string {
	params := url.Values{}
	params.Set("membership", "true")
	if opts.Page != 0 {
		params.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.Size != 0 {
		params.Set("per_page", strconv.Itoa(opts.Size))
	}
	return params.Encode()
}

func encodeCommitListOptions(opts scm.CommitListOptions) string {
	params := url.Values{}
	if opts.Page != 0 {
		params.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.Size != 0 {
		params.Set("per_page", strconv.Itoa(opts.Size))
	}
	if opts.Ref != "" {
		params.Set("ref_name", opts.Ref)
	}
	return params.Encode()
}

func encodeIssueListOptions(opts scm.IssueListOptions) string {
	params := url.Values{}
	if opts.Page != 0 {
		params.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.Size != 0 {
		params.Set("per_page", strconv.Itoa(opts.Size))
	}
	if opts.Open && opts.Closed {
		params.Set("state", "all")
	} else if opts.Closed {
		params.Set("state", "closed")
	} else if opts.Open {
		params.Set("state", "opened")
	}
	return params.Encode()
}

func encodePullRequestListOptions(opts scm.PullRequestListOptions) string {
	params := url.Values{}
	if opts.Page != 0 {
		params.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.Size != 0 {
		params.Set("per_page", strconv.Itoa(opts.Size))
	}
	if opts.Open && opts.Closed {
		params.Set("state", "all")
	} else if opts.Closed {
		params.Set("state", "closed")
	} else if opts.Open {
		params.Set("state", "opened")
	}
	return params.Encode()
}
