// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Download all backups
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

func RestoreDownloadPost(ctx context.Context, body []byte, args []string) (*string, error) {
	inputErr := os.WriteFile(path.Join(utils.WorkDir("magicbox-mongodb"), "bloblist.json"), body, 0644)
	if inputErr != nil {
		return nil, inputErr
	}

	_, pwsherr := execution.ExecutePowerShell("john", "*", "magicbox-mongodb", "30-restore", "20-download-all.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	return nil, nil

}