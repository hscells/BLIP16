package basm

import (
	"bufio"
	"github.com/hscells/blip16/blip"
	"io"
	"strconv"
	"strings"
)

type token string

const (
	nop  token = "nop"
	jp   token = "jp"
	jpeq token = "jpeq"
	jplt token = "jplt"
	jpgt token = "jpgt"
	mov  token = "mov"
	lda  token = "lda"
	inc  token = "inc"
	dec  token = "dec"
	rda  token = "rda"
	add  token = "add"
	sub  token = "sub"
	mul  token = "mul"
	div  token = "div"
	and  token = "and"
	or   token = "or"
	xor  token = "xor"
	sla  token = "sla"
	sra  token = "sra"
)

var mapping = map[token]uint8{
	nop: 0x00, jp: 0x01, jpeq: 0x02, jplt: 0x03, jpgt: 0x04,
	mov: 0x10, lda: 0x11, inc: 0x12, dec: 0x13, rda: 0x14,
	add: 0x20, sub: 0x21, mul: 0x22, div: 0x22,
	and: 0x30, or: 0x31, xor: 0x32, sla: 0x33, sra: 0x34,
}

func Tokenise(reader io.Reader) []token {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)
	var tok token
	var tokens []token
	for scanner.Scan() {
		r := rune(scanner.Bytes()[0])
		if r == '\r' || r == 10 || r == 32 {
			if len(tok) > 0 {
				tokens = append(tokens, tok)
				tok = ""
			}
			continue
		}
		tok += token(r)
	}
	return tokens
}

func Parse(toks []token) ([]uint8, error) {
	data := make([]uint8, 0xffff-blip.OffCD)
	vars := make(map[token]uint16)
	labels := make(map[token]uint16)
	labelsLocs := make(map[uint16]token)
	var state int
	var varTok token
	var i uint16
	for _, tok := range toks {
		// Check if the token is an op code.
		if op, ok := mapping[tok]; ok {
			data[i] = op
			i++
			continue
		}
		// Check if the token is a label.
		if tok[0] == '$' {
			if _, ok := vars[tok]; !ok {
				varTok = tok
				state = 1
			} else {
				data[i] = uint8(vars[tok] >> 8)
				data[i+1] = uint8(vars[tok])
				i += 2
			}
			goto update
		}

		if tok[0] == ':' {
			// The current token is be the definition
			// of a label.
			labels[tok] = blip.OffCD + i
			goto update
		}

		if tok[0] == '.' {
			labelsLocs[i] = tok
			i += 2
			goto update
		}

		// Now decide what to do with the number.
		switch state {
		case 0:
			// Otherwise the token must be an 8-bit hex number.
			// BitSize must be doubled because ParseInt uses signed integers.
			num, err := strconv.ParseInt(string(tok), 16, 16)
			if err != nil {
				return nil, err
			}
			data[i] = uint8(num)
			i++
		case 1:
			// Otherwise the token must be an 16-bit hex number.
			// BitSize must be doubled because ParseInt uses signed integers.
			num, err := strconv.ParseInt(string(tok), 16, 32)
			if err != nil {
				return nil, err
			}
			// Labels must be offset.
			vars[varTok] = uint16(num)
			state = 0
		}
	update:
	}

	//fmt.Println(labels)
	//fmt.Println(labelsLocs)

	// Now, finally the remaining labels
	// that were not defined in the first pass
	// can now be substituted.
	for addr, tok := range labelsLocs {
		// Retrieve the address of the label definition.
		ptr := labels[token(strings.Replace(string(tok), ".", ":", -1))]
		//fmt.Printf("%s(%04x)->%04x[%02x;%02x]\n", tok, addr, ptr, uint8(ptr>>8), uint8(ptr))
		// Substitute the label for the pointer.
		data[addr] = uint8(ptr >> 8)
		data[addr+1] = uint8(ptr)
	}

	return data, nil
}
