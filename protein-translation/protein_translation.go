package protein

import "errors"

var (
	ErrStop        = errors.New("STOP")
	ErrInvalidBase = errors.New("Invalid codon")
)

func FromRNA(rna string) ([]string, error) {
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}

	proteins := make([]string, 0)
	var stop bool
	for i := 0; i < len(rna) && !stop; i += 3 {
		l, err := FromCodon(rna[i : i+3])

		switch {
		case err == ErrStop:
			stop = true
			continue
		case err != nil:
			return nil, err
		}

		proteins = append(proteins, l)
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
