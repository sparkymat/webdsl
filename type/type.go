package types

type Type int32

const (
	Int8       Type = iota
	Int16      Type = iota
	Int32      Type = iota
	Int64      Type = iota
	Uint8      Type = iota
	Uint16     Type = iota
	Uint32     Type = iota
	Uint64     Type = iota
	Float32    Type = iota
	Float64    Type = iota
	Complex64  Type = iota
	Complex128 Type = iota
	Byte       Type = iota
	Rune       Type = iota
	String     Type = iota
)
