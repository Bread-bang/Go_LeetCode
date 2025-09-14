func removeDuplicates(nums []int) int {
    k := 0
    length := len(nums)

    set := NewSet[int]()

    for _, v := range nums {
        if !set.Contains(v) {            
			set.Add(v)
			nums[k] = v
            k++
        }
    }

    for i := k; i < length; i++ {
        nums[i] = 0
    }

    return k
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](elements ...T) Set[T] {
    s := make(Set[T])
    for _, el := range elements {
        s.Add(el)
    }

    return s
}

func (s Set[T]) Add(element T) {
    s[element] = struct{}{}
}

func (s Set[T]) Contains(element T) bool {
    _, found := s[element]
    return found
}