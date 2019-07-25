package console

import "gopkg.in/alecthomas/kingpin.v2"

func ConfigureCommands(app *kingpin.Application)  {
	cmd_addUser := &AddUserCommand{}
	addUser := app.Command("user-add", "create users.").Action(cmd_addUser.Run)
	// 配置命令入参flag
	addUser.Flag("data", "user data.").Short('d').StringVar(&cmd_addUser.Data)
	_ = addUser
}