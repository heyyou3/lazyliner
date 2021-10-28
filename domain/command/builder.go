package command

type implCommand struct {
	Command
}

type Builder interface {
	SetOnelineCommand(string) Builder
	SetDescription(string) Builder
	SetTags([]string) Builder
	NewCommand() Command
}

func Build() Builder {
	return &implCommand{}
}

func (c *implCommand) SetOnelineCommand(olc string) Builder {
	c.Command.onelineCommand = olc
	return c
}

func (c *implCommand) SetDescription(d string) Builder {
	c.Command.description = d
	return c
}

func (c *implCommand) SetTags(t []string) Builder {
	c.Command.tags = t
	return c
}

func (c *implCommand) NewCommand() Command {
	return c.Command
}
