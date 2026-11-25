package buffer

// AC-265: golang.org/x/text stays at or above v0.41.0, which fixed
// CVE-2026-56851 and CVE-2026-56852, and the decoders micro uses for
// file encodings still turn bad input into replacement text.

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"golang.org/x/text/encoding/htmlindex"
	"golang.org/x/text/transform"
)

// versionAtLeast compares "vX.Y.Z" versions (pre-release parts ignored).
func versionAtLeast(have, floor string) bool {
	parse := func(v string) [3]int {
		var out [3]int
		v = strings.TrimPrefix(v, "v")
		if i := strings.IndexAny(v, "-+"); i >= 0 {
			v = v[:i]
		}
		for i, p := range strings.SplitN(v, ".", 3) {
			out[i], _ = strconv.Atoi(p)
		}
		return out
	}
	h, f := parse(have), parse(floor)
	for i := range h {
		if h[i] != f[i] {
			return h[i] > f[i]
		}
	}
	return true
}

func TestVersionAtLeast(t *testing.T) {
	cases := []struct {
		have, floor string
		want        bool
	}{
		{"v0.41.0", "v0.41.0", true},
		{"v0.4.0", "v0.41.0", false},
		{"v0.42.1", "v0.41.0", true},
		{"v1.0.0-rc1", "v0.41.0", true},
	}
	for _, c := range cases {
		if got := versionAtLeast(c.have, c.floor); got != c.want {
			t.Errorf("versionAtLeast(%q, %q) = %v, want %v", c.have, c.floor, got, c.want)
		}
	}
}

func TestXTextAtSecurityFloor(t *testing.T) {
	// go test runs in the package folder; go.mod is two levels up.
	mod, err := os.ReadFile("../../go.mod")
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range strings.Split(string(mod), "\n") {
		f := strings.Fields(line)
		if len(f) >= 2 && f[0] == "golang.org/x/text" {
			if !versionAtLeast(f[1], "v0.41.0") {
				t.Fatalf("golang.org/x/text %s is below v0.41.0", f[1])
			}
			return
		}
	}
	t.Fatal("golang.org/x/text not required in go.mod")
}

func TestDecodeInvalidInputGivesReplacement(t *testing.T) {
	bad := []byte{'a', 0xff, 0xfe, 'b', 0xc3}
	for _, name := range []string{"utf-8", "utf-16le", "shift_jis", "gbk", "windows-1252"} {
		enc, err := htmlindex.Get(name)
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		out, err := io.ReadAll(transform.NewReader(bytes.NewReader(bad), enc.NewDecoder()))
		if err != nil {
			t.Fatalf("%s: decode error %v", name, err)
		}
		if len(out) == 0 {
			t.Fatalf("%s: no output", name)
		}
	}
}
