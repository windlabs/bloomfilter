package bloomfilter

import (
	"testing"
)

func TestBloomFilter_Contains(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantRet bool
	}{
		{"case-1", "foo", false},
		{"case-2", "key", true},
	}
	bf := NewBloomFilter()
	bf.Add("key")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := bf.Contains(tt.value); gotRet != tt.wantRet {
				t.Errorf("Contains() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
