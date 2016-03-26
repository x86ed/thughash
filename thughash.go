package thughash

import (
	"math"
	"strconv"

	"github.com/x86ed/thughash/words"
)

func slicepos(slice []string, value string) float64 {
	for p, v := range slice {
		if v == value {
			return float64(p)
		}
	}
	return float64(-1)
}

func slicerange(slice string, index int) int {
	var output int
	output = index + 1
	return output
}

func (th ThugHash) generateRemainder(raw int64) int {
	output := int(raw) + 1000
	if output > 9999 {
		output = 9999
	}
	return output
}

//ThugHash exported struct containing all the representations of the hash
type ThugHash struct {
	Index     int
	Subject   string
	Verb      string
	End       string
	Remainder int
	QuickHash string
}

//Generate generates a thughhash struct from a given 64 bit float
func (th *ThugHash) Generate(seed float64) {
	var sm, vm float64
	sm = math.Mod(seed, 262144)
	remainder := int64(seed / 262144)

	vm = math.Mod(sm, 4096)
	subject := int64(sm / 4096)

	end := int(math.Mod(vm, 64))
	verb := int64(vm / 64)

	th.Subject = Subjects[subject]
	th.Verb = Verbs[verb]
	th.End = Ends[end]
	th.Remainder = th.generateRemainder(remainder)
	th.QuickHash = th.ThirtyTwoEncode(seed)
}

//GenerateFrom generates a thughHash struct from a string...can be either a 7 digit number or a textual representation
func (th *ThugHash) GenerateFrom(seed string) {
	var seedNum float64
	if len(seed) <= 7 {
		seedInt, _ := strconv.ParseInt(seed, 32, 64)
		seedNum = float64(seedInt)
	} else {
		seedNum = th.Degenerate(seed)
	}
	th.Generate(seedNum)
}

//Degenerate creates a thughhash struct by parsing a thughash string...is pretty cool
func (th *ThugHash) Degenerate(hash string) float64 {
	var output, pos float64
	output = 0
	pos = 0
	mutantHash := hash
	for s := range mutantHash {
		testkey := hash[:s]
		pos = slicepos(Subjects, testkey)
		if pos > -1 {
			si := slicerange(mutantHash, s)
			mutantHash = mutantHash[si:]
			break
		}
	}
	output += pos * 4096
	for v := range mutantHash {
		testkey := mutantHash[:v]
		pos = slicepos(Verbs, testkey)
		if pos > -1 {
			vi := slicerange(mutantHash, v)
			mutantHash = mutantHash[vi:]
			break
		}
	}
	output += pos * 64
	for e := range mutantHash {
		testkey := mutantHash[:e]
		pos = slicepos(Ends, testkey)
		if pos > -1 {
			ei := slicerange(mutantHash, e)
			mutantHash = mutantHash[ei:]
			break
		}
	}
	output += pos
	rem, _ := strconv.ParseFloat(mutantHash, 64)
	rem -= 1000
	rem = rem * 262144
	output += rem
	return output
}

//ThirtyTwoEncode Base32 encodes the thughash int from a number. returns a string value
func (th ThugHash) ThirtyTwoEncode(hash float64) string {
	output := strconv.FormatInt(int64(hash), 32)
	for len(output) < 7 {
		output = "0" + output
	}
	return output
}

//MakeSlug Outputs a slug from the thughash's word values
func (th *ThugHash) MakeSlug() string {
	output := th.Subject + "-" + th.Verb + "-" + th.End + "-" + strconv.Itoa(th.Remainder)
	return output
}

//MatchHash returns a base 32 string ifthe input is a thughhash and vice versa
func (th *ThugHash) MatchHash(hash string) string {
	var output string
	if len(hash) <= 7 {
		output = th.MakeSlug()
	} else {
		output = th.QuickHash
	}
	return output
}
