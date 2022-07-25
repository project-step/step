package planner

import (
	"strings"

	"github.com/stepneko/neko-session/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var rules []string
var tabs int = 0

type MySqlListener struct {
	*parser.BaseMySqlParserListener
}

func NewMySqlListener() *MySqlListener {
	return &MySqlListener{}
}

func (l *MySqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	for i := 0; i < tabs; i++ {
		print(" ")
	}
	ctx.GetStart()
	print(rules[ctx.GetRuleIndex()])
	print("  ")
	println(tabs)
	for i := 0; i < tabs; i++ {
		print(" ")
	}
	println(ctx.GetText())
	for i := 0; i < tabs; i++ {
		print(" ")
	}
	println("--------------------------------------------")
	tabs += 2
	return
}

func (l *MySqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	tabs -= 2
}

func (l *MySqlListener) VisitTerminal(node antlr.TerminalNode) {
	for i := 0; i < tabs; i++ {
		print(" ")
	}
	print("terminal node")
	print("  ")
	println(tabs)
	for i := 0; i < tabs; i++ {
		print(" ")
	}
	println("--------------------------------------------")
}

func BuildTree(query string) error {
	stream := antlr.NewInputStream(strings.ToUpper(query))
	lexer := parser.NewMySqlLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	sqlParser := parser.NewMySqlParser(tokenStream)
	rules = sqlParser.RuleNames
	l := NewMySqlListener()

	tree := sqlParser.SqlStatements()
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
	t := tree.ToStringTree(sqlParser.RuleNames, sqlParser)
	println(t)

	return nil
}
