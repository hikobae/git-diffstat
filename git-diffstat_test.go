package main

import "testing"

var parseLineTest = []struct {
	line   string
	result result
}{
	{"25      17      git-diffstat.go", result{path: "git-diffstat.go", add: "25", delete: "17"}},
	{"-       -       a.png", result{path: "a.png", add: "-", delete: "-"}},
}

func TestParseLine(t *testing.T) {
	var r result
	for _, tc := range parseLineTest {
		if err := parseLine(tc.line, &r); err != nil {
			t.Errorf("parseLine() error, err=%v", err)
			continue
		}
		if r.path != tc.result.path || r.add != tc.result.add || r.delete != tc.result.delete {
			t.Errorf("expected %s %s %s, actual %s %s %s", tc.result.path, tc.result.add, tc.result.delete, r.path, r.add, r.delete)
			continue
		}
	}
}
