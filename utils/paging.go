package utils

// GetFromAndLimit Formats per page and page to from/limit with max per page check
func GetFromAndLimit(perPage, page int) (from, limit int) {

	if page == 0 {
		page = 1
	}

	if perPage > 100 || perPage == 0 {
		perPage = 25
	}

	return (page - 1) * perPage, perPage
}
func GetFromAndLimitWoLimit(perPage, page int) (from, limit int) {

	if page == 0 {
		page = 1
	}

	if perPage == 0 {
		perPage = 25
	}

	return (page - 1) * perPage, perPage
}
