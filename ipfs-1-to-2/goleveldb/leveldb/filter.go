// Copyright (c) 2012, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package leveldb

import (
	"github.com/ipfs/fs-repo-migrations/ipfs-1-to-2/goleveldb/leveldb/filter"
)

type iFilter struct {
	filter.Filter
}

func (f iFilter) Contains(filter, key []byte) bool {
	return f.Filter.Contains(filter, iKey(key).ukey())
}

func (f iFilter) NewGenerator() filter.FilterGenerator {
	return iFilterGenerator{f.Filter.NewGenerator()}
}

type iFilterGenerator struct {
	filter.FilterGenerator
}

func (g iFilterGenerator) Add(key []byte) {
	g.FilterGenerator.Add(iKey(key).ukey())
}
