package semver_test

import (
	"testing"

	"github.com/anttikivi/agricola/internal/semver"
)

var tests = []struct { //nolint:gochecknoglobals
	in  string
	out string
}{
	{"", ""},

	{"0.1.0-alpha.24+sha.19031c2.darwin.amd64", "0.1.0-alpha.24"},
	{"0.1.0-alpha.24+sha.19031c2-darwin-amd64", "0.1.0-alpha.24"},

	{"bad", ""},
	{"1-alpha.beta.gamma", ""},
	{"1-pre", ""},
	{"1+meta", ""},
	{"1-pre+meta", ""},
	{"1.2-pre", ""},
	{"1.2+meta", ""},
	{"1.2-pre+meta", ""},
	{"1.0.0-alpha", "1.0.0-alpha"},
	{"1.0.0-alpha.1", "1.0.0-alpha.1"},
	{"1.0.0-alpha.beta", "1.0.0-alpha.beta"},
	{"1.0.0-beta", "1.0.0-beta"},
	{"1.0.0-beta.2", "1.0.0-beta.2"},
	{"1.0.0-beta.11", "1.0.0-beta.11"},
	{"1.0.0-rc.1", "1.0.0-rc.1"},
	{"1", ""},
	{"1.0", ""},
	{"1.0.0", "1.0.0"},
	{"1.2", ""},
	{"1.2.0", "1.2.0"},
	{"1.2.3-456", "1.2.3-456"},
	{"1.2.3-456.789", "1.2.3-456.789"},
	{"1.2.3-456-789", "1.2.3-456-789"},
	{"1.2.3-456a", "1.2.3-456a"},
	{"1.2.3-pre", "1.2.3-pre"},
	{"1.2.3-pre+meta", "1.2.3-pre"},
	{"1.2.3-pre.1", "1.2.3-pre.1"},
	{"1.2.3-zzz", "1.2.3-zzz"},
	{"1.2.3", "1.2.3"},
	{"1.2.3+meta", "1.2.3"},
	{"1.2.3+meta-pre", "1.2.3"},
	{"1.2.3+meta-pre.sha.256a", "1.2.3"},

	{"vbad", ""},
	{"v1-alpha.beta.gamma", ""},
	{"v1-pre", ""},
	{"v1+meta", ""},
	{"v1-pre+meta", ""},
	{"v1.2-pre", ""},
	{"v1.2+meta", ""},
	{"v1.2-pre+meta", ""},
	{"v1.0.0-alpha", "1.0.0-alpha"},
	{"v1.0.0-alpha.1", "1.0.0-alpha.1"},
	{"v1.0.0-alpha.beta", "1.0.0-alpha.beta"},
	{"v1.0.0-beta", "1.0.0-beta"},
	{"v1.0.0-beta.2", "1.0.0-beta.2"},
	{"v1.0.0-beta.11", "1.0.0-beta.11"},
	{"v1.0.0-rc.1", "1.0.0-rc.1"},
	{"v1", ""},
	{"v1.0", ""},
	{"v1.0.0", "1.0.0"},
	{"v1.2", ""},
	{"v1.2.0", "1.2.0"},
	{"v1.2.3-456", "1.2.3-456"},
	{"v1.2.3-456.789", "1.2.3-456.789"},
	{"v1.2.3-456-789", "1.2.3-456-789"},
	{"v1.2.3-456a", "1.2.3-456a"},
	{"v1.2.3-pre", "1.2.3-pre"},
	{"v1.2.3-pre+meta", "1.2.3-pre"},
	{"v1.2.3-pre.1", "1.2.3-pre.1"},
	{"v1.2.3-zzz", "1.2.3-zzz"},
	{"v1.2.3", "1.2.3"},
	{"v1.2.3+meta", "1.2.3"},
	{"v1.2.3+meta-pre", "1.2.3"},
	{"v1.2.3+meta-pre.sha.256a", "1.2.3"},

	{"agerbad", ""},
	{"ager1-alpha.beta.gamma", ""},
	{"ager1-pre", ""},
	{"ager1+meta", ""},
	{"ager1-pre+meta", ""},
	{"ager1.2-pre", ""},
	{"ager1.2+meta", ""},
	{"ager1.2-pre+meta", ""},
	{"ager1.0.0-alpha", "1.0.0-alpha"},
	{"ager1.0.0-alpha.1", "1.0.0-alpha.1"},
	{"ager1.0.0-alpha.beta", "1.0.0-alpha.beta"},
	{"ager1.0.0-beta", "1.0.0-beta"},
	{"ager1.0.0-beta.2", "1.0.0-beta.2"},
	{"ager1.0.0-beta.11", "1.0.0-beta.11"},
	{"ager1.0.0-rc.1", "1.0.0-rc.1"},
	{"ager1", ""},
	{"ager1.0", ""},
	{"ager1.0.0", "1.0.0"},
	{"ager1.2", ""},
	{"ager1.2.0", "1.2.0"},
	{"ager1.2.3-456", "1.2.3-456"},
	{"ager1.2.3-456.789", "1.2.3-456.789"},
	{"ager1.2.3-456-789", "1.2.3-456-789"},
	{"ager1.2.3-456a", "1.2.3-456a"},
	{"ager1.2.3-pre", "1.2.3-pre"},
	{"ager1.2.3-pre+meta", "1.2.3-pre"},
	{"ager1.2.3-pre.1", "1.2.3-pre.1"},
	{"ager1.2.3-zzz", "1.2.3-zzz"},
	{"ager1.2.3", "1.2.3"},
	{"ager1.2.3+meta", "1.2.3"},
	{"ager1.2.3+meta-pre", "1.2.3"},
	{"ager1.2.3+meta-pre.sha.256a", "1.2.3"},

	{"agricolabad", ""},
	{"agricola1-alpha.beta.gamma", ""},
	{"agricola1-pre", ""},
	{"agricola1+meta", ""},
	{"agricola1-pre+meta", ""},
	{"agricola1.2-pre", ""},
	{"agricola1.2+meta", ""},
	{"agricola1.2-pre+meta", ""},
	{"agricola1.0.0-alpha", "1.0.0-alpha"},
	{"agricola1.0.0-alpha.1", "1.0.0-alpha.1"},
	{"agricola1.0.0-alpha.beta", "1.0.0-alpha.beta"},
	{"agricola1.0.0-beta", "1.0.0-beta"},
	{"agricola1.0.0-beta.2", "1.0.0-beta.2"},
	{"agricola1.0.0-beta.11", "1.0.0-beta.11"},
	{"agricola1.0.0-rc.1", "1.0.0-rc.1"},
	{"agricola1", ""},
	{"agricola1.0", ""},
	{"agricola1.0.0", "1.0.0"},
	{"agricola1.2", ""},
	{"agricola1.2.0", "1.2.0"},
	{"agricola1.2.3-456", "1.2.3-456"},
	{"agricola1.2.3-456.789", "1.2.3-456.789"},
	{"agricola1.2.3-456-789", "1.2.3-456-789"},
	{"agricola1.2.3-456a", "1.2.3-456a"},
	{"agricola1.2.3-pre", "1.2.3-pre"},
	{"agricola1.2.3-pre+meta", "1.2.3-pre"},
	{"agricola1.2.3-pre.1", "1.2.3-pre.1"},
	{"agricola1.2.3-zzz", "1.2.3-zzz"},
	{"agricola1.2.3", "1.2.3"},
	{"agricola1.2.3+meta", "1.2.3"},
	{"agricola1.2.3+meta-pre", "1.2.3"},
	{"agricola1.2.3+meta-pre.sha.256a", "1.2.3"},
}

func TestIsValid(t *testing.T) {
	t.Parallel()

	for _, tt := range tests {
		ok := semver.IsValid(tt.in)
		if ok != (tt.out != "") {
			t.Errorf("IsValid(%q) = %v, want %v", tt.in, ok, !ok)
		}
	}
}

func TestVersionString(t *testing.T) {
	t.Parallel()

	for _, tt := range tests {
		// Don't test the cases where the versions don't parse.
		if tt.out != "" {
			v, _ := semver.Parse(tt.in)

			ok := v.String() == tt.out
			if !ok {
				t.Errorf("Version{%q}.String() = %v, want %v", tt.in, v, tt.out)
			}
		}
	}
}
