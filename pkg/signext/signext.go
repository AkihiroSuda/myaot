package signext

// signBit(4) =  0b0000'1000
func signBit(pos int) int {
	return 1 << (pos - 1)
}

// signMask(4) = 0b1111'1000
func signMask(pos int) int {
	return ^(signBit(pos) - 1)
}

func SignExt(val, pos int) int {
	if val&signBit(pos) != 0 {
		return val | signMask(pos)
	}
	return val
}
