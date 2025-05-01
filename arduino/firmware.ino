/*
 * Arduino Command Processor
 * 
 * This sketch reads commands from the serial port and executes them.
 * Commands:
 * M,pin,mode - Sets the pin mode (M,6,OUTPUT)
 * D,pin,value - Sets the digital pin value (D,6,HIGH or D,6,LOW)
 * AB,pin,value - Sets analog value (PWM) for pin (AB,9,128)
 * BZ,pin,state - Controls a buzzer (BZ,8,ON or BZ,8,OFF)
 */

void setup() {
  Serial.begin(9600); // Initialize serial communication at 9600 baud
  Serial.println("Arduino ready to receive commands");
}

void loop() {
  if (Serial.available() > 0) {
    String command = Serial.readStringUntil('\n');
    processCommand(command);
  }
}

void processCommand(String command) {
  // Trim any whitespace
  command.trim();
  
  // Split the command by commas
  int firstComma = command.indexOf(',');
  int secondComma = command.indexOf(',', firstComma + 1);
  
  if (firstComma == -1 || secondComma == -1) {
    Serial.println("Invalid command format. Expected: TYPE,PIN,VALUE");
    return;
  }
  
  String cmdType = command.substring(0, firstComma);
  int pin = command.substring(firstComma + 1, secondComma).toInt();
  String value = command.substring(secondComma + 1);
  
  // Process the command based on type
  if (cmdType.equals("M")) {
    // Mode command
    if (value.equals("OUTPUT")) {
      pinMode(pin, OUTPUT);
      Serial.print("Set pin ");
      Serial.print(pin);
      Serial.println(" to OUTPUT mode");
    } else if (value.equals("INPUT")) {
      pinMode(pin, INPUT);
      Serial.print("Set pin ");
      Serial.print(pin);
      Serial.println(" to INPUT mode");
    } else {
      Serial.println("Invalid mode. Expected: OUTPUT or INPUT");
    }
  } else if (cmdType.equals("D")) {
    // Digital write command
    if (value.equals("HIGH")) {
      digitalWrite(pin, HIGH);
      Serial.print("Set pin ");
      Serial.print(pin);
      Serial.println(" to HIGH");
    } else if (value.equals("LOW")) {
      digitalWrite(pin, LOW);
      Serial.print("Set pin ");
      Serial.print(pin);
      Serial.println(" to LOW");
    } else {
      Serial.println("Invalid value. Expected: HIGH or LOW");
    }
  } else if (cmdType.equals("BZ")) {
    // Buzzer command
    if (value.equals("ON")) {
      pinMode(pin, OUTPUT);
      digitalWrite(pin, HIGH);
      Serial.print("Buzzer on pin ");
      Serial.print(pin);
      Serial.println(" turned ON");
    } else if (value.equals("OFF")) {
      digitalWrite(pin, LOW);
      Serial.print("Buzzer on pin ");
      Serial.print(pin);
      Serial.println(" turned OFF");
    } else {
      Serial.println("Invalid buzzer state. Expected: ON or OFF");
    }
  } else {
    Serial.println("Invalid command type. Expected: M (mode), D (digital), AB (analog), or BZ (buzzer)");
  }
} 