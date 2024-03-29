// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Backup all databases
---
*/
package cmds

import (
	"context"
	"os"
	"path"

	"github.com/365admin/magicbox-mongodb/execution"
	"github.com/365admin/magicbox-mongodb/utils"
)

func BackupAllPost(ctx context.Context, body []byte, args []string) (*string, error) {
	inputErr := os.WriteFile(path.Join(utils.WorkDir("magicbox-mongodb"), "databaseservices.json"), body, 0644)
	if inputErr != nil {
		return nil, inputErr
	}

	result, pwsherr := execution.ExecutePowerShell("john", "*", "magicbox-mongodb", "10-backup", "20 backup-all.ps1", "", "-upload", args[0])
	if pwsherr != nil {
		return nil, pwsherr
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))
	return nil, nil

}
