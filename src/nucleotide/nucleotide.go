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
type RnaStrand []nucleotide

var m map[string]nucleotide = make(map[string]nucleotide)

var (
	A nucleotide = make_nucleotide("A")
	C nucleotide = make_nucleotide("C")
	G nucleotide = make_nucleotide("G")
	T nucleotide = make_nucleotide("T")
	U nucleotide = make_nucleotide("U")
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

func RnaStrandFromDnaStrand(dnaStrand DnaStrand) RnaStrand {
	rnaStrand := make([]nucleotide, 0, 0)

	for _, el := range dnaStrand {
		if el == T {
			rnaStrand = append(rnaStrand, U)
		} else {
			rnaStrand = append(rnaStrand, el)
		}
	}

	return rnaStrand
}

func StrandToString(strand []nucleotide) string {
	strandChars := make([]string, 0, 0)
	for _, el := range strand {
		strandChars = append(strandChars, string(el))
	}
	return strings.Join(strandChars, "")
}
