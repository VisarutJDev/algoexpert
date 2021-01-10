package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ShiftLinkedList(head *LinkedList, k int) *LinkedList {
	// Write your code here.
	if k == 0 {
		return head
	}
	counter := 0
	headToCount := head
	var array []int
	for {
		if headToCount.Next == nil {
			counter++
			array = append(array, headToCount.Value)
			break
		}
		counter++
		array = append(array, headToCount.Value)
		headToCount = headToCount.Next
	}

	var newArray []int

	if k > 0 {
		k = k % len(array)

		lasttoHead := array[(len(array) - k):]
		headTolast := array[:len(array)-k]

		newArray = append(lasttoHead, headTolast...)
	} else {
		k = k * -1 % len(array)
		headToLast := array[:k]
		lastToHead := array[k:]

		newArray = append(lastToHead, headToLast...)
	}

	newHead := &LinkedList{}
	for i := range newArray {
		if i == 0 {
			newHead.Value = newArray[i]
			newHead.Next = &LinkedList{}
			continue
		}
		recur := newHead.Next
		j := 1
		for {
			if i == j {
				if i != len(newArray)-1 {
					recur.Next = &LinkedList{}
				}
				recur.Value = newArray[i]
				break
			}
			if recur.Next == nil {
				recur.Next = &LinkedList{}
			}
			recur = recur.Next
			j++
		}

	}
	return newHead
}
