package src

// import "dna"
import (
	"nucleotide"
	"testing"
)

func TestGetCompliment1(t *testing.T) {
	actual, _ := nucleotide.GetCompliment(nucleotide.A)

	if actual != nucleotide.T {
		t.Errorf("Failed, actual: %v, expected %v", actual, nucleotide.T)
	}

}

func TestGetCompliment2(t *testing.T) {
	actual, _ := nucleotide.GetCompliment(nucleotide.T)

	if actual != nucleotide.A {
		t.Errorf("Failed, actual: %v, expected %v", actual, nucleotide.A)
	}

}

func TestGetCompliment3(t *testing.T) {
	actual, _ := nucleotide.GetCompliment(nucleotide.C)

	if actual != nucleotide.G {
		t.Errorf("Failed, actual: %v, expected %v", actual, nucleotide.G)
	}

}

func TestGetCompliment4(t *testing.T) {
	actual, _ := nucleotide.GetCompliment(nucleotide.G)

	if actual != nucleotide.C {
		t.Errorf("Failed, actual: %v, expected %v", actual, nucleotide.C)
	}
}

func TestGetComplimentFromStrand(t *testing.T) {
	input, expected := nucleotide.DnaStrandFromString("AAAACCCGGT"), "ACCGGGTTTT"
	actual := nucleotide.StrandToString(
		nucleotide.GetComplimentFromStrand(input))

	if actual != expected {
		t.Errorf("Failed, actual: %v, expected: %v", actual, expected)
	}
}
