// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: Database Restore
---
*/
package endpoints

import (
	"context"
	"encoding/json"
	"os"
	"path"

	"github.com/swaggest/usecase"

	"github.com/365admin/magicbox-mongodb/execution"
	"github.com/365admin/magicbox-mongodb/schemas"
	"github.com/365admin/magicbox-mongodb/utils"
)

func RestoreListtarPost() usecase.Interactor {
	type Request struct {
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *schemas.Backupcontent) error {

		_, err := execution.ExecutePowerShell("john", "*", "magicbox-mongodb", "30-restore", "57-listtarfile.ps1", "")
		if err != nil {
			return err
		}

		resultingFile := path.Join(utils.WorkDir("magicbox-mongodb"), "backupcontent.json")
		data, err := os.ReadFile(resultingFile)
		if err != nil {
			return err
		}
		resultObject := schemas.Backupcontent{}
		err = json.Unmarshal(data, &resultObject)
		*output = resultObject
		return err

	})
	u.SetTitle("Database Restore")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Backup")
	return u
}
