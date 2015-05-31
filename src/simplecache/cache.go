package simplecache

import (
  "errors"
)

func getCache() Cache {

  emptyMap := map[string]string{}
  return Cache{data: emptyMap, locked: false}
}

type Cache struct {
  data map[string]string
  locked bool
}

func (c Cache) wait() {
    for {
      if (!c.locked) {
        break
      }
    }
}

func (c Cache) Get(key string) (string, error) {
  c.wait()
  c.locked = true
	value, found := c.data[key]
	c.locked = false
  if !found {
		return "", errors.New("Not found")
	}

	return value, nil
}

func (c Cache) Set(key string, value string) {
  c.wait()
  c.locked = true
	c.data[key] = value
  c.locked = false
}
