# MCP for IOT

<p align="center">
  <strong>Model Completion Protocol (MCP) Server for IoT Devices (Arduino)</strong>
</p>

[![Go](https://github.com/sukeesh/mcp-iot-go/workflows/Go/badge.svg)](https://github.com/sukeesh/mcp-iot-go/actions)

## Overview
MCP for IOT is a Go implementation of the [Model Completion Protocol](https://modelcontextprotocol.io/introduction) (MCP) server that allows AI models like Claude to directly interact with IoT devices, specifically Arduino boards. This server acts as a bridge between Claude and your Arduino hardware, enabling AI-controlled hardware automation.

## Features
- 🔌 List available serial ports
- 📡 Read data from Arduino's serial port
- 🚦 Control digital pins (LEDs, relays, etc.)
- 🔊 Control buzzers with on/off and duration settings
- 🤖 Seamless integration with Claude AI

## Prerequisites
- [Go](https://go.dev/doc/install) (version 1.24 or later)
- Arduino hardware with USB connection
- [Claude Desktop App](https://claude.ai/download) or compatible MCP client

## Installation

### Build from Source
```bash
git clone https://github.com/sukeesh/mcp-iot-go.git
cd mcp-iot-go
go install
```

The binary will be installed to your GOBIN directory, which should be in your PATH.

## Arduino Setup

1. Connect your Arduino to your computer
2. Upload the provided Arduino firmware:
   ```
   arduino/firmware.ino
   ```
   
This firmware accepts commands via serial communication:
- `M,pin,mode` - Sets pin mode (e.g., `M,6,OUTPUT`)
- `D,pin,value` - Controls digital pins (e.g., `D,6,HIGH` or `D,6,LOW`)
- `BZ,pin,state` - Controls buzzers (e.g., `BZ,8,ON` or `BZ,8,OFF`)

## Configuration with MCP Client

Configure Claude Desktop:
1. Open Claude Desktop → Settings → Developer → Edit Config
2. Add the following to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "iot": {
      "command": "<path-to-mcp-iot-go-binary>",
      "env": {}
    }
  }
}
```

3. Restart Claude Desktop

## Available Tools

| Tool | Description |
|------|-------------|
| `port_list` | Lists all available serial ports |
| `read_serial_line` | Reads a single line of data from a specified serial port |
| `write_digital` | Writes a digital value (HIGH/LOW) to a pin on a specified port |
| `buzzer_control` | Controls a buzzer connected to an Arduino pin with optional duration |

## Usage Examples

After setup, you can interact with your Arduino hardware directly through Claude:

- "List all available serial ports"
- "Read data from the Arduino on port /dev/cu.usbmodem12401"
- "Turn on the red LED on pin 6"
- "Activate the buzzer on pin 8 for 2 seconds"

## Development

### Project Structure
- `main.go` - MCP server setup and tool definitions
- `internal/tools.go` - Tool implementation for Arduino communication
- `arduino/firmware.ino` - Arduino firmware for processing commands

### Dependencies
- [mcp-go](https://github.com/mark3labs/mcp-go) - Go library for MCP implementation
- [go.bug.st/serial](https://pkg.go.dev/go.bug.st/serial) - Serial port communication library


## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.
