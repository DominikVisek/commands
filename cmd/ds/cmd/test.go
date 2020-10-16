package cmd

import (
	"github.com/spf13/cobra"

	"ds/pkg/output"
	"ds/pkg/util"
)

var selfUpdateCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  `test`,
	Run: func(cmd *cobra.Command, args []string) {
		util.Info("Dowloading new wesion of ds!")
		fileUrl := "https://static.javatpoint.com/go/images/go-tutorial.jpg"
		err := util.DownloadFile("logo.jpg", fileUrl, true)

		if err != nil {
			output.UserOut.Panic(err)
		}
		util.Success("Downloaded: " + fileUrl)
	},
}

func init() {
	RootCmd.AddCommand(selfUpdateCmd)
}
