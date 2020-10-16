package main

import (
	"os"
	"path/filepath"

	"ds/cmd/ds/cmd"
	"ds/pkg/util"
)

var targetDir = ".gotmp/bin"

func main() {
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		err = os.MkdirAll(targetDir, 0755)
		util.CheckErr(err)
	}
	err := cmd.RootCmd.GenBashCompletionFile(filepath.Join(targetDir, "ds_bash_autocomplete.sh"))
	if err != nil {
		util.Failed("could not generate ds_bash_autocomplete.sh: %v", err)
	}
	err = cmd.RootCmd.GenZshCompletionFile(filepath.Join(targetDir, "ds_zsh_autocomplete.sh"))
	if err != nil {
		util.Failed("could not generate ds_zsh_autocomplete.sh: %v", err)
	}
	err = cmd.RootCmd.GenFishCompletionFile(filepath.Join(targetDir, "ds_fish_autocomplete.sh"), true)
	if err != nil {
		util.Failed("could not generate ds_fish_autocomplete.sh: %v", err)
	}
	err = cmd.RootCmd.GenPowerShellCompletionFile(filepath.Join(targetDir, "ds_powershell_autocomplete.ps1"))
	if err != nil {
		util.Failed("could not generate ds_powershell_autocomplete.ps1: %v", err)
	}
}
