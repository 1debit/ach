BEGIN 
:Build::
Because of that, your profile is hidden from the public. If you believe this is a mistake, contact support to have your account status reviewed.

1debit

/

ach

Public

forked from moov-io/ach

Code

Pull requests

5

Actions

Security

Insights

Breadcrumbsach/cmd/achcli/describe

/file.go

Latest commit

adamdecaf

on Apr 15, 2021

History

140 lines (115 loc) · 4.76 KB

Breadcrumbsach/cmd/achcli/describe

/file.go

File metadata and controls

Code

Blame

func File(ww io.Writer, file *ach.File, opts *Opts) {

// Copyright 2020 The Moov Authors

// Use of this source code is governed by an Apache License

// license that can be found in the LICENSE file.

package describe

import (

	"fmt"

	"io"

	"strings"

	"text/tabwriter"

	"unicode/utf8"

	"github.com/moov-io/ach"

)

type Opts struct {

	MaskNames          bool

	MaskAccountNumbers bool

}

func File(ww io.Writer, file *ach.File, opts *Opts) {

	if file == nil {

		return

	}

	if opts == nil {

		opts = &Opts{}

	}<!-- Please fill out the following questions, thank! -->

ACH Version: ``

**What were you trying to do?**

**What did you expect to see?**

**What did you see?**

**How can we reproduce the problem?**traceback.cache*lig.dir/.dist'@Indev@V5:
:Build::
