package stacks

/* This is a reimplementation of
   github.com/golang-collections/collections/stack to better suit my
   use case, so in order to comply with the MIT license (as this
   package contains a significant portion of that package), the
   license, including the copyright notice, for that package is
   included in the file COLLECTIONS_LICENSE */

// FloatStack definitions
type (
	FloatStack struct {
		Top    *FloatNode
		length int
	}
	FloatNode struct {
		Value    float64
		Previous *FloatNode
	}
)

// Create a new stack
func NewFloatStack() *FloatStack {
	return &FloatStack{nil, 0}
}

// Return the number of items in the stack
func (this *FloatStack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *FloatStack) Peek() (float64, error) {
	if this.length == 0 {
		return 0, ErrorEmptyStack
	}
	return this.Top.Value, nil
}

// Pop the top item of the stack and return it
func (this *FloatStack) Pop() (float64, error) {
	if this.length == 0 {
		return 0, ErrorEmptyStack
	}

	n := this.Top
	this.Top = n.Previous
	this.length--
	return n.Value, nil
}

// Push a value onto the top of the stack
func (this *FloatStack) Push(value float64) {
	n := &FloatNode{value, this.Top}
	this.Top = n
	this.length++
}
