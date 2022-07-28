package planner

import (
	"fmt"

	"github.com/pingcap/tidb/parser/ast"
	"github.com/stepneko/neko-session/state"
)

type colX struct {
	tableName  string
	colNames   []string
	colIndices []int
}

func (v *colX) Enter(in ast.Node) (ast.Node, bool) {
	if name, ok := in.(*ast.ColumnName); ok {
		v.colNames = append(v.colNames, name.Name.O)
	}
	if name, ok := in.(*ast.TableName); ok {
		v.tableName = name.Name.O
		println(name.Name.O)
	}
	return in, false
}

func (v *colX) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func extract(rootNode *ast.StmtNode) (*LogicalPlan, error) {
	v := &colX{}
	(*rootNode).Accept(v)

	// Convert column name to column indicies
	for _, colName := range v.colNames {
		info, exist := state.SchemaMap[v.tableName][colName]
		if !exist {
			return nil, fmt.Errorf("unknown column name: %s", colName)
		}
		v.colIndices = append(v.colIndices, info.Index)
	}

	return &LogicalPlan{
		TableName:   v.tableName,
		Columns:     v.colIndices,
		ColumnNames: v.colNames,
	}, nil
}
