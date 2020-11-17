package ch03

import "errors"

type ThreeStack struct {
	array                              []string
	width1, width2, width3, i1, i2, i3 int
}

func NewThreeStack() *ThreeStack {
	return &ThreeStack{
		array:  make([]string, 30),
		width1: 10,
		width2: 10,
		width3: 10,
		i1:     -1,
		i2:     -1,
		i3:     -1,
	}
}

func (t *ThreeStack) Push1(s string) error {
	if s == "" {
		return errors.New("empty string is not allowed")
	}
	if t.i1+1 < t.width1 {
		t.i1++
		t.array[t.i1] = s
	} else {
		old := t.array
		newArray := make([]string, len(old)+5)
		for i, s := range old {
			if i < t.width1 {
				newArray[i] = s
			} else if (i <= (t.width1 + t.width2)) && (i >= t.width1) {
				newArray[i+5] = s
			} else if (i <= (t.width3 + t.width2 + t.width1)) && (i >= (t.width2 + t.width1)) {
				newArray[i+5] = s
			}
		}
		t.array = newArray
		t.width1 = t.width1 + 5
		t.i1++
		t.array[t.i1] = s

	}
	return nil
}
func (t *ThreeStack) Push2(s string) error {
	if s == "" {
		return errors.New("empty string is not allowed")
	}
	if t.i2+1 < t.width2 {
		t.i2++
		t.array[t.width1+t.i2] = s
	} else {
		old := t.array
		newArray := make([]string, len(old)+5)
		for i, s := range old {
			if i < t.width1 {
				newArray[i] = s
			} else if (i < (t.width1 + t.width2)) && (i >= t.width1) {
				newArray[i] = s
			} else if (i < (t.width3 + t.width2 + t.width1)) && (i >= (t.width1 + t.width2)) {
				newArray[i+5] = s
			}
		}
		t.array = newArray
		t.width2 = t.width2 + 5
		t.i2++
		t.array[t.width1+t.i2] = s
	}
	return nil
}
func (t *ThreeStack) Push3(s string) error {
	if s == "" {
		return errors.New("empty string is not allowed")
	}
	if t.i3+1 < t.width3 {
		t.i3++
		t.array[t.width2+t.width1+t.i3] = s
	} else {
		old := t.array
		newArray := make([]string, len(old)+5)

		for i, s := range old {
			if i < t.i1 {
				newArray[i] = s
			} else if (i < (t.width1 + t.width2)) && (i >= t.width1) {
				newArray[i] = s
			} else if (i < (t.width3 + t.width2 + t.width1)) && (i >= (t.width1 + t.width2)) {
				newArray[i] = s
			}
		}
		t.array = newArray
		t.width3 = t.width3 + 5
		t.i3++
		t.array[t.width2+t.width1+t.i3] = s
	}
	return nil
}
func (t *ThreeStack) Pop1() (string, error) {
	if t.array[0] == "" {
		return "", errors.New("empty")
	}
	r := t.array[t.i1]
	t.array[t.i1] = ""
	t.i1--
	return r, nil
}
func (t *ThreeStack) Peek1() (string, error) {
	if t.array[0] == "" {
		return "", errors.New("empty")
	}
	return t.array[t.i1], nil
}
func (t *ThreeStack) Pop2() (string, error) {
	if t.array[t.width1] == "" {
		return "", errors.New("empty")
	}
	r := t.array[t.i2+t.width1]
	t.array[t.i2+t.width1] = ""
	t.i2--
	return r, nil
}
func (t *ThreeStack) Peek2() (string, error) {
	if t.array[t.width1] == "" {
		return "", errors.New("empty")
	}
	return t.array[t.i2+t.width1], nil
}
func (t *ThreeStack) Pop3() (string, error) {
	if t.array[t.width2+t.width2] == "" {
		return "", errors.New("empty")
	}
	r := t.array[t.width2+t.width2+t.i3]
	t.array[t.width2+t.width2+t.i3] = ""
	t.i3--
	return r, nil
}
func (t *ThreeStack) Peek3() (string, error) {
	if t.array[t.width2+t.width2] == "" {
		return "", errors.New("empty")
	}
	return t.array[t.width2+t.width2+t.i3], nil
}
func (t *ThreeStack) IsEmpty1() bool {
	return t.array[0] == ""
}
func (t *ThreeStack) IsEmpty2() bool {
	return t.array[t.width1] == ""
}
func (t *ThreeStack) IsEmpty3() bool {
	return t.array[t.width2+t.width1] == ""
}
