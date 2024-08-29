package stack

type MinStack struct {
	stack      [][]int // value, minimum
	CurrentMin int
}

func Constructor() MinStack {
	return MinStack{stack: make([][]int, 0)}
}

func (s *MinStack) Push(sToAdd int) {
	if !(len(s.stack) > 0) {
		s.CurrentMin = sToAdd
	}
	if sToAdd <= s.CurrentMin {
		s.stack = append(s.stack, []int{sToAdd, sToAdd})
		s.CurrentMin = sToAdd
	} else {
		s.stack = append(s.stack, []int{sToAdd, s.CurrentMin})
	}

}

func (s *MinStack) Pop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
	}
	if len(s.stack) > 0 {
		s.CurrentMin = s.FullTop()[1]
	} else {
		s.CurrentMin = 0
	}
}

func (s *MinStack) FullTop() []int {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1]
	} else {
		return nil
	}
}

func (s *MinStack) Top() int {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1][0]
	} else {
		return -1
	}
}

func (s *MinStack) GetMin() int {
	return s.FullTop()[1]
}
