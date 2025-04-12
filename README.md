# Alt+Q Key Monitor

A simple Go program that monitors and counts the number of times the 'Q' key is pressed while the Alt key is held down.

## Features

- Detects when Alt key is pressed (both left and right Alt keys)
- Counts the number of 'Q' key presses while Alt is held
- Provides real-time feedback in the console
- Graceful exit handling with Ctrl+C

## Prerequisites

- Go 1.x installed on your system
- Windows operating system (uses Windows-specific system calls)

## Installation

1. Clone this repository or download the source code
2. Navigate to the project directory
3. Build the program:
   ```bash
   go build
   ```

## Usage

1. Run the compiled executable:
   ```bash
   ./alt-q-monitor
   ```
   or
   ```bash
   go run main.go
   ```

2. The program will start monitoring for Alt+Q combinations
3. Press and hold Alt, then press 'Q' multiple times
4. Release Alt to see the count of 'Q' presses
5. Press Ctrl+C to exit the program

## Example Output

```
Alt+Q Key Monitor
Press Alt and then press 'q' multiple times
Release Alt to see the count of 'q' presses
Press Ctrl+C to exit

Alt pressed - tracking 'q' presses
'q' pressed while Alt is held (count: 1)
'q' pressed while Alt is held (count: 2)
Alt released - 'q' was pressed 2 times while Alt was held
```

## Technical Details

The program uses Windows system calls through the `user32.dll` to monitor keyboard input. It specifically tracks:
- VK_LMENU (Left Alt key)
- VK_RMENU (Right Alt key)
- VK_Q (Q key)

The program polls the keyboard state every 20 milliseconds to detect key presses.

## License

This project is open source and available under the MIT License. 