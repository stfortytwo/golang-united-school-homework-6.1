package golang_united_school_homework

import "errors"

const (
	errMesIndexOutRange = "index doesn't exist or out of the range"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity <= len(b.shapes) {
		return errors.New("no room left in the box")
	} else {
		b.shapes = append(b.shapes, shape)
		return nil
	}
	//panic("implement me")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes)-1 < i || i < 0 {
		return nil, errors.New(errMesIndexOutRange)
	} else {
		return b.shapes[i], nil
	}
	//panic("implement me")
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i > len(b.shapes)-1 || i < 0 {
		return nil, errors.New(errMesIndexOutRange)
	} else {
		returnShape := b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return returnShape, nil
	}
	//panic("implement me")

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i > len(b.shapes)-1 || i < 0 {
		return nil, errors.New(errMesIndexOutRange)
	} else {
		returnShape := b.shapes[i]
		b.shapes[i] = shape
		return returnShape, nil
	}
	//panic("implement me")
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sumPerimeter float64 = 0
	for i := range b.shapes {
		sumPerimeter += b.shapes[i].CalcPerimeter()
	}
	return sumPerimeter
	//panic("something went really wrong")
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sumArea float64 = 0
	for i := range b.shapes {
		sumArea += b.shapes[i].CalcArea()
	}
	return sumArea
	//panic("something went really wrong")
}

// RemoveAllCircles removes all circles in the list;
// whether circles are not exist in the list returns an error
func (b *box) RemoveAllCircles() error {
	var removedCirclesCounter int = 0
	for i := 0; i < len(b.shapes); {
		if _, ok := b.shapes[i].(Circle); ok {
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
			removedCirclesCounter++
		} else {
			i++
		}
	}
	if removedCirclesCounter == 0 {
		return errors.New("there are no circles to annihilate")
	} else {
		return nil
	}
	//panic("implement me")
}
