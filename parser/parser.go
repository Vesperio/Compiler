package parser

import (
	. "github.com/deckarep/golang-set/v2"
)

const Epsilon = 'ε'

// Production 产生式
type Production struct {
	L rune   // 左部
	R []rune // 右部
}

// Grammar 文法
type Grammar struct {
	T     Set[rune]    // 所有的终结符
	N     Set[rune]    // 所有的非终结符
	Prods []Production // 产生式
}

// 构造函数
func NewGrammar() *Grammar {
	return &Grammar{
		T:     NewThreadUnsafeSet[rune](),
		N:     NewThreadUnsafeSet[rune](),
		Prods: make([]Production, 0),
	}
}

// NULLABLE 集
var Nullable = NewThreadUnsafeSet[rune]()

// FIRST 集
var First = make(map[rune]Set[rune])

// FOLLOW 集
var Follow = make(map[rune]Set[rune])

// Init NULLABLE 集、FIRST 集、FOLLOW 集
func (g Grammar) Init() {
	Nullable.Clear()

	it := g.N.Iterator()
	for e := range it.C {
		First[e] = NewThreadUnsafeSet[rune]()
	}
	it.Stop()

	it = g.N.Iterator()
	for e := range it.C {
		Follow[e] = NewThreadUnsafeSet[rune]()
	}
	it.Stop()
}

// 求 NULLABLE 集
func (g Grammar) GetNullableSet() {
	isChanging := true
	for isChanging {
		isChanging = false
		// 遍历所有产生式
		for _, p := range g.Prods {
			if p.R[0] == Epsilon {
				if Nullable.Contains(p.L) == false {
					Nullable.Add(p.L)
					isChanging = true
				}
				continue
			}

			flag := true
			// 遍历产生式右部的所有元素
			for _, e := range p.R {
				if !(g.N.Contains(e) && Nullable.Contains(e)) {
					flag = false
					break
				}
			}
			if flag == true {
				if Nullable.Contains(p.L) == false {
					Nullable.Add(p.L)
					isChanging = true
				}
			}
		}
	}
}

// 求 FIRST 集
func (g Grammar) GetFirstSet() {
	isChanging := true
	for isChanging {
		isChanging = false
		// 遍历每条产生式
		for _, p := range g.Prods {
			// 遍历该条产生式中的每个元素
			for _, e := range p.R {
				if g.T.Contains(e) { // 若该元素是终结符
					if First[p.L].Contains(e) == false {
						First[p.L].Add(e)
						isChanging = true
					}
					break
				} else if g.N.Contains(e) { // 若该元素是非终结符
					if First[p.L].IsSuperset(First[e]) == false {
						First[p.L] = First[p.L].Union(First[e])
						isChanging = true
					}
					if Nullable.Contains(e) == false {
						break
					}
				}
			}
		}
	}
}

// 求 FOLLOW 集
func (g Grammar) GetFollowSet() {
	isChanging := true
	for isChanging {
		isChanging = false
		// 遍历每条产生式
		for _, p := range g.Prods {
			tmp := Follow[p.L].Clone()
			for i := len(p.R) - 1; i >= 0; i-- { // 逆序遍历该条产生式中的每个元素
				e := p.R[i]
				if g.T.Contains(e) {
					tmp.Clear()
					tmp.Add(e)
				} else if g.N.Contains(e) {
					if Follow[e].IsSuperset(tmp) == false {
						Follow[e] = Follow[e].Union(tmp)
						isChanging = true
					}

					if Nullable.Contains(e) == false {
						tmp.Clear()
					}
					tmp = tmp.Union(First[e])
				}
			}
		}
	}
}

// 针对每一条产生式计算 FIRST_S 集
func (g Grammar) GetFirstS(p Production) (firstS Set[rune]) {
	firstS = NewThreadUnsafeSet[rune]()
	for _, e := range p.R {
		if g.T.Contains(e) {
			firstS.Add(e)
			return
		} else if g.N.Contains(e) {
			firstS = firstS.Union(First[e])
			if Nullable.Contains(e) == false {
				return
			}
		}
	}
	firstS = firstS.Union(Follow[p.L])
	return
}
