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

func (i *IotMcpServer) FetchTemperature() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mode := &serial.Mode{
			BaudRate: 9600,
		}
		port, err := serial.Open("/dev/cu.usbmodem12401", mode)
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
