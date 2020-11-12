package main

import (
	"bytes"
	"strings"
	"testing"
)

func testRun(t *testing.T) {
	testCases := []struct {
		name   string
		col    int
		op     string
		exp    string
		files  []string
		errMsg string
	}{
		{name: "RunAvg1File", col: 3, op: "avg", exp: "227.6\n",
			files:  []string{"./testdata/example.csv"},
			errMsg: ""},
		{name: "RunAvgMultiFiles", col: 3, op: "avg", exp: "233.84\n",
			files: []string{"./testdata/example.csv", "./testdata/example2.csv"}, errMsg: "",
		},
		{name: "RunFailRead", col: 2, op: "avg", exp: "",
			files:  []string{"./testdata/example.csv", "./testdata/fakefile.csv"},
			errMsg: "Cannot open file"},
		{name: "RunFailColumn", col: 0, op: "avg", exp: "", files: []string{"./testdata/example.csv"}, errMsg: "Invalid column"},
		{name: "RunFailNoFiles", col: 2, op: "avg", exp: "",
			files:  []string{},
			errMsg: "No input files"},
		{name: "RunFailOperation", col: 2, op: "invalid", exp: "", files: []string{"./testdata/example.csv"},
			errMsg: "Operation not supported: invalid",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var res bytes.Buffer
			err := run(tc.files, tc.op, tc.col, &res)

			if tc.errMsg != "" {
				if err == nil {
					t.Errorf("Expected error. Got nil instead")
				}
			}

			if !strings.Contains(err.Error(), tc.errMsg) {
				t.Errorf("Unexpected error: %q", err)
			}
			if res.String() != tc.exp {
				t.Errorf("Expected %q, got %q instead", tc.exp, &res)
			}
		})
	}
}
