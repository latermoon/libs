package bitmaps

const (
	mask = 8
)

type Bitmaps []uint8

// make a new bitmaps
func New(size int) Bitmaps {
	count := int(size / mask)
	if size%mask > 0 {
		count += 1
	}
	return make(Bitmaps, count)
}

// Size return the size of bitmaps, and the size may larger than initialize size
func (b Bitmaps) Size() int {
	return len(b) * mask
}

// SetBit set the bit at offset, true means 1, false means 0
func (b Bitmaps) SetBit(offset int, value bool) {
	if value {
		b[(offset / mask)] |= 1 << uint8(offset%mask)
	} else {
		b[(offset / mask)] ^= 1 << uint8(offset%mask)
	}
}

// GetBit return the bit value at offset, true means 1, false means 0
func (b Bitmaps) GetBit(offset int) bool {
	return b[(offset/mask)]&(1<<uint8(offset%mask)) > 0
}
