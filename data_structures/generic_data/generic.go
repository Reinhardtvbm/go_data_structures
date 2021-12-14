package rg

type T struct {
	Int     *int
	Int8    *int8
	Int16   *int16
	Int32   *int32
	Int64   *int64
	String  *string
	Char    *byte
	Float32 *float32
	Float64 *float64
	t       int
}

func AsInt() *T {
	return newT(0)
}

func AsInt8() *T {
	return newT(1)
}

func AsInt16() *T {
	return newT(2)
}

func AsInt32() *T {
	return newT(3)
}

func AsInt64() *T {
	return newT(4)
}

func AsString() *T {
	return newT(5)
}

func AsChar() *T {
	return newT(6)
}

func AsFloat32() *T {
	return newT(7)
}

func AsFloat64() *T {
	return newT(8)
}

func newT(choice int) *T {

	var t *T = new(T)
	t.t = choice

	t.Int = nil
	t.Int8 = nil
	t.Int16 = nil
	t.Int32 = nil
	t.Int64 = nil
	t.Char = nil
	t.Float32 = nil
	t.Float64 = nil
	t.String = nil

	switch choice {
	case 0:
		t.Int = new(int)
	case 1:
		t.Int8 = new(int8)
	case 2:
		t.Int16 = new(int16)
	case 3:
		t.Int32 = new(int32)
	case 4:
		t.Int64 = new(int64)
	case 5:
		t.String = new(string)
	case 6:
		t.Char = new(byte)
	case 7:
		t.Float32 = new(float32)
	case 8:
		t.Float64 = new(float64)
	}

	return t
}
