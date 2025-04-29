package main

import (
	"fmt"
	"log"
	"time"

	"go.bug.st/serial"
)

func main() {
	// List available ports
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Available ports:")
	for _, port := range ports {
		log.Println(port)
	}

	mode := &serial.Mode{
		BaudRate: 9600,
	}

	// Update this to match your actual port
	portName := "/dev/cu.usbmodem12401"
	log.Printf("Connecting to %s...\n", portName)

	port, err := serial.Open(portName, mode)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// Wait for connection to establish
	log.Println("Waiting for connection to establish...")
	time.Sleep(2 * time.Second)

	// Set pin mode to OUTPUT
	pinMode := fmt.Sprintf("M,6,OUTPUT\n")
	_, err = port.Write([]byte(pinMode))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Set pin 6 to OUTPUT mode")

	// Blink the LED 5 times
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

	log.Println("Program completed")
}
