package internal

import (
	"bufio"
	"context"
	"fmt"
	"os"
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
			fmt.Fprintf(os.Stderr, "Error reading from serial port: %v\n", err)
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
		fmt.Fprintf(os.Stderr, "CALLING WRITE DIGITAL...\n")

		portName := request.Params.Arguments["portName"].(string)
		pin := int(request.Params.Arguments["pin"].(float64))
		value := request.Params.Arguments["value"].(string)

		fmt.Fprintf(os.Stderr, "Values are : %s %d %s\n", portName, pin, value)

		mode := &serial.Mode{
			BaudRate: 9600,
		}
		port, err := serial.Open(portName, mode)
		if err != nil {
			return nil, err
		}
		defer port.Close()
		time.Sleep(2 * time.Second)

		pinMode := fmt.Sprintf("M,%d,OUTPUT\n", pin)
		fmt.Fprintf(os.Stderr, "PIN MODE IS BEING SET to %s", pinMode)
		_, err = port.Write([]byte(pinMode))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing pin mode: %v\n", err)
			return nil, err
		}

		time.Sleep(500 * time.Millisecond)

		for i := 0; i < 5; i++ {
			// Turn LED ON
			commandOn := fmt.Sprintf("D,%d,HIGH\n", pin)
			fmt.Fprintf(os.Stderr, "COMMAND IS BEING SET TO %s", commandOn)
			_, err = port.Write([]byte(commandOn))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing HIGH command: %v\n", err)
				return nil, err
			}

			// Wait 1 second
			time.Sleep(1 * time.Second)

			// Turn LED OFF
			commandOff := fmt.Sprintf("D,%d,LOW\n", pin)
			fmt.Fprintf(os.Stderr, "COMMAND IS BEING SET TO %s", commandOff)
			_, err = port.Write([]byte(commandOff))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing LOW command: %v\n", err)
				return nil, err
			}

			// Wait 1 second
			time.Sleep(1 * time.Second)
		}

		return mcp.NewToolResultText(fmt.Sprintf("Set pin %d to %s", pin, value)), nil
	}
}
