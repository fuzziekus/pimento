package ui

import "reflect"

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
	"NoPass":   15,
	"All":      15,
}

type Flags struct {
	All         bool
	ItemName    bool
	UserName    bool
	Password    bool
	Tag         bool
	NoPass      bool
	flagVar     int
	NoHeader    bool
	FormatCSV   bool
	FormatTable bool
}

func (e *Flags) calcCondition() {
	e.flagVar = 0

	// カラム指定がある場合は、All フラグは false にし、
	// 対象カラムだけ更新するようにする
	if e.All {
		e.NoPass = false
	}

	if e.ItemName || e.UserName || e.Password || e.Tag {
		e.All = false
		e.NoPass = false
	}

	

	// CSV が指定されていれば、CSV出力を優先
	if e.FormatCSV {
		e.FormatTable = false
	}

	// フィールド名を動的に取得して e.flagVar を更新するため
	// reflect を利用
	rtEditFlag := reflect.TypeOf(Flags{})
	rvEditFlag := reflect.ValueOf(e).Elem()
	for i := 0; i < rtEditFlag.NumField(); i++ {
		f := rtEditFlag.Field(i)
		for key, val := range FlagMapping {
			v := rvEditFlag.FieldByName(key).Interface()
			if f.Name == key && v.(bool) {
				e.flagVar |= val
			}
		}
	}
}
