package command

import "lazyliner/domain/command"

type ChooseCommandDTO struct {
	Commands []chooseCommandDTO
}

type chooseCommandDTO struct {
	Command     string
	Description string
	Tags        []string
}

type ChooseCommandUseCase interface {
	ChooseCommandDTO() *ChooseCommandDTO
}

type chooseCommandUseCase struct {
	CommandRepository command.Repository
}

func NewChooseCommandUseCase(r command.Repository) ChooseCommandUseCase {
	return &chooseCommandUseCase{r}
}

func (ccu *chooseCommandUseCase) ChooseCommandDTO() *ChooseCommandDTO {
	var internalDTOs []chooseCommandDTO
	cmds := ccu.CommandRepository.Fetch()
	for _, c := range cmds {
		internalDTO := chooseCommandDTO{
			Command:     c.OnelineCommand(),
			Description: c.Description(),
			Tags:        c.Tags(),
		}
		internalDTOs = append(internalDTOs, internalDTO)
	}
	return &ChooseCommandDTO{
		Commands: internalDTOs,
	}
}
