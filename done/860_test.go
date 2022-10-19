package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "860"
  title: design-circular-queue
  lang: golang
  type: large
  inputString: |-
    ["MyCircularQueue","enQueue","enQueue","enQueue","enQueue","Rear","isFull","deQueue","enQueue","Rear"]
    [[3],[1],[2],[3],[4],[],[],[],[4],[]]
*/

type MyCircularQueue struct {
	size   int
	length int
	q      []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{size: k, q: make([]int, k)}
}

func (q *MyCircularQueue) EnQueue(v int) bool {
	if q.IsFull() {
		return false
	}
	q.q[q.length] = v
	q.length++
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	copy(q.q, q.q[1:])
	q.length--
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.q[0]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.q[q.length-1]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.length == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.length == q.size
}

func Test_860(t *testing.T) {
	type Data struct {
		Methods []string
		Values  [][]int
	}
	NewTestcases(t).
		AddExpectation([]any{nil, true, 1, 1, true, 1, 2, true, 2, 2, true, 2, 3}).
		AddInput(Data{
			Methods: MakeStringSlice(`["MyCircularQueue","enQueue","Front","Rear","enQueue","Front","Rear","deQueue","Front","Rear","enQueue","Front","Rear"]`),
			Values:  Make2DMatrixInt("[[2],[1],[],[],[2],[],[],[],[],[],[3],[],[]]"),
		}).
		AddExpectation([]any{nil, true, false, false, true, false, 3, 3, true, -1, -1}).
		AddInput(Data{
			Methods: MakeStringSlice(`["MyCircularQueue","isEmpty","isFull","deQueue","enQueue","enQueue","Front","Rear","deQueue","Front","Rear"]`),
			Values:  Make2DMatrixInt("[[1],[],[],[],[3],[2],[],[],[],[],[]]"),
		}).
		AddExpectation([]any{nil, true, true, true, false, 3, true, true, true, 4}).
		AddInput(Data{
			Methods: MakeStringSlice(`["MyCircularQueue","enQueue","enQueue","enQueue","enQueue","Rear","isFull","deQueue","enQueue","Rear"]`),
			Values:  Make2DMatrixInt("[[3],[1],[2],[3],[4],[],[],[],[4],[]]"),
		}).
		Each(func(a *assert.Assertions, td TestData) {
			var actual []any
			input := td.Input.(Data)
			methods, values := input.Methods, input.Values
			var q MyCircularQueue
			for i := range input.Methods {
				var v any
				switch methods[i] {
				case "MyCircularQueue":
					q = Constructor(values[i][0])
				case "enQueue":
					v = q.EnQueue(values[i][0])
				case "deQueue":
					v = q.DeQueue()
				case "Front":
					v = q.Front()
				case "Rear":
					v = q.Rear()
				case "isEmpty":
					v = q.IsEmpty()
				case "isFull":
					v = q.IsFull()
				}
				actual = append(actual, v)
			}

			a.Equal(td.Expectation, actual)
		})
}
