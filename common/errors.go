package common


type DataError struct {
	Code int
	Err string
}

func (e DataError) Error() string {
	return e.Err
}
