package magic8ball

/*
MIT License

Copyright (c) 2022 Eric Fialkowski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

import (
	"math/rand"
	"testing"
)

func TestMagic8Ball_Answer(t *testing.T) {
	mb := New()
	a := mb.Answer()
	if !isPossibleResponse(a) {
		t.Failed()
	}
}

func TestMagic8Ball_Provided_Rand_Answer(t *testing.T) {
	mb := NewWithRandom(*rand.New(rand.NewSource(1)))

	// since the rand source is fixed, it should always return the same value
	if mb.Answer() != "It is decidedly so." {
		t.Fail()
	}
}

// In case we ever want more tests
func isPossibleResponse(a string) bool {
	var pass bool
	switch a {
	case "It is certain.":
		pass = true
	case "It is decidedly so.":
		pass = true
	case "Without a doubt.":
		pass = true
	case "Yes - definitely.":
		pass = true
	case "You may rely on it.":
		pass = true
	case "As I see it, yes.":
		pass = true
	case "Most likely.":
		pass = true
	case "Outlook good.":
		pass = true
	case "Yes.":
		pass = true
	case "Signs point to yes.":
		pass = true
	case "Reply hazy, try again.":
		pass = true
	case "Ask again later.":
		pass = true
	case "Better not tell you now.":
		pass = true
	case "Cannot predict now.":
		pass = true
	case "Concentrate and ask again.":
		pass = true
	case "Don't count on it.":
		pass = true
	case "My reply is no.":
		pass = true
	case "My sources say no.":
		pass = true
	case "Outlook not so good.":
		pass = true
	case "Very doubtful.":
		pass = true
	default:
		pass = false
	}
	return pass
}
