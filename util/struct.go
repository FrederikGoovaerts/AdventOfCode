package util

type StringSet map[string]struct{}

func (s *StringSet) Has(in string) bool {
	_, present := (*s)[in]
	return present
}

func (s *StringSet) Add(in string) {
	(*s)[in] = EMPTY_STRUCT
}

func (s *StringSet) Delete(in string) {
	delete(*s, in)
}

var EMPTY_STRUCT = struct{}{}
