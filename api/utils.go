package api

/*
 * Takes a string representing a field we're selecting on
 * and returns ".*" if it's empty, allowing us to match
 * on ANY rather than NONE as a default.
 */
func DefMatch(f string) string {
	if f == "" {
		return ".*"
	} else {
		return f
	}
}
