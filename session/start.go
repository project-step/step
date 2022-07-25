package session

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"

	"github.com/go-mysql-org/go-mysql/server"
)

func onConnect(c net.Conn, cfg *SessionManagerConfig) error {
	connIdent := c.RemoteAddr().String()
	h, err := NewHandler(cfg, connIdent)
	if err != nil {
		// return err
	}
	conn, err := server.NewConn(c, cfg.User, cfg.Pass, h)
	if err != nil {
		return err
	}

	defer func() {
		if err := h.Finalize(); err != nil {
			fmt.Println("Finalize handler failed", err)
		}
	}()

	for {
		if err := conn.HandleCommand(); err != nil {
			if conn.Conn == nil {
				return nil
			}
			fmt.Printf("handle command error: %v\n", err)
			return err
		}
	}
}

func Start(cmd *cobra.Command, sessionManagerCfg *SessionManagerConfig) error {
	l, err := net.Listen("tcp4", fmt.Sprintf(":%d", sessionManagerCfg.Port))
	if err != nil {
		return err
	}

	println("Listening to port: ", sessionManagerCfg.Port)

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}
		go onConnect(c, sessionManagerCfg)
	}
}
