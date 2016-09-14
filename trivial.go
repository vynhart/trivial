// Package trivial provides trivial functionality. ^_^
package trivial

func StringsContain(s []string, e string) bool {
	for _, x := range s {
		if x == e {
			return true
		}
	}
	return false
}
