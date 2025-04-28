package internal

import (
	"bufio"
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"log"
	"strings"
	"time"

	"go.bug.st/serial"
)

type IotMcpServer struct {
}

func NewIotMcpServer() *IotMcpServer {
	return &IotMcpServer{}
}

func (i *IotMcpServer) ReadSerialLine() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		portName := request.Params.Arguments["portName"].(string)
		mode := &serial.Mode{
			BaudRate: 9600,
		}
		port, err := serial.Open(portName, mode)
		if err != nil {
			return nil, err
		}
		defer port.Close()

		time.Sleep(2 * time.Second)
		reader := bufio.NewReader(port)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading from serial port:", err)
		}

		line = strings.TrimSpace(line)
		return mcp.NewToolResultText(line), nil
	}
}

func (i *IotMcpServer) GetPortList() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		ports, err := serial.GetPortsList()
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(strings.Join(ports, "\n")), nil
	}
}
