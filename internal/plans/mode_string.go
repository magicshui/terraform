// Code generated by "stringer -type Mode"; DO NOT EDIT.

package plans

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NormalMode-0]
	_ = x[DestroyMode-68]
	_ = x[RefreshOnlyMode-82]
}

const (
	_Mode_name_0 = "NormalMode"
	_Mode_name_1 = "DestroyMode"
	_Mode_name_2 = "RefreshOnlyMode"
)

func (i Mode) String() string {
	switch {
	case i == 0:
		return _Mode_name_0
	case i == 68:
		return _Mode_name_1
	case i == 82:
		return _Mode_name_2
	default:
		return "Mode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
