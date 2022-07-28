package state

import "strconv"

type ColumnType int

const (
	ColumnType_Int ColumnType = iota
	ColumnType_String
)

const RowCount int = 1000000

type ColumnInfo struct {
	Index int
	Typ   ColumnType
}

// Hacky mock for a table in memory
var HashMap = make(map[string][][][]byte)
var SchemaMap = make(map[string]map[string]*ColumnInfo)

func Init() {
	println("warming up, loading in data")

	tableUserName := "users"
	SchemaMap[tableUserName] = make(map[string]*ColumnInfo)
	SchemaMap[tableUserName]["userid"] = &ColumnInfo{
		Index: 0,
		Typ:   ColumnType_Int,
	}
	SchemaMap[tableUserName]["name"] = &ColumnInfo{
		Index: 1,
		Typ:   ColumnType_String,
	}
	SchemaMap[tableUserName]["gender"] = &ColumnInfo{
		Index: 2,
		Typ:   ColumnType_String,
	}

	HashMap[tableUserName] = [][][]byte{}
	// Init columns in HashMap
	for i := 0; i < 3; i++ {
		HashMap[tableUserName] = append(HashMap[tableUserName], [][]byte{})
	}

	for i := 0; i < RowCount; i++ {
		HashMap[tableUserName][0] = append(HashMap[tableUserName][0], []byte(strconv.Itoa(i)))
		HashMap[tableUserName][1] = append(HashMap[tableUserName][1], []byte("zhanglaosan"))
		HashMap[tableUserName][2] = append(HashMap[tableUserName][2], []byte("male"))
	}

	println(len(HashMap[tableUserName][0]))
	println(len(HashMap[tableUserName][1]))
	println(len(HashMap[tableUserName][2]))

	println("warm up finished")
}
