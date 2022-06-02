package main

type Set map[Token]struct{}

func (s Set) Add(t Token) {
	s[t] = struct{}{}
}
