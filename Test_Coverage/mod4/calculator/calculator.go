package calculator

type Calculator struct {
	name string
}

func Sum(a, b uint8) uint16 {                    // 1st Code branch
	return uint16(a) + uint16(b) // 1st Statement
}

func Multiply(a, b int8) int16 {				// 2nd Code branch
	res := int16(a) * int16(b) // 2nd Statement
	return res //3rd Statement
}

func Abs(a int8) int8 {							// 3rd Code branch and 4th Statement
	if a < 0 {    // 4th Code branch (inner)
		return -a //5th Statement (inner)
	}
    // 6th Statement (inner)
	return a // 5th Code branch (inner)
}