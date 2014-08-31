package dna

import (
	"nucleotide"
)

type Dna struct {
	strand nucleotide.DnaStrand
}

func Init(dna *Dna, strand string) {
	dna.strand = nucleotide.DnaStrandFromString(strand)
}

func Count_ACGT(dna Dna) (int, int, int, int) {
	return Count(dna, "A"), Count(dna, "C"), Count(dna, "G"), Count(dna, "T")
}

func Count(dna Dna, s string) int {
	count := 0

	for _, el := range dna.strand {
		if string(el) == s {
			count++
		}
	}

	return count
}
