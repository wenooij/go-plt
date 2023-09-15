package plt

type ByKey []Key

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i] < a[j] }

type ByRecord []Record

func (a ByRecord) Len() int           { return len(a) }
func (a ByRecord) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRecord) Less(i, j int) bool { return a[i].RelTime < a[j].RelTime }

type ByLineRecord []LineRecord

func (a ByLineRecord) Len() int           { return len(a) }
func (a ByLineRecord) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLineRecord) Less(i, j int) bool { return a[i].RelTime < a[j].RelTime }

func IntegrateKeys(keys []Key, records []Record) map[Key]float64 {
	ints := make(map[Key]float64, len(keys))
	for _, r := range records {
		ints[r.Key] += r.Value
	}
	return ints
}

type ByIntegratedKey struct {
	Keys []Key
	Ints map[Key]float64
}

func (a ByIntegratedKey) Len() int           { return len(a.Keys) }
func (a ByIntegratedKey) Swap(i, j int)      { a.Keys[i], a.Keys[j] = a.Keys[j], a.Keys[i] }
func (a ByIntegratedKey) Less(i, j int) bool { return a.Ints[a.Keys[j]] < a.Ints[a.Keys[i]] }
