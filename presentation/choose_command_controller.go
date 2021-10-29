package presentation

import (
	"fmt"
	ic "lazyliner/infra/command"
	uc "lazyliner/usecase/command"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/ktr0731/go-fuzzyfinder"
)

func ChooseCommandController() error {
	repo, err := ic.NewTomlRepository("./config.toml")
	if err != nil {
		return err
	}
	cmds := uc.NewChooseCommandUseCase(repo).ChooseCommandDTO().Commands
	idx, err := fuzzyfinder.FindMulti(
		cmds,
		func(i int) string {
			return fmt.Sprintf("%s,%s,%s", cmds[i].Command, cmds[i].Description, strings.Join(cmds[i].Tags, ","))
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Command %s\nDesc %s\nTags %s", cmds[i].Command, cmds[i].Description, strings.Join(cmds[i].Tags, ","))
		}),
	)
	if err != nil {
		return err
	}
	clipboard.WriteAll(cmds[idx[0]].Command)
	return nil
}
