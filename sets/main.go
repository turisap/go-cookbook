package main

type Set struct {
	elements map[any]bool
}

func NewSet() Set {
	return Set{elements: make(map[any]bool)}
}

func (s *Set) Add(el any) {
	s.elements[el] = false
}

func (s *Set) Delete(el any) {
	delete(s.elements, el)
}

func (s *Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Set) Len() int {
	return len(s.elements)
}

func (s *Set) Has(el any) bool {
	_, ok := s.elements[el]
	return ok
}

func (s *Set) Remove(el any) {
	delete(s.elements, el)
}

func (s *Set) Copy() Set {
	n := NewSet()

	for k, _ := range s.elements {
		n.Add(k)
	}

	return n
}

func Union(sets ...Set) (u Set) {
	u = sets[0].Copy()
	for _, set := range sets[1:len(sets)] {
		for k := range set.elements {
			u.Add(k)
		}
	}
	return
}

func Intersect(sets ...Set) (i Set) {
	i = sets[0].Copy()
	for k := range i.elements {
		for _, set := range sets[1:len(sets)] {
			if !set.Has(k) {
				i.Remove(k)
			}
		}
	}
	return
}

func (s *Set) List() (list []any) {
	for k, _ := range s.elements {
		list = append(list, k)
	}

	return
}

func (s Set) Difference(t Set) Set {
	for k := range t.elements {
		if s.Has(k) {
			delete(s.elements, k)
		}
	}
	return s
}

func main() {

}
