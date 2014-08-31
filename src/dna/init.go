package dna

import (
	"strings"
)

type Dna struct {
	strand string
}

func Init(dna *Dna, strand string) {
	dna.strand = strand
}

func Count_ACGT(dna Dna) (int, int, int, int) {
	return Count(dna, "A"), Count(dna, "C"), Count(dna, "G"), Count(dna, "T")
}

func Count(dna Dna, s string) int {
	return strings.Count(dna.strand, s)
}
