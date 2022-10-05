package bitmap

import (
	"math/rand"
	"testing"
)

func TestSetBit(t *testing.T) {
	size := 1000
	offsets := make([]int64, size)
	for i := 0; i < size; i++ {
		offsets[i] = int64(i)
	}
	rand.Shuffle(size, func(i, j int) {
		offsets[i], offsets[j] = offsets[j], offsets[i]
	})
	offsets = offsets[0 : size/5]
	offsetMap := make(map[int64]struct{})

	bm := New()
	for _, offset := range offsets {
		offsetMap[offset] = struct{}{}
		bm.SetBit(offset, 1)
	}

	for i := 0; i < bm.BitSize(); i++ {
		offset := int64(i)
		_, expect := offsetMap[offset]
		actual := bm.GetBit(offset) > 0
		if expect != actual {
			t.Errorf("wrong value at %d", offset)
		}
	}

	bm2 := New()
	bm2.SetBit(15, 1)
	if bm2.GetBit(15) != 1 {
		t.Error("wrong value")
	}
	if bm2.GetBit(16) != 0 {
		t.Error("wrong value")
	}
}
