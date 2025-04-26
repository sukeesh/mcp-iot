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

	temperatureModule := mcp.NewTool("get_temperature",
		mcp.WithDescription("Read temperature from Temperature sensor module"),
	)
	s.AddTool(temperatureModule, i.FetchTemperature())

	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func main() {
	mcpMain()
}
