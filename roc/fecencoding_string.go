// Code generated by "stringer -type=FecEncoding -trimprefix=FecEncoding"; DO NOT EDIT.

package roc

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FecEncodingDisable - -1]
	_ = x[FecEncodingDefault-0]
	_ = x[FecEncodingRs8m-1]
	_ = x[FecEncodingLdpcStaircase-2]
}

const _FecEncoding_name = "DisableDefaultRs8mLdpcStaircase"

var _FecEncoding_index = [...]uint8{0, 7, 14, 18, 31}

func (i FecEncoding) String() string {
	i -= -1
	if i < 0 || i >= FecEncoding(len(_FecEncoding_index)-1) {
		return "FecEncoding(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _FecEncoding_name[_FecEncoding_index[i]:_FecEncoding_index[i+1]]
}
