package planner

import (
	"fmt"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/parser/test_driver"
)

func parse(sql string) (*ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return &stmtNodes[0], nil
}

func BuildTree(query string) error {
	astNode, err := parse("SELECT a, b FROM t")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", *astNode)
	return nil
}
