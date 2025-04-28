package main

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/sukeesh/mcp-iot-go/internal"
	"log"
	"os"
)

func setEnvs() {
	_ = os.Getenv("KEY")
}

func mcpMain() {
	setEnvs()

	s := server.NewMCPServer(
		"IOT MCP Server",
		"0.0.1",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	i := internal.NewIotMcpServer()

	readSerialLineModule := mcp.NewTool("read_serial_line",
		mcp.WithDescription("Read a single line of data from any serial port (port and baud configurable)"),
		mcp.WithString("portName",
			mcp.Required(),
			mcp.Description("The name of the port to read the serial line from. Example: /dev/cu.usbmodem12401"),
		),
	)
	s.AddTool(readSerialLineModule, i.ReadSerialLine())

	portsListModule := mcp.NewTool("port_list",
		mcp.WithDescription("List all the ports available"))
	s.AddTool(portsListModule, i.GetPortList())

	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func main() {
	mcpMain()
}
