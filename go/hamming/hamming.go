package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	var errnum int = 0
	var err error = errors.New("Hamming Error")
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				errnum++
			}
		}

		if errnum == 0 {
			return errnum, nil
		} else {
			return errnum, nil
		}
	} else {
		return -1, err
	}
}
