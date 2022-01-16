package structures

// MovieRead wrap data for filter movie
type MovieRead struct {
	SearchWord string `schema:"searchword"`
	Pagination int    `schema:"pagination"`
}
