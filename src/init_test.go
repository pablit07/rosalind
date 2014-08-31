package src

import (
	"dna"
	"nucleotide"
	"testing"
)

func TestCount(t *testing.T) {
	_dna := dna.Dna{}

	dna.Init(&_dna, "AAA")

	count_A := dna.Count(_dna, "A")

	if count_A != 3 {
		t.Errorf("Count of A was not 3, actual %d, tested on %q", count_A, _dna)
	}
}

func TestCount2(t *testing.T) {
	_dna := dna.Dna{}

	dna.Init(&_dna, "ACCAGGGTA")

	count_A := dna.Count(_dna, "A")

	if count_A != 3 {
		t.Errorf("Count of A was not 3, actual %d, tested on %q", count_A, _dna)
	}
}

func TestCountAll(t *testing.T) {
	_dna := dna.Dna{}

	dna.Init(&_dna, "ACCAGGGTA")

	count_A, count_C, count_G, count_T := dna.Count_ACGT(_dna)

	if count_A != 3 {
		t.Errorf("Count of A was not 3, actual %d, tested on %q", count_A, _dna)
	}
	if count_C != 2 {
		t.Errorf("Count of C was not 2, actual %d, tested on %q", count_C, _dna)
	}
	if count_G != 3 {
		t.Errorf("Count of G was not 3, actual %d, tested on %q", count_G, _dna)
	}
	if count_T != 1 {
		t.Errorf("Count of T was not 1, actual %d, tested on %q", count_T, _dna)
	}
}

func TestCountAll2(t *testing.T) {
	_dna := dna.Dna{}

	dna.Init(&_dna, "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC")

	count_A, count_C, count_G, count_T := dna.Count_ACGT(_dna)

	if count_A != 20 {
		t.Errorf("Count of A was not 20, actual %d, tested on %q", count_A, _dna)
	}
	if count_C != 12 {
		t.Errorf("Count of C was not 12, actual %d, tested on %q", count_C, _dna)
	}
	if count_G != 17 {
		t.Errorf("Count of G was not 17, actual %d, tested on %q", count_G, _dna)
	}
	if count_T != 21 {
		t.Errorf("Count of T was not 21, actual %d, tested on %q", count_T, _dna)
	}
}

func TestMakeDnaStrand(t *testing.T) {
	nucleotide.NewDnaStrand(nucleotide.A, nucleotide.C, nucleotide.G, nucleotide.T)
}

func TestMakeDnaStrand2(t *testing.T) {
	strand := nucleotide.DnaStrandFromString("AAA")

	if len(strand) != 3 {
		t.Errorf("Length of strand was not 3, actual=%d", len(strand))
	}
}

func TestStrandToString(t *testing.T) {
	strand := nucleotide.DnaStrandFromString("GATGGAACTTGACTACGTAAATT")

	actual := nucleotide.StrandToString(strand)

	if actual != "GATGGAACTTGACTACGTAAATT" {
		t.Errorf("Error, expected %q but was %q", "GATGGAACTTGACTACGTAAATT", actual)
	}

}

func TestDnaToRna(t *testing.T) {
	strand := nucleotide.DnaStrandFromString("GATGGAACTTGACTACGTAAATT")

	actual := nucleotide.StrandToString(nucleotide.RnaStrandFromDnaStrand(strand))

	if actual != "GAUGGAACUUGACUACGUAAAUU" {
		t.Errorf("Error, expected %q but was %q", "GAUGGAACUUGACUACGUAAAUU", actual)
	}
}
