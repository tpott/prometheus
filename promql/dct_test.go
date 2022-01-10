// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package promql_test

import (
	"context"
    "testing"

    "github.com/stretchr/testify/require"

    "github.com/prometheus/prometheus/promql"
)

func TestRotate(t *testing.T) {
	vals := []float64{4, 3, 5, 10}
	i, rotated := promql.RotateToMax(vals)
	require.Equal(t, 3, i)
	require.Equal(t, []float64{10, 4, 3, 5}, rotated)
	recovered := promql.Unrotate(i, rotated)
	require.Equal(t, vals, recovered)
}

func TestDCT(t *testing.T) {
	ctx := context.Background()
	// Example values from https://docs.scipy.org/doc/scipy/reference/generated/scipy.fftpack.dct.html
	vals := []float64{4, 3, 5, 10}
	require.InDeltaSlice(t, []float64{30, -8, 6, -2}, promql.DCT1(ctx, vals), 1e-6)
}
