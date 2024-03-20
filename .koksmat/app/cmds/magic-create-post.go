// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Create Rooms
---
*/
package cmds

import (
	"context"

	"github.com/365admin/magicbox-mongodb/execution"
	"github.com/365admin/magicbox-mongodb/utils"
)

func MagicCreatePost(ctx context.Context, args []string) (*string, error) {

	result, pwsherr := execution.ExecutePowerShell("john", "*", "magicbox-mongodb", "00-magic", "10-create.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))
	return nil, nil

}