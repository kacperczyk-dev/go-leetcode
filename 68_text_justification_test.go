package main

import (
	"reflect"
	"testing"
)

func TestFullJustify(t *testing.T) {
	type In struct {
		words    []string
		maxWidth int
	}
	cases := []struct {
		in   In
		want []string
	}{
		{In{[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20}, []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "}},
		{In{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16}, []string{"What   must   be", "acknowledgment  ", "shall be        "}},
		{In{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16}, []string{"This    is    an", "example  of text", "justification.  "}},
	}

	for _, c := range cases {
		got := fullJustify(c.in.words, c.in.maxWidth)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("fullJustify(%v, %v) == %v, want %v", c.in.words, c.in.maxWidth, got, c.want)
		}
	}
}
