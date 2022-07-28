package state

type DataHandleStatus int

const (
	DataHandleStatus_NotReady DataHandleStatus = iota
	DataHandleStatus_Loading
	DataHandleStatus_Ready
	DataHandleStatus_Error
)

var QueryMap = make(map[string]TableDataHandle)

type TableDataHandle interface {
	Load(tableName string, columns []int) error
	Cols() int
	Rows() int
	GetData(rowInd int, colInd int) []byte
	SetStatus(status DataHandleStatus)
	GetStatus() DataHandleStatus
}

type SimpleTableDataHandle struct {
	Status  DataHandleStatus
	DataSet [][][]byte
}

func NewSimpleTableDataHandle() *SimpleTableDataHandle {
	return &SimpleTableDataHandle{
		Status:  DataHandleStatus_NotReady,
		DataSet: [][][]byte{},
	}
}

func (tds *SimpleTableDataHandle) Load(tableName string, columns []int) error {
	dataSource := HashMap[tableName]
	rowCount := len(dataSource[0])
	colCount := len(columns)

	for colInd := 0; colInd < colCount; colInd++ {
		tds.DataSet = append(tds.DataSet, [][]byte{})
		realColInd := columns[colInd]
		for rowInd := 0; rowInd < rowCount; rowInd++ {
			tds.DataSet[colInd] = append(tds.DataSet[colInd], dataSource[realColInd][rowInd])
		}
	}

	return nil
}

func (tds *SimpleTableDataHandle) Cols() int {
	return len(tds.DataSet)
}

func (tds *SimpleTableDataHandle) Rows() int {
	return len(tds.DataSet[0])
}

func (tds *SimpleTableDataHandle) GetData(rowId int, colId int) []byte {
	return tds.DataSet[colId][rowId]
}

func (tds *SimpleTableDataHandle) SetStatus(status DataHandleStatus) {
	tds.Status = status
}

func (tds *SimpleTableDataHandle) GetStatus() DataHandleStatus {
	return tds.Status
}
