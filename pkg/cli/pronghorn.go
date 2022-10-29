package cli

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tietang/props/ini"
	"pronghorn/pkg/boot"
)

func NewCommand() *cobra.Command {
	cfgFile := ""
	cmd := &cobra.Command{
		Use:   "Pronghorn",
		Short: "Pronghorn is a proxy server",
		RunE: func(*cobra.Command, []string) error {
			//获取程序运行文件所在路径

			// cfgFile could be defined by user or default `.rr.yaml`
			// this check added just to be safe
			if cfgFile == "" {
				log.Fatalln("no configuration file provided")
			}

			//加载配置文件
			conf := ini.NewIniFileCompositeConfigSource(cfgFile)

			boot := boot.NewBootApplication(conf)

			boot.Start()

			return nil
		},
	}
	f := cmd.PersistentFlags()
	f.StringVarP(&cfgFile, "config", "c", "", "config file")
	return cmd
}
