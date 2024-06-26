// Code generated by "stringer -type=Operation -linecomment"; DO NOT EDIT.

package iceberg

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OpTrue-0]
	_ = x[OpFalse-1]
	_ = x[OpIsNull-2]
	_ = x[OpNotNull-3]
	_ = x[OpIsNan-4]
	_ = x[OpNotNan-5]
	_ = x[OpLT-6]
	_ = x[OpLTEQ-7]
	_ = x[OpGT-8]
	_ = x[OpGTEQ-9]
	_ = x[OpEQ-10]
	_ = x[OpNEQ-11]
	_ = x[OpStartsWith-12]
	_ = x[OpNotStartsWith-13]
	_ = x[OpIn-14]
	_ = x[OpNotIn-15]
	_ = x[OpNot-16]
	_ = x[OpAnd-17]
	_ = x[OpOr-18]
}

const _Operation_name = "TrueFalseIsNullNotNullIsNaNNotNaNLessThanLessThanEqualGreaterThanGreaterThanEqualEqualNotEqualStartsWithNotStartsWithInNotInNotAndOr"

var _Operation_index = [...]uint8{0, 4, 9, 15, 22, 27, 33, 41, 54, 65, 81, 86, 94, 104, 117, 119, 124, 127, 130, 132}

func (i Operation) String() string {
	if i < 0 || i >= Operation(len(_Operation_index)-1) {
		return "Operation(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Operation_name[_Operation_index[i]:_Operation_index[i+1]]
}
