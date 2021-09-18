package main

import (
	"fmt"
	"time"
)

type LoopQueue struct {
	start  int
	end    int
	length int
	name   string
	data   []interface{}
}

func (lq *LoopQueue) InitQueue(length int, name string) bool {
	if nil == lq || length <= 0 {
		return false
	}
	lq.data = make([]interface{}, length)
	lq.length = length
	lq.name = name
	lq.start = 0
	lq.end = 0
	return true
}
func (lq *LoopQueue) Push(data interface{}) bool {
	if nil == lq {
		panic("LoopQueue is nil")
	}
	if lq.isFull() {
		return false
	}
	var end int = lq.getEnd()
	lq.data[end] = data
	lq.end = (end + 1) % lq.length
	return true
}
func (lq *LoopQueue) Pop() (bool, interface{}) {
	if nil == lq {
		panic("LoopQueue is nil")
	}
	if lq.isEmpty() {
		return false, nil
	}
	var start = lq.getStart()
	var startValue interface{} = lq.data[start]
	lq.start = (start + 1) % lq.length
	return true, startValue
}
func (lq *LoopQueue) isEmpty() bool {
	if nil == lq {
		panic("LoopQueue is nil")
	}
	if lq.getStart() == lq.getEnd() {
		return true
	}
	return false
}
func (lq *LoopQueue) isFull() bool {
	if nil == lq {
		panic("LoopQueue is nil")
	}
	if lq.getEnd()+1 == lq.getStart() {
		return true
	}
	return false
}
func (lq *LoopQueue) getStart() int {
	return lq.start % lq.length
}
func (lq *LoopQueue) getEnd() int {
	return lq.end % lq.length
}

var Q LoopQueue

func Create() {
	var index int = 0
	for {
		ret := Q.Push(index)
		if ret {
			fmt.Println("PushOk", "index=", index)
			index++
		} else {
			fmt.Println("PushError", "index=", index)
		}
		time.Sleep(1e9)
	}
}
func Consum() {
	for {
		ret, data := Q.Pop()
		if ret {
			fmt.Println("PopSucc", "data=", data)
		} else {
			fmt.Println("PopError")
		}
		time.Sleep(1e9)
	}
}

//实现环形队列
func main() {
	Q.InitQueue(10, "test")
	go Create()
	go Consum()
	for {
		time.Sleep(1e9)
	}
}
