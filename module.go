package main

import (
	"github.com/go-ini/ini"
	"github.com/spf13/cobra"
	"log"
)

var boolMapping = map[bool]string{
	true:  "yes",
	false: "no",
}

var cmdAddModule, cmdDelModule *cobra.Command

var moduleComment, moduleUid, moduleGid, moduleAuthUser, moduleSecrets string
var moduleReadOnly, mouduleListable bool

func addModule(cmd *cobra.Command, args []string) {

	moduleName := args[0]
	modulePath := args[1]

	cfg, err := ini.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	section, err := cfg.NewSection(moduleName)

	if err != nil {
		log.Fatal(err)
	}
	_, err = section.NewKey("path", modulePath)
	_, err = section.NewKey("comment", moduleComment)
	_, err = section.NewKey("read only", boolMapping[moduleReadOnly])
	_, err = section.NewKey("list", boolMapping[mouduleListable])
	_, err = section.NewKey("uid", moduleUid)
	_, err = section.NewKey("gid", moduleGid)
	_, err = section.NewKey("auth users", moduleAuthUser)
	_, err = section.NewKey("secrets files", moduleSecrets)

	cfg.SaveTo(configPath)
}

func delModule(cmd *cobra.Command, args []string) {

	moduleName := args[0]

	cfg, err := ini.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg.DeleteSection(moduleName)
	cfg.SaveTo(configPath)
}

func init() {
	cmdAddModule = &cobra.Command{
		Use:   "add-module [module name] [absolute path]",
		Short: "Add modules into rsync config",
		Long:  "Add modules into rsync config",
		Run:   addModule,
	}

	cmdAddModule.Flags().StringVarP(&moduleComment, "comment", "C", "Unknown Module", "Comment for a module")
	cmdAddModule.Flags().BoolVarP(&moduleReadOnly, "readonly", "r", false, "Is a module read only")
	cmdAddModule.Flags().BoolVarP(&mouduleListable, "listable", "L", true, "Is a module listable")
	cmdAddModule.Flags().StringVarP(&moduleUid, "module-uid", "U", "user", "Uid for chroot")
	cmdAddModule.Flags().StringVarP(&moduleGid, "module-gid", "G", "user", "Gid for chroot")
	cmdAddModule.Flags().StringVarP(&moduleAuthUser, "auth-user", "A", "user", "User for authentication")
	cmdAddModule.Flags().StringVarP(&moduleSecrets, "secret-file", "S", "/etc/rsyncd.secrets", "Secret file for this module")

	cmdDelModule = &cobra.Command{
		Use:   "del-module [module name] [absolute path]",
		Short: "Remove modules from rsync config",
		Long:  "Remove modules from rsync config",
		Run:   delModule,
	}
}
