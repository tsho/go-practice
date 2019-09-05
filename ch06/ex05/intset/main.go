package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

const br = 32 << (^uint(0) >> 63)

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/br, uint(x%br)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/br, uint(x%br)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", br*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() (count int) {
	for _, word := range s.words {
		count += popcount(word)
	}
	return
}

// Return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var ret IntSet
	ret.words = make([]uint, len(s.words))
	copy(ret.words, s.words)
	return &ret
}


func popcount(x uint) (count int) {
	for x != 0 {
		count++
		x &= x - 1
	}
	return
}

func (s *IntSet) AddAll(xall ...int) {
	for _, x := range xall {
		s.Add(x)
	}
}

// Intersects two sets
func (s *IntSet) IntersectWith(t *IntSet) {
	if len(s.words) < len(t.words) {
		for i := range s.words {
			if i < len(t.words) {
				s.words[i] &= t.words[i]
			} else {
				s.words[i] = 0
			}
		}
	} else {
		for i := range t.words {
			if i < len(s.words) {
				t.words[i] &= s.words[i]
			} else {
				t.words[i] = 0
			}
		}
		s, t = t, s
	}
}

// Returns returns elements of second set without elements of second set
func (s *IntSet) DifferenceWith(t *IntSet) {
	if len(s.words) < len(t.words) {
		tmp := s.Copy()
		tmp.IntersectWith(t)
		for i := range s.words {
			s.words[i] ^= tmp.words[i]
		}
	} else {
		tmp := t.Copy()
		tmp.IntersectWith(s)
		for i := range t.words {
			t.words[i] ^= tmp.words[i]
		}
	t, s = s, t
	}
}


// Returns symmetric difference of two sets
func (s *IntSet) SymmetricDifference(t *IntSet) {
	if len(s.words) < len(t.words) {
		for i, tword := range t.words {
			if i < len(s.words) {
				s.words[i] ^= tword
			} else {
				s.words = append(s.words, tword)
			}
		}
	} else {
		for i, sword := range s.words {
			if i < len(t.words) {
				t.words[i] ^= sword
			} else {
				t.words = append(t.words, sword)
			}
		}
		s, t = t, s
	}
}

func (s IntSet) Elems() (res []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < br; j++ {
			if word&(1<<uint(j)) != 0 {
				res = append(res, br*i+j)
			}
		}
	}
	return
}