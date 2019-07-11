package api

/*
Command contains a parsed incoming command to the Parking
*/
type Command struct {
  Cmd string
  Args []interface{}
}

/*
CommandApi describes any input which can be used to read
commands.
*/
type CommandApi interface {
  /*
  Start should initiate and loop over an input source which
  is providing commands.
  */
  Start() error
  /*
  Chan should return a channel which can be used for reading
  parsed and validated commands.
  */
  Chan() (<-chan Command, error)
}

func newCommand(cmd string, args []interface{}) Command {
  return Command{
    Cmd: cmd,
    Args: args,
  }
}
