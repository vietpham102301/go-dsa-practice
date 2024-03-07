package main

import "math/rand"

type RandomizedSet struct {
	valueArray []int
	setMap     map[int]int
}

func Constructor() RandomizedSet {
	obj := RandomizedSet{}
	obj.setMap = make(map[int]int)
	obj.valueArray = []int{}
	return obj
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.setMap[val]
	if !ok {
		this.valueArray = append(this.valueArray, val)
		l := len(this.valueArray) - 1
		this.setMap[val] = l
		return true
	}

	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	l := len(this.valueArray) - 1
	index, ok := this.setMap[val]
	if ok {
		lastValue := this.valueArray[l]
		this.valueArray[index] = lastValue
		this.setMap[lastValue] = index
		this.valueArray = this.valueArray[:l]
		delete(this.setMap, val)
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	l := len(this.valueArray)
	index := rand.Intn(l)
	return this.valueArray[index]
}

func main() {
	//Write your test here!
}
