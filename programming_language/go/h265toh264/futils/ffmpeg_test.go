package futils

import (
	"testing"
)

func TestCheckH265Format(t *testing.T) {
	filePath1 := "/Users/yangchao/Downloads/testfile.mp4"
	filePath2 := "/Users/yangchao/Downloads/output.mp4"
	result1, err := CheckH265Format(&filePath1)
	if err != nil {
		t.Error("")
	}
	if !result1 {
		t.Error("")
	}
	result2, err := CheckH265Format(&filePath2)
	if err != nil {
		t.Error("")
	}
	if result2 {
		t.Error("")
	}

}
