// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package leveldb

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/MetalBlockchain/metalgo/database"
	"github.com/MetalBlockchain/metalgo/utils/logging"
)

func TestInterface(t *testing.T) {
	for _, test := range database.Tests {
		folder := t.TempDir()
		db, err := New(folder, nil, logging.NoLog{}, "", prometheus.NewRegistry())
		if err != nil {
			t.Fatalf("leveldb.New(%q, logging.NoLog{}) errored with %s", folder, err)
		}

		defer db.Close()

		test(t, db)

		// The database may have been closed by the test, so we don't care if it
		// errors here.
		_ = db.Close()
	}
}

func BenchmarkInterface(b *testing.B) {
	for _, size := range database.BenchmarkSizes {
		keys, values := database.SetupBenchmark(b, size[0], size[1], size[2])
		for _, bench := range database.Benchmarks {
			folder := b.TempDir()

			db, err := New(folder, nil, logging.NoLog{}, "", prometheus.NewRegistry())
			if err != nil {
				b.Fatal(err)
			}

			bench(b, db, "leveldb", keys, values)

			// The database may have been closed by the test, so we don't care if it
			// errors here.
			_ = db.Close()
		}
	}
}
