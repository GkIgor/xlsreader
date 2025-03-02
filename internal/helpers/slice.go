package helpers

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func MakeSliceSet(str []string) map[string]struct{} {
	set := make(map[string]struct{}, len(str))

	for _, item := range str {
		set[item] = struct{}{}
	}
	return set
}
