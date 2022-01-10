package promql

import (
	"context"
	// "errors"
	"math"
)

// type Name Actual

// Index is meant to wrap int, but be 0 indexed, i.e. non-negative, for value
// slices.
type Index = int

// Value represents the input space and is meant to make it easier to codemod
type Value = float64

// FValue represents the frequency space
type FValue = float64

// matrix is a slice of column values. Example accessor: `mat[rowI][colJ]`
type matrix = [][]Value

// column is a slice of values
type column = []Value

// IndexInvalid is the only negative value of type Index allowed
const IndexInvalid = Index(-1)

// Normalization?
// scipy's default (None), i.e. ??
// matlab's default, scipy's "ortho", i.e. ??
// TODO scipy.fftpack.dct(np.array([4, 3, 5, 10]), type=1, norm=None)

// Notes from https://en.wikipedia.org/wiki/Discrete_cosine_transform#DCT-I
// X_k is the output value in the k-th column.
// x_k is the input value in the k-th column.

// https://docs.scipy.org/doc/scipy/reference/generated/scipy.fftpack.dct.html

// DCT type 1
func DCT1(ctx context.Context, vals []Value) []FValue {
	n := len(vals)
	if n <= 1 {
		return nil
	}
	// TODO check "identity" cache
	dctMat := make(matrix, n)
	for i := 0; i < n; i++ {
		dctMat[i] = make([]Value, n)
		for j := 0; j < n; j++ {
			// TODO why did we normalize everything by 2?
			if i == 0 {
				dctMat[i][j] = 1 // 0.5
				continue
			} else if i == n - 1 && j % 2 == 0 {
				dctMat[i][j] = 1 // 0.5
				continue
			} else if i == n - 1 {
				dctMat[i][j] = -1 // -0.5
				continue
			}
			dctMat[i][j] = math.Cos((math.Pi / float64(n - 1)) * float64(i) * float64(j)) * 2
		}
	}
	return matMult(matrix{vals}, dctMat)[0]
}

// DCT type 2, the default DCT
func DCT2(ctx context.Context, vals []Value) []FValue {
	n := len(vals)
	if n == 0 {
		return nil
	}
	ret := make([]FValue, n)
	return ret
}

// DCT type 3, the default IDCT
func DCT3(ctx context.Context, vals []Value) []FValue {
	n := len(vals)
	if n == 0 {
		return nil
	}
	ret := make([]FValue, n)
	return ret
}

// DCT type 4
func DCT4(ctx context.Context, vals []Value) []FValue {
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

// Unrotate takes the original index that the max was at, and reverses the
// transformation that RotateToMax applies. It will return an empty slice if
// it's passed an empty slice, or an invalid index.
func Unrotate(i Index, vals []Value) []Value {
	n := len(vals)
	if n == 0 || i == IndexInvalid || i >= n {
		return nil
	}
	return append(vals[n-i:], vals[0:n-i]...)
}

var identityCache map[Index]matrix

// identity returns an identity matrix of size n x n
func identity(ctx context.Context, n Index) matrix {
	if identityCache == nil {
		identityCache = make(map[Index]matrix)
	}
	v, ok := identityCache[n]
	if ok {
		return v
	}
	mat := make(matrix, n)
	for i := 0; i <= n; i++ {
		// TODO add `1` at index `i`
		mat = append(mat, make([]Value, n))
	}
	return mat
}

// transpose returns the transposed matrix
func transpose(m matrix) matrix {
	nRows := len(m)
	if nRows == 0 {
		return nil
	}
	nCols := len(m[0])
	if nCols == 0 {
		return nil
	}
	ret := make(matrix, nCols)
	for i := 0; i < nCols; i++ {
		// ret = append(ret, make([]Value, nRows))
		ret[i] = make([]Value, nRows)
	}
	for i := 0; i < nCols; i++ {
		for j := 0; j < nRows; j++ {
			ret[i][j] = m[j][i]
		}
	}
	return ret
}

// matMult returns the matrix multiplied result of m * n
func matMult(m, n matrix) matrix {
	mRows := len(m)
	if mRows == 0 {
		return nil
	}
	mCols := len(m[0])
	if mCols == 0 {
		return nil
	}
	nRows := len(n)
	if nRows == 0 {
		return nil
	}
	nCols := len(n[0])
	if nCols == 0 {
		return nil
	}
	if mCols != nCols {
		// Maybe return an error instead?
		return nil
	}
	ret := make(matrix, mRows)
	for i := 0; i < mRows; i++ {
		ret[i] = make([]Value, nCols)
		for j := 0; j < nCols; j++ {
			for k := 0; k < mCols; k++ {
				ret[i][j] += m[i][k] * n[k][j]
			}
		}
	}
	return ret
}
