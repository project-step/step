package state

type DataSetStatus int

const (
	DataSetStatus_NotReady DataSetStatus = iota
	DataSetStatus_Loading
	DataSetStatus_Ready
	DataSetStatus_Error
)

type TableDataSet struct {
	Status  DataSetStatus
	DataSet [][][]byte
}

func NewTableDataSet() *TableDataSet {
	return &TableDataSet{
		Status:  DataSetStatus_NotReady,
		DataSet: [][][]byte{},
	}
}

func (tds *TableDataSet) LoadData(tableName string, columns []int) error {
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

var QueryMap = make(map[string]*TableDataSet)
