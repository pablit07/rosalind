package nucleotide

import (
	"strings"
)

type nucleotide string

func make_nucleotide(char string) nucleotide {
	n := nucleotide(char)
	m[char] = n
	return n
}

type DnaStrand []nucleotide

var m map[string]nucleotide = make(map[string]nucleotide)

var (
	A nucleotide = make_nucleotide("A")
	C nucleotide = make_nucleotide("C")
	G nucleotide = make_nucleotide("G")
	T nucleotide = make_nucleotide("T")
)

func NewDnaStrand(nts ...nucleotide) DnaStrand {
	return nts
}

func DnaStrandFromString(s string) DnaStrand {
	chars := strings.SplitN(s, "", len(s))

	strand := make([]nucleotide, 0, 0)

	for _, el := range chars {
		n := m[el]
		strand = append(strand, n)
	}

	return strand
}
