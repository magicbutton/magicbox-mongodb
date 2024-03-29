// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Forward To Database
---
*/
package cmds

import (
	"context"

	"github.com/365admin/magicbox-mongodb/execution"
	"github.com/365admin/magicbox-mongodb/utils"
)

func ConnectForwardPost(ctx context.Context, args []string) (*string, error) {

	result, pwsherr := execution.ExecutePowerShell("john", "*", "magicbox-mongodb", "07-connect", "20-forward.ps1", "", "-databasename", args[0])
	if pwsherr != nil {
		return nil, pwsherr
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))
	return nil, nil

}
