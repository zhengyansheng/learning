package unit

import (
	"testing"
	"unicode/utf8"
	"unit_test/funcs"
)

func FuzzReverse(f *testing.F) {

	testcases := []string{"hello world", " ", "!12345"}

	for _, tc := range testcases {
		f.Add(tc) // 输入测试种子
	}

	f.Fuzz(func(t *testing.T, orig string) {

		if !utf8.ValidString(orig) { //忽略构造的无效字符（非UTF-8编码字符串)
			return
		}

		rev, err := funcs.Reverse(orig)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !utf8.ValidString(rev) {
			t.Fatalf("Reverse produced invalid UTF-8 string %q", rev)
		}

		doubleRev, err := funcs.Reverse(rev)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if orig != doubleRev {
			t.Fatalf("Before: %q, after: %q", orig, doubleRev)
		}
	})
}

// TestFuzz 单元测试
func TestReverse(t *testing.T) {
	testcases := []struct {
		in   string
		want string
	}{
		{
			in:   "hello world",
			want: "dlrow olleh",
		},
		{
			in:   " ",
			want: " ",
		},
	}

	for _, tc := range testcases {
		rev, err := funcs.Reverse(tc.in)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if rev != tc.want {
			t.Errorf("Reverse: %q, want: %q", rev, tc.want)
		}
	}
}
