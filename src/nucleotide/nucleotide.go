package nucleotide

import (
	"strings"
)

type nucleotide string

func make_nucleotide(n nucleotide) {
	m[string(n)] = n
}

type DnaStrand []nucleotide
type RnaStrand []nucleotide

var m map[string]nucleotide = make(map[string]nucleotide)

const (
	A nucleotide = "A"
	C nucleotide = "C"
	G nucleotide = "G"
	T nucleotide = "T"
	U nucleotide = "U"
)

func init() {
	make_nucleotide(A)
	make_nucleotide(C)
	make_nucleotide(G)
	make_nucleotide(T)
	make_nucleotide(U)
}

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

func GetCompliment(n nucleotide) (result nucleotide, fail bool) {

	switch n {
	case A:
		result = T
	case T:
		result = A
	case G:
		result = C
	case C:
		result = G
	default:
		fail = true
	}

	return
}

func GetComplimentFromStrand(strand DnaStrand) (result DnaStrand) {

	for _, n := range strand {
		compliment, _ := GetCompliment(n)
		result = append(result, compliment)
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return
}
