package magicapp

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/365admin/magicbox-mongodb/cmds"
	"github.com/365admin/magicbox-mongodb/utils"
)

func RegisterCmds() {
	RootCmd.PersistentFlags().StringVarP(&utils.Output, "output", "o", "", "Output format (json, yaml, xml, etc.)")

	healthCmd := &cobra.Command{
		Use:   "health",
		Short: "Health",
		Long:  `Everything you need to be able to do to deploy, use and operate MongoDB on the Magicbox platform`,
	}
	HealthPingPostCmd := &cobra.Command{
		Use:   "ping  pong",
		Short: "Ping",
		Long:  `Simple ping endpoint`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.HealthPingPost(ctx, args)
		},
	}
	healthCmd.AddCommand(HealthPingPostCmd)
	HealthCoreversionPostCmd := &cobra.Command{
		Use:   "coreversion ",
		Short: "Core Version",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.HealthCoreversionPost(ctx, args)
		},
	}
	healthCmd.AddCommand(HealthCoreversionPostCmd)

	RootCmd.AddCommand(healthCmd)
	magicCmd := &cobra.Command{
		Use:   "magic",
		Short: "Magic Buttons",
		Long:  `Everything you need to be able to do to deploy, use and operate MongoDB on the Magicbox platform`,
	}
	MagicBackupPostCmd := &cobra.Command{
		Use:   "backup ",
		Short: "Backup All",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.MagicBackupPost(ctx, args)
		},
	}
	magicCmd.AddCommand(MagicBackupPostCmd)

	RootCmd.AddCommand(magicCmd)
	discoverCmd := &cobra.Command{
		Use:   "discover",
		Short: "Discovery",
		Long:  `Everything you need to be able to do to deploy, use and operate MongoDB on the Magicbox platform`,
	}
	DiscoverDatabasesPostCmd := &cobra.Command{
		Use:   "databases ",
		Short: "Database Discovery",
		Long:  `Discover databases in the cluster`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.DiscoverDatabasesPost(ctx, args)
		},
	}
	discoverCmd.AddCommand(DiscoverDatabasesPostCmd)

	RootCmd.AddCommand(discoverCmd)
	connectCmd := &cobra.Command{
		Use:   "connect",
		Short: "Connect",
		Long:  `Everything you need to be able to do to deploy, use and operate MongoDB on the Magicbox platform`,
	}
	ConnectConnectionstringsPostCmd := &cobra.Command{
		Use:   "connectionstrings  databasename",
		Short: "Get connection strings",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ConnectConnectionstringsPost(ctx, args)
		},
	}
	connectCmd.AddCommand(ConnectConnectionstringsPostCmd)
	ConnectForwardPostCmd := &cobra.Command{
		Use:   "forward  databasename",
		Short: "Forward To Database",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ConnectForwardPost(ctx, args)
		},
	}
	connectCmd.AddCommand(ConnectForwardPostCmd)

	RootCmd.AddCommand(connectCmd)
	backupCmd := &cobra.Command{
		Use:   "backup",
		Short: "Backup",
		Long:  `Everything you need to be able to do to deploy, use and operate MongoDB on the Magicbox platform`,
	}
	BackupAllPostCmd := &cobra.Command{
		Use:   "all  upload",
		Short: "Backup all databases",
		Long:  `Backup all databases in the cluster`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			body, err := os.ReadFile(args[0])
			if err != nil {
				panic(err)
			}

			cmds.BackupAllPost(ctx, body, args)
		},
	}
	backupCmd.AddCommand(BackupAllPostCmd)

	RootCmd.AddCommand(backupCmd)
	restoreCmd := &cobra.Command{
		Use:   "restore",
		Short: "Restore",
		Long:  `Everything you need to be able to do to deploy, use and operate MongoDB on the Magicbox platform`,
	}
	RestoreListPostCmd := &cobra.Command{
		Use:   "list ",
		Short: "List backup blobs",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.RestoreListPost(ctx, args)
		},
	}
	restoreCmd.AddCommand(RestoreListPostCmd)
	RestoreDownloadPostCmd := &cobra.Command{
		Use:   "download ",
		Short: "Download all backups",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			body, err := os.ReadFile(args[0])
			if err != nil {
				panic(err)
			}

			cmds.RestoreDownloadPost(ctx, body, args)
		},
	}
	restoreCmd.AddCommand(RestoreDownloadPostCmd)
	RestoreUnarchivePostCmd := &cobra.Command{
		Use:   "unarchive  database",
		Short: "Database Restore",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.RestoreUnarchivePost(ctx, args)
		},
	}
	restoreCmd.AddCommand(RestoreUnarchivePostCmd)
	RestoreListtarPostCmd := &cobra.Command{
		Use:   "listtar  database",
		Short: "Database Restore",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.RestoreListtarPost(ctx, args)
		},
	}
	restoreCmd.AddCommand(RestoreListtarPostCmd)

	RootCmd.AddCommand(restoreCmd)
	provisionCmd := &cobra.Command{
		Use:   "provision",
		Short: "Provision",
		Long:  `Everything you need to be able to do to deploy, use and operate MongoDB on the Magicbox platform`,
	}
	ProvisionAppdeployproductionPostCmd := &cobra.Command{
		Use:   "appdeployproduction ",
		Short: "App deploy to production",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ProvisionAppdeployproductionPost(ctx, args)
		},
	}
	provisionCmd.AddCommand(ProvisionAppdeployproductionPostCmd)

	RootCmd.AddCommand(provisionCmd)
}
