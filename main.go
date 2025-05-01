package main

import (
	"log"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/sukeesh/mcp-iot-go/internal"
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

	writeDigitalModule := mcp.NewTool("write_digital",
		mcp.WithDescription("Write a digital value to a pin on a port. Example: /dev/cu.usbmodem12401,1,HIGH"),
		mcp.WithString("portName",
			mcp.Required(),
			mcp.Description("The name of the port to write the digital value to. Example: /dev/cu.usbmodem12401"),
		),
		mcp.WithNumber("pin",
			mcp.Required(),
			mcp.Description("The pin to write the digital value to. Example: 1, 7, 6"),
		),
		mcp.WithString("value",
			mcp.Required(),
			mcp.Description("The value to write to the pin. Example: HIGH"),
		),
	)
	s.AddTool(writeDigitalModule, i.WriteDigital())

	buzzerControlModule := mcp.NewTool("buzzer_control",
		mcp.WithDescription("Control a buzzer connected to an Arduino pin"),
		mcp.WithString("portName",
			mcp.Required(),
			mcp.Description("The name of the port to control the buzzer on. Example: /dev/cu.usbmodem12401"),
		),
		mcp.WithNumber("pin",
			mcp.Required(),
			mcp.Description("The pin number where the buzzer is connected. Example: 8"),
		),
		mcp.WithString("state",
			mcp.Required(),
			mcp.Description("The state to set the buzzer to. Either 'ON' or 'OFF'"),
		),
		mcp.WithNumber("duration",
			mcp.Description("Duration in milliseconds to keep the buzzer on (if state is 'ON'). Set to 0 for no auto-off."),
		),
	)
	s.AddTool(buzzerControlModule, i.BuzzerControl())

	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func main() {
	mcpMain()
}
