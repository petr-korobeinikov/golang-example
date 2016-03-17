package util

// MakeInterfaceSliceFromIntSlice makes []interface from []int.
func MakeInterfaceSliceFromIntSlice(in []int) (out []interface{}) {
	out = make([]interface{}, len(in))
	for i, d := range in {
		out[i] = d
	}
	return
}
