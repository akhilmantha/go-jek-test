package api

import (
  "os"
  "bufio"
  "io"
  "strings"
)

type FileCommandApi struct {
  out chan Command
  reader io.Reader
}

func (f *FileCommandApi) Start() error {
  scanner := bufio.NewScanner(f.reader)
  for scanner.Scan() {
    inp := strings.TrimSpace(scanner.Text())
    cmd, err := parse(inp)
    if err != nil {
      close(f.out)
      return err
    }
    if cmd.Cmd == ExitCommand {
      break
    }
    f.out <- cmd
  }
  select {
  case <-f.out:
  default:
    close(f.out)
  }
  return nil
}

func (f *FileCommandApi) Chan() (chan Command, error) {
  return f.out, nil
}

func NewFileCommandApi(filename string) (CommandApi, error) {
  file, err := os.Open(filename)
  if err != nil {
    return nil, err
  }

  return &FileCommandApi{
    out: make(chan Command, 1),
    reader: file,
  }, nil
}

func NewStdinApi() CommandApi {
  return &FileCommandApi{
    out: make(chan Command, 1),
    reader: os.Stdin,
  }
}
