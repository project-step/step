package cli

import (
	"github.com/spf13/cobra"

	"github.com/stepneko/step/session"
)

var sessionManagerCfg = session.SessionManagerConfig{}

func Initialize(cmd *cobra.Command) {
	flags := cmd.Flags()
	flags.SortFlags = false

	// Step server configs
	flags.IntVarP(&sessionManagerCfg.Port, "port", "P", 3309, "Listen port of step server")
	flags.StringVar(&sessionManagerCfg.User, "user", "root", "User name of step server")
	flags.StringVar(&sessionManagerCfg.Pass, "pass", "", "Password of step server")

	// MySQL server configs
	flags.StringVar(&sessionManagerCfg.MySQL.Host, "mysql.host", "127.0.0.1", "MySQL server host name")
	flags.IntVar(&sessionManagerCfg.MySQL.Port, "mysql.port", 3306, "MySQL server port")
	flags.StringVar(&sessionManagerCfg.MySQL.User, "mysql.user", "root", "MySQL server user name")
	flags.StringVar(&sessionManagerCfg.MySQL.Pass, "mysql.pass", "", "MySQL password")
	flags.StringVar(&sessionManagerCfg.MySQL.Db, "mysql.db", "", "MySQL database name")
	flags.StringVar(&sessionManagerCfg.MySQL.Options, "mysql.options", "charset=utf8mb4", "MySQL server connection options")
}

func CreateSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "session",
		Short:   "start a node as Step session manager",
		Example: `  step session --port 3309 --mysql.user root --mysql.pass 123456`,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return session.Start(cmd, &sessionManagerCfg)
		},
	}
	Initialize(cmd)
	return cmd
}
