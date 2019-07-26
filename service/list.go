package service

import (
	"container/list"
	"fmt"
)

func stacktest() {
	s := list.New()
	fmt.Print("--Stack--\nPush: ")
	for i := 0; i < 10; i++ {
		s.PushBack(i) // 後ろに入れる
		fmt.Print(i, ",")
	}
	fmt.Print("\nPop : ")
	for i := 0; i < 10; i++ {
		r := s.Remove(s.Front())
		fmt.Print(r, ",")
	}
}

func queuetest() {
	s := list.New()
	fmt.Print("\n--Queue--\nEnqueue: ")
	for i := 0; i < 10; i++ {
		s.PushFront(i) // 前に入れる
		fmt.Print(i, ",")
	}
	fmt.Print("\nDequeue: ")
	for i := 0; i < 10; i++ {
		r := s.Remove(s.Front())
		fmt.Print(r, ",")
	}
}
