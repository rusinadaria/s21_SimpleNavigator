package stack

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{data: []int{}}
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Pop() (int, bool){
	if len(s.data) == 0 {
        return 0, false
    }
    
    n := len(s.data) - 1
    value := s.data[n]
    s.data = s.data[:n]

    return value, true 
}

func (s *Stack) Top() int {
	return len(s.data) - 1
}