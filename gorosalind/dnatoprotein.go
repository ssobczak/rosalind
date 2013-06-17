package gorosalind

var codonMap = map[string]byte{
	"ATT": 'I', "ATC": 'I', "ATA": 'I', "ATG": 'M', "ACT": 'T', "ACC": 'T', "ACA": 'T', "ACG": 'T',
	"AAT": 'N', "AAC": 'N', "AAA": 'K', "AAG": 'K', "AGT": 'S', "AGC": 'S', "AGA": 'R', "AGG": 'R',
	"CTT": 'L', "CTC": 'L', "CTA": 'L', "CTG": 'L', "CCT": 'P', "CCC": 'P', "CCA": 'P', "CCG": 'P',
	"CAT": 'H', "CAC": 'H', "CAA": 'Q', "CAG": 'Q', "CGT": 'R', "CGC": 'R', "CGA": 'R', "CGG": 'R',
	"GTT": 'V', "GTC": 'V', "GTA": 'V', "GTG": 'V', "GCT": 'A', "GCC": 'A', "GCA": 'A', "GCG": 'A',
	"GAT": 'D', "GAC": 'D', "GAA": 'E', "GAG": 'E', "GGT": 'G', "GGC": 'G', "GGA": 'G', "GGG": 'G',
	"TTT": 'F', "TTC": 'F', "TTA": 'L', "TTG": 'L', "TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
	"TAT": 'Y', "TAC": 'Y', "TGT": 'C', "TGC": 'C', "TGG": 'W',
}

func dnaCodonToProtein(codon []byte) (ret byte, ok bool) {
	ret, ok = codonMap[string(codon)]
	return
}

func isStart(codon []byte) bool {
	return string(codon) == "ATG"
}

func isStop(codon []byte) bool {
	cs := string(codon)
	return cs == "TAA" || cs == "TAG" || cs == "TGA"
}

func reverseCompl(codon []byte) (ret []byte) {
	ret = make([]byte, len(codon))
	for k, v := range codon {
		ret[len(codon)-k] = v
	}
	return
}
