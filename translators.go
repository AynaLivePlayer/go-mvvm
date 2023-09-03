package gmvvm

import "strconv"

type TranslatorSameType[D any] struct {
}

func (t TranslatorSameType[D]) ToModel(value D) (D, bool) {
	return value, true
}

func (t TranslatorSameType[D]) ToView(data D) D {
	return data
}

type TranslatorIntToString struct{}

func (t TranslatorIntToString) ToModel(value string) (int, bool) {
	val, err := strconv.Atoi(value)
	return val, err == nil
}

func (t TranslatorIntToString) ToView(data int) string {
	return strconv.Itoa(data)
}

type TranslatorStringToInt struct{}

func (t TranslatorStringToInt) ToModel(value int) (string, bool) {
	return strconv.Itoa(value), true
}

func (t TranslatorStringToInt) ToView(data string) int {
	val, err := strconv.Atoi(data)
	if err != nil {
		return 0
	}
	return val
}
