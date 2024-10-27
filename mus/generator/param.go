package generator

func param(p string) string {
	// TODO Remove this check.
	// if p == "" {
	// 	return ""
	// }
	return p + ", "
}
