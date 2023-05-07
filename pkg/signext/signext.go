package signext

// signBit(4) =  0b0000'1000
func signBit(bits int) int {
	return 1 << (bits - 1)
}

// signMask(4) = 0b1111'1000
func signMask(bits int) int {
	return ^(signBit(bits) - 1)
}

func SignExt(val, bits int) int {
	if val&signBit(bits) != 0 {
		return val | signMask(bits)
	}
	return val
}
