package cmd

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/UKHomeOffice/kmm/pkg/kubeadm"
)

// versionCmd represents the version command
var AddonsCmd = &cobra.Command{
	Use:   "addons",
	Short: "Will deploy cluster resources",
	Long:  "Will deploy / redeploy essential cluster resources",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := getKmmConfig(cmd)
		if err == nil {
			err = kubeadm.Addons(cfg.KubeadmCfg)
		}
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
	},
}

func init() {
	RootCmd.AddCommand(AddonsCmd)
}