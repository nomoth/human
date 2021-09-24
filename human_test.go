package human_test

import (
	"testing"

	"github.com/nomoth/human"
)

type testCase struct {
	count  uint64
	result string
}

var testCases = []testCase{
	{0, "0 B"},
	{1, "1 B"},
	{1024, "1.0 KiB"},
	{1024 * 1024, "1.0 MiB"},
	{1024 * 1024 * 1024, "1.0 GiB"},
	{1024 * 1024 * 1024 * 1024, "1.0 TiB"},
	{1024 * 1024 * 1024 * 1024 * 1024, "1.0 PiB"},
	{1024 * 1024 * 1024 * 1024 * 1024 * 1024, "1.0 EiB"},
	{1.5 * 1024, "1.5 KiB"},
	{999 * 1024 * 1024 * 1024, "999.0 GiB"},
	{3333, "3.3 KiB"},
	{18446744073709551615, "16.0 EiB"}, //MAX VALUE
}

func TestBytesCount(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.result, func(t *testing.T) {
			r := human.BytesCount(tc.count)
			if r != tc.result {
				t.Errorf("expected '%s' got '%s'", tc.result, r)
			}
		})
	}
}
