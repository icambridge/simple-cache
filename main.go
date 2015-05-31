package main

import (
	"bufio"
	"log"
	"net"
	"net/textproto"
	"simplecache"
	"strings"
	"github.com/mgutz/str"
)

var commands = map[string]simplecache.Command{}

func main() {

	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	buildCommands()
	log.Println("Simple cache started")



	for {

		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			log.Println("Connection made from " + c.RemoteAddr().String())
			reader := bufio.NewReader(c)
			tp := textproto.NewReader(reader)

			for {
				line, err := tp.ReadLine()
				if err != nil {
					break
				}
				log.Println(line)
				handleCommand(c, line)
			}
			log.Println("Connection closed")
			c.Close()
		}(conn)
	}
}

func buildCommands() {

		rawCommands := simplecache.GetCommands()

		for _, command := range rawCommands {
			key := command.GetName()
			commands[key] = command
		}
}

func handleCommand(c net.Conn, line string) {
	parts := str.ToArgv(line)
	commandName := strings.ToLower(parts[0])

	if commandName == "" {
		return
	}

	command, found := commands[commandName]

	if !found {
		command = commands["notfound"]
	}


	output := command.HandlePayload(parts)

	if output == "" {
		return
	}

	writeLn(c, output)
}

func writeLn(c net.Conn, line string) {

	c.Write([]byte(line + "\r\n"))
}
