package session

import (
	"fmt"

	"github.com/go-mysql-org/go-mysql/client"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/server"
	"github.com/stepneko/neko-session/planner"

	"github.com/hashicorp/go-multierror"
)

type ConnHandler struct {
	server.EmptyHandler

	cfg       *SessionManagerConfig
	mysqlConn *client.Conn
	connIdent string
}

func NewHandler(config *SessionManagerConfig, connIdent string) (*ConnHandler, error) {
	h := ConnHandler{
		cfg:       config,
		connIdent: connIdent,
	}
	if err := h.initialize(); err != nil {
		return nil, err
	}
	return &h, nil

}

func (h *ConnHandler) initialize() error {
	mysqlCfg := h.cfg.MySQL
	mysqlConn, err := client.Connect(
		fmt.Sprintf("%s:%d", mysqlCfg.Host, mysqlCfg.Port),
		mysqlCfg.User,
		mysqlCfg.Pass,
		mysqlCfg.Db,
	)
	if err != nil {
		return err
	}
	if err := mysqlConn.Ping(); err != nil {
		return err
	}
	h.mysqlConn = mysqlConn

	return nil
}

func (h *ConnHandler) handleSelect(query string) (*mysql.Result, error) {
	return planner.PlanQeury(query)
}

func (h *ConnHandler) Finalize() error {
	var result error
	if h.mysqlConn != nil {
		err := h.mysqlConn.Close()
		if err != nil {
			result = multierror.Append(result, err)
		}
	} else {
	}
	return result
}

func (h *ConnHandler) UseDB(db string) error {
	var result error
	if err := h.mysqlConn.UseDB(db); err != nil {
		result = multierror.Append(result, err)
	}
	return result
}

func (h *ConnHandler) HandleQuery(query string) (*mysql.Result, error) {
	switch parseCmdType(query) {
	case CmdType_Select:
		return h.handleSelect(query)
	default:
		return h.mysqlConn.Execute(query)
	}
}

func (h *ConnHandler) HandleFieldList(table string, fieldWildcard string) ([]*mysql.Field, error) {
	fields, err1 := h.mysqlConn.FieldList(table, fieldWildcard)

	return fields, err1
}

func (h *ConnHandler) HandleStmtPrepare(query string) (int, int, interface{}, error) {
	stmt, err := h.mysqlConn.Prepare(query)

	return stmt.ParamNum(), stmt.ColumnNum(), nil, err
}

func (h *ConnHandler) HandleStmtExecute(context interface{}, query string, args []interface{}) (*mysql.Result, error) {
	result, err1 := h.mysqlConn.Execute(query, args...)

	return result, err1
}
