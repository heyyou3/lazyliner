package command

import (
	"oneline-lazy/domain/command"

	"github.com/BurntSushi/toml"
)

type TomlConfigs struct {
	Commands []TomlConfig `toml:"commands"`
}

type TomlConfig struct {
	Command     string   `toml:"command"`
	Description string   `toml:"description"`
	Tags        []string `toml:"tags"`
}

type TomlRepository struct {
	Config *TomlConfigs
}

func NewTomlRepository(filePath string) (*TomlRepository, error) {
	var config TomlConfigs
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		return nil, err
	}
	return &TomlRepository{Config: &config}, nil
}

var _ command.Repository = &TomlRepository{}

func (tr *TomlRepository) Fetch() []command.Command {
	var commands []command.Command
	for _, cmd := range tr.Config.Commands {
		cmdBuiler := command.Build().
			SetOnelineCommand(cmd.Command).
			SetDescription(cmd.Description).
			SetTags(cmd.Tags)
		commands = append(commands, cmdBuiler.NewCommand())
	}
	return commands
}
