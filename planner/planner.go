package planner

import (
	"sync"

	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/parser/test_driver"
	"github.com/stepneko/neko-dataflow/edge"
	"github.com/stepneko/neko-dataflow/iterator"
	"github.com/stepneko/neko-dataflow/operators"
	"github.com/stepneko/neko-dataflow/request"
	"github.com/stepneko/neko-dataflow/scope"
	"github.com/stepneko/neko-dataflow/step"
	"github.com/stepneko/neko-dataflow/timestamp"
	"github.com/stepneko/neko-dataflow/worker"
	"github.com/stepneko/neko-session/state"
)

func result(logicalPlan *LogicalPlan, tableState *state.TableDataSet) (*mysql.Result, error) {
	res := &mysql.Result{
		Resultset: mysql.NewResultset(len(tableState.DataSet)),
	}
	// HACK just print data from server side
	for _, n := range logicalPlan.ColumnNames {
		print(n)
		print("  ")
	}
	println()

	for rowInd := 0; rowInd < len(tableState.DataSet[0]); rowInd++ {
		for colInd := 0; colInd < len(tableState.DataSet); colInd++ {
			print(string(tableState.DataSet[colInd][rowInd]))
			print("  ")
		}
		println()
	}
	// print actual data
	// TODO add data from tableState into the res to return to client conn
	return res, nil
}

func execPlan(logicalPlan *LogicalPlan) (*mysql.Result, error) {
	var wg sync.WaitGroup
	q := logicalPlan.Query
	tableState, exist := state.QueryMap[q]
	if !exist {
		wg.Add(1)
		tableState = state.NewTableDataSet()
		ch := make(chan request.InputDatum)

		go step.Start(func(w worker.Worker) error {
			state.QueryMap[logicalPlan.Query] = tableState
			w.Dataflow(func(s scope.Scope) error {
				operators.
					NewInput(s, ch).
					Inspect(func(e edge.Edge, msg *request.Message, ts timestamp.Timestamp) (iterator.Iterator[*request.Message], error) {
						if msg.ToString() == "init" {
							tableState.Status = state.DataSetStatus_Loading
							if err := tableState.LoadData(logicalPlan.TableName, logicalPlan.Columns); err != nil {
								tableState.Status = state.DataSetStatus_Error
								return nil, err
							}
							tableState.Status = state.DataSetStatus_Ready
							wg.Done()
						}
						return nil, nil
					})
				return nil
			})
			return nil
		})
		ch <- request.NewInputRaw(request.NewMessage([]byte("init")), *timestamp.NewTimestamp())
	}
	wg.Wait()
	return result(logicalPlan, tableState)
}

func parse(sql string) (*ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}
	return &stmtNodes[0], nil
}

func PlanQeury(query string) (*mysql.Result, error) {
	astNode, err := parse(query)
	if err != nil {
		return nil, err
	}
	logicalPlan, err := extract(astNode)
	if err != nil {
		return nil, err
	}
	logicalPlan.Query = query
	return execPlan(logicalPlan)
}
