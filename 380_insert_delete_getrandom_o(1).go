package main

import "math/rand"

// https://leetcode.com/problems/insert-delete-getrandom-o1/description/

type RandomizedSet struct {
	m map[int]int
	a []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.m[val]
	if !ok {
		this.a = append(this.a, val)
		this.m[val] = len(this.a) - 1
	}
	// fmt.Println("Added ", val)
	// fmt.Println("a = ", this.a)
	// fmt.Println("m = ", this.m)
	return !ok
}

func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.m[val]
	if ok {
		delete(this.m, val)
		newV := this.a[len(this.a)-1]
		this.a[i] = newV
		if newV != val {
			this.m[newV] = i
		}
		this.a = this.a[:len(this.a)-1]
		// fmt.Println("Removed ", val)
		// fmt.Println("a = ", this.a)
		// fmt.Println("m = ", this.m)
	}
	return ok
}

func (this *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(this.a)) // Element at which position we want
	return this.a[randomIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
