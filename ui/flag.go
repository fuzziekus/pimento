package ui

var ColumnOrder = []string{
	"ItemName",
	"UserName",
	"Password",
	"Tag",
}

var FlagMapping = map[string]int{
	"ItemName": 1,
	"UserName": 2,
	"Password": 4,
	"Tag":      8,
	"NoPass":   11,
	"All":      15,
}

type Flags struct {
	All      bool
	ItemName bool
	UserName bool
	Password bool
	Tag      bool
	NoPass   bool
	flagVar  int
	NoHeader bool
}
