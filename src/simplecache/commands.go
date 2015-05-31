package simplecache

import "fmt"

func GetCommands() []Command {
  cache := getCache()

  return []Command{
    GetCommand{cache},
    SetCommand{cache},
    NotFoundCommand{},
  }
}


type Command interface {
  GetName() string
  HandlePayload(payload []string) string
}


type GetCommand struct {
  cache Cache
}

func (cmd GetCommand) GetName() string {
  return "get"
}

func (cmd GetCommand) HandlePayload(payload []string) string {
		key := payload[1]
		value, err := cmd.cache.Get(key)
    strLen := len(value)
    if err != nil {
      return "$-1"
    }

		return fmt.Sprintf("$%d\r\n%s", strLen, value)
}

type SetCommand struct {
  cache Cache
}

func (cmd SetCommand) GetName() string {
  return "set"
}

func (cmd SetCommand) HandlePayload(payload []string) string {

    if len(payload) < 3 {
      return "-ERR syntax error"
    }

    key := payload[1]
    value := payload[2]

  	cmd.cache.Set(key, value)

    return fmt.Sprintf("+OK")
}


type NotFoundCommand struct {
}

func (cmd NotFoundCommand) GetName() string {
  return "notfound"
}

func (cmd NotFoundCommand) HandlePayload(payload []string) string {
  	return fmt.Sprintf("-ERR unknown command '%s'", payload[0])
}
