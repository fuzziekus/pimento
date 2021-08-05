package ui

var ColumnOrder = []string{
	"Description",
	"UserId",
	"Password",
	"Memo",
}

var FlagMapping = map[string]int{
	"Description": 1,
	"UserId":      2,
	"Password":    4,
	"Memo":        8,
	"NoPass":      11,
	"All":         15,
}

type Flags struct {
	All         bool
	Description bool
	UserId      bool
	Password    bool
	Memo        bool
	NoPass      bool
	flagVar     int
	NoHeader    bool
}
