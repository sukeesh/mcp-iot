package internal

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

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

func (i *IotMcpServer) WriteDigital() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		portName := request.Params.Arguments["portName"].(string)
		pin := int(request.Params.Arguments["pin"].(float64))
		value := request.Params.Arguments["value"].(string)

		mode := &serial.Mode{
			BaudRate: 9600,
		}
		port, err := serial.Open(portName, mode)
		if err != nil {
			return nil, err
		}
		defer port.Close()
		time.Sleep(2 * time.Second)

		pinMode := fmt.Sprintf("M,6,OUTPUT\n")
		_, err = port.Write([]byte(pinMode))
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(500 * time.Millisecond)

		for i := 0; i < 5; i++ {
			// Turn LED ON
			log.Println("Setting pin 6 HIGH")
			commandOn := fmt.Sprintf("D,6,HIGH\n")
			_, err = port.Write([]byte(commandOn))
			if err != nil {
				log.Fatal(err)
			}

			// Wait 1 second
			time.Sleep(1 * time.Second)

			// Turn LED OFF
			log.Println("Setting pin 6 LOW")
			commandOff := fmt.Sprintf("D,6,LOW\n")
			_, err = port.Write([]byte(commandOff))
			if err != nil {
				log.Fatal(err)
			}

			// Wait 1 second
			time.Sleep(1 * time.Second)
		}

		return mcp.NewToolResultText(fmt.Sprintf("Set pin %d to %s", pin, value)), nil
	}
}
