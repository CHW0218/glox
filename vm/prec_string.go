// Code generated by "stringer -type=Prec"; DO NOT EDIT.

package vm

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PrecNone-0]
	_ = x[PrecAssign-1]
	_ = x[PrecOr-2]
	_ = x[PrecAnd-3]
	_ = x[PrecEqual-4]
	_ = x[PrecComp-5]
	_ = x[PrecTerm-6]
	_ = x[PrecFactor-7]
	_ = x[PrecUnary-8]
	_ = x[PrecCall-9]
	_ = x[PrecPrimary-10]
}

const _Prec_name = "PrecNonePrecAssignPrecOrPrecAndPrecEqualPrecCompPrecTermPrecFactorPrecUnaryPrecCallPrecPrimary"

var _Prec_index = [...]uint8{0, 8, 18, 24, 31, 40, 48, 56, 66, 75, 83, 94}

func (i Prec) String() string {
	if i < 0 || i >= Prec(len(_Prec_index)-1) {
		return "Prec(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Prec_name[_Prec_index[i]:_Prec_index[i+1]]
}
