package smpp34

import "strconv"

type Field interface {
	Length() interface{}
	Value() interface{}
	String() string
	ByteArray() []byte
}

type VariableField struct {
	value []byte
}

type FixedField struct {
	size  uint8
	value uint8
}

func NewVariableField(v []byte) Field {
	i := &VariableField{v}
	f := Field(i)
	return f
}

func NewFixedField(v uint8) Field {
	i := &FixedField{1, v}
	f := Field(i)
	return f
}

func (v *VariableField) Length() interface{} {
	l := len(v.value)
	return l
}

func (v *VariableField) Value() interface{} {
	return v.value
}

func (v *VariableField) String() string {
	return string(v.value)
}

func (v *VariableField) ByteArray() []byte {
	return v.value
}

func (f *FixedField) Length() interface{} {
	return uint8(1)
}

func (f *FixedField) Value() interface{} {
	return f.value
}

func (f *FixedField) String() string {
	return strconv.Itoa(int(f.value))
}

func (f *FixedField) ByteArray() []byte {
	return packUi8(f.value)
}
