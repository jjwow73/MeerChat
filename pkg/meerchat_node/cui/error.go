package meerchat_node

import "fmt"

type CommandError struct {
	OriginalError error
}

func (commandErr *CommandError) Error() string {
	return fmt.Sprintf("command error : %v", commandErr.OriginalError)
}