package command

type Command struct {
	onelineCommand string
	description    string
	tags           []string
}

func (c *Command) OnelineCommand() string {
	return c.onelineCommand
}

func (c *Command) Description() string {
	return c.description
}

func (c *Command) Tags() []string {
	return c.tags
}
