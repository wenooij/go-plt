package plt

type Format struct {
	seriesKey

	barFormat
}

func (f *Format) Series(key string) *Format {
	f.seriesKey = key
	return f
}

func Series(key string) *Format {
	return new(Format).Series(key)
}
