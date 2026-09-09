package main

type Deque struct {
	data []int
}

func (d *Deque) IsEmpty() bool {
	return len(d.data) == 0
}

func (d *Deque) Front() int {
	return d.data[0]
}

func (d *Deque) Back() int {
	return d.data[len(d.data)-1]
}

func (d *Deque) PushBack(val int) {
	d.data = append(d.data, val)
}

func (d *Deque) PopFront() {
	d.data = d.data[1:]
}

func (d *Deque) PopBack() {
	d.data = d.data[:len(d.data)-1]
}

func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) == 0 {
		return nil
	}

	dq := &Deque{}
	result := []int{}

	for i := 0; i < len(nums); i++ {

		// Remove indices outside current window
		if !dq.IsEmpty() && dq.Front() <= i-k {
			dq.PopFront()
		}

		// Remove smaller elements from back
		// explaination: every other element in the new window when added this new value, if its lesser we are removing it bcz new one is max for the window
		for !dq.IsEmpty() && nums[i] >= nums[dq.Back()] {
			dq.PopBack()
		}

		// Add current index
		dq.PushBack(i)

		// Window is ready
		if i >= k-1 {
			result = append(result, nums[dq.Front()])
		}
	}

	return result
}
