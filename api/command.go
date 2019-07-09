package api

type Command struct {
  Cmd string
  Args []interface{}
}

type CommandApi interface {
  Start() error
  Chan() (chan Command, error)
}

func newCommand(cmd string, args []interface{}) Command {
  return Command{
    Cmd: cmd,
    Args: args,
  }
}
