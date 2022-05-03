package warmup

import "testing"

func TestSockMerchant(t *testing.T) {
	var (
		n  int32 = 7
		ar       = []int32{1, 2, 1, 2, 1, 3, 2}
	)
	SockMerchant(n, ar)
}
