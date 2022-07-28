package planner

type LogicalPlan struct {
	Query       string
	TableName   string
	ColumnNames []string
	Columns     []int
}
