package multivalues

func First(values ...interface{}) interface{} {
	if len(values) == 0 {
		return nil
	}
	return values[0]
}

func Last(values ...interface{}) interface{} {
	if len(values) == 0 {
		return nil
	}
	return values[len(values)-1]
}

func ToSlices(values ...interface{}) []interface{} {
	return values
}
