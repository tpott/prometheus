package promql

// import (
	// "errors"
	// "math"
// )

// type Name Actual

// Index is meant to wrap int, but be 0 indexed, i.e. non-negative, for value
// slices.
type Index = int

// Value represents the input space and is meant to make it easier to codemod
type Value = float64

// FValue represents the frequency space
type FValue = float64

type matrix = [][]Value

// IndexInvalid is the only negative value of type Index allowed
const IndexInvalid = Index(-1)

// Normalization?
// scipy's default (None), i.e. ??
// matlab's default, scipy's "ortho", i.e. ??
// scipy.fftpack.dct(np.array([4, 3, 5, 10]), type=1, norm=None)

// DCT type 1
func DCT1(vals []Value) []FValue {
	n := len(vals)
	if n == 0 {
		return nil
	}
	ret := make([]FValue, n)
	return ret
}

// DCT type 2, the default DCT
func DCT2(vals []Value) []FValue {
	n := len(vals)
	if n == 0 {
		return nil
	}
	ret := make([]FValue, n)
	return ret
}

// DCT type 3, the default IDCT
func DCT3(vals []Value) []FValue {
	n := len(vals)
	if n == 0 {
		return nil
	}
	ret := make([]FValue, n)
	return ret
}

// DCT type 4
func DCT4(vals []Value) []FValue {
	n := len(vals)
	if n == 0 {
		return nil
	}
	ret := make([]FValue, n)
	return ret
}

// DCT default function is DCT type 2
var DCT = DCT2

// IDCT default function is DCT type 3
var IDCT = DCT3

// RotateToMax returns the index of the max value in the input slice and a slice
// rotated to that point. If an empty slice is used, it returns -1, nil
// TODO testify https://go.dev/play/p/pBXY74jDMW1
func RotateToMax(vals []Value) (Index, []Value) {
	n := len(vals)
	if n == 0 {
		return IndexInvalid, nil
	}
	maxI := Index(0)
	var maxV *Value
	for i, v := range vals {
		if maxV == nil || v > *maxV {
			maxI = i
			maxV = &vals[i]
		}
	}
	return maxI, append(vals[maxI:n], vals[0:maxI]...)
}

// identity returns an identity matrix of size n x n
func identity(n Index) matrix {
	mat := make([][]Value, n)
	for i := 0; i <= n; i++ {
		// TODO add `1` at index `i`
		mat = append(mat, make([]Value, n))
	}
	return mat
}

// matMult returns the matrix multiplied result of m * n
func matMult(m, n matrix) (matrix) {
	return nil
}
