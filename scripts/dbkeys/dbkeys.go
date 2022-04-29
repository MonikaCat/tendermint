package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/creachadair/ctrl"
	tmdb "github.com/tendermint/tm-db"
)

var (
	dbSpec   = flag.String("db", "", `Database spec ("backend:path", required)`)
	statBase = flag.Int("base", 4, "Histogram sample base")
	statBars = flag.Int("bar", 20, "Histogram bar length")
)

func main() {
	flag.Parse()

	ctrl.Run(func() error {
		backend, path, err := parseDatabaseSpec(*dbSpec)
		if err != nil {
			ctrl.Fatalf("Parsing -db: %v", err)
		}

		// tm-db expects the target directory and the path to be separate.
		// Default to the current working directory if one is not specified.
		dir, base := filepath.Split(path)
		if dir == "" {
			dir = "."
		}
		base = strings.TrimSuffix(base, ".db")

		db, err := tmdb.NewDB(base, tmdb.BackendType(backend), dir)
		if err != nil {
			ctrl.Fatalf("NewDB(%q, %q, %q): %v", base, backend, dir, err)
		}
		defer db.Close()

		// Iterate through all the keys and aggregate statistics.
		// The DB interface has a Stats method, but it isn't always implemented.
		it, err := db.Iterator(nil, nil)
		if err != nil {
			ctrl.Fatalf("Iterator: %v", err)
		}

		var h = histogram{base: *statBase}
		for it.Valid() {
			h.addSample(len(it.Value()))
			fmt.Printf("%#q\n", string(it.Key()))
			it.Next()
		}
		if err := it.Close(); err != nil {
			ctrl.Fatalf("Closing iterator: %v", err)
		}

		h.WriteTo(os.Stdout)
		return nil
	})
}

func parseDatabaseSpec(s string) (backend, path string, _ error) {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) == 1 {
		return "", "", errors.New("invalid backend:path spec")
	} else if parts[0] == "" || parts[1] == "" {
		return "", "", errors.New("backend and path must be non-empty")
	}
	return parts[0], parts[1], nil
}

type histogram struct {
	samples int
	sum     int
	max     int
	base    int
	buckets []int // b[i] is count of sizes ≤ base^i
}

func (h *histogram) addSample(size int) {
	h.samples++
	h.sum += size
	if size > h.max {
		h.max = size
	}
	if h.base <= 0 {
		h.base = 4
	}

	v := size
	for i := 0; ; i++ {
		if len(h.buckets) <= i {
			h.buckets = append(h.buckets, 0)
		}

		w := v / h.base
		if w == 0 {
			h.buckets[i]++
			return
		}
		v = w
	}
}

func (h *histogram) WriteTo(w io.Writer) (int64, error) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "n = %d\n", h.samples)
	fmt.Fprintf(&buf, "max = %d\n", h.max)
	fmt.Fprintf(&buf, "avg = %.1f\n", float64(h.sum)/float64(h.samples))
	fmt.Fprintf(&buf, "[%d^i]\n", h.base)
	for i, v := range h.buckets {
		incr := int(20 * float64(v) / float64(h.samples))
		fmt.Fprintf(&buf, "%-2d %-6d %s\n", i, v, strings.Repeat("=", incr))
	}
	return io.Copy(w, &buf)
}
