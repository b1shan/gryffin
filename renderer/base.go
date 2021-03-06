// Copyright 2015, Yahoo Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package renderer

import (
	"github.com/yahoo/gryffin"
)

type BaseRenderer struct {
	chanResponse chan *gryffin.Scan
	chanLinks    chan *gryffin.Scan
}

func (r *BaseRenderer) Do(s *gryffin.Scan) {
	// Dummy operation, just close the channels.
	defer close(r.chanResponse)
	defer close(r.chanLinks)
}

func (r *BaseRenderer) GetRequestBody() <-chan *gryffin.Scan {
	return r.chanResponse
}

func (r *BaseRenderer) GetLinks() <-chan *gryffin.Scan {
	return r.chanLinks
}
