package parser

func RequestMethodValid(method string) bool {
	methods := []string{"GET", "POST", "DELETE", "PUT", "PATCH"}
	for _, m := range methods {
		if method == m {
			return true
		}
	}
	return false
}
