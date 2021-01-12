package mmp

import (
	"os"

	//"github.com/pkg/term"
	"x6a.dev/pkg/msg"
)

func readKey() []byte {
	bytes := make([]byte, 32)

	n, err := os.Stdin.Read(bytes)
	if err != nil {
		msg.Alertf("Unable to read from stdin: %v", err)
		return nil
	}

	// msg.Debugf("\n%d bytes read: %v\n", n, bytes[0:n])

	return bytes[0:n]
}

/*
func readKey2() []byte {
	t, err := term.Open("/dev/tty")
	if err != nil {
		return nil
	}

	if err := term.RawMode(t); err != nil {
		return nil
	}

	bytes := make([]byte, 32)

	n, err := t.Read(bytes)
	if err != nil {
		return nil
	}

	// msg.Debugf("\n%d bytes read: %v\n", n, bytes)

	// if n == 3 && bytes[0] == 27 && bytes[1] == 91 {
	// 	if bytes[2] != 49 && bytes[2] != 50 {
	// 		input = bytes
	// 	}
	// }
	// if n == 3 && bytes[0] == 27 && bytes[1] == 79 {
	// 	input = bytes
	// }
	// if n == 1 {
	// 	input = append(input, bytes[0])
	// }

	// if len(input) == 0 {
	// 	input = bytes
	// }
	// input = bytes[0:n]

	if err := t.Restore(); err != nil {
		return nil
	}
	if err := t.Close(); err != nil {
		return nil
	}

	return bytes[0:n]
}
*/
/*
// Returns either an ascii code, or (if input is an arrow) a Javascript key code.
func getChar() (ascii int, keyCode int, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
		// Three-character control sequence, beginning with "ESC-[".

		// Since there are no ASCII codes for arrow keys, we use
		// Javascript key codes.
		if bytes[2] == 65 {
			// Up
			keyCode = 38
		} else if bytes[2] == 66 {
			// Down
			keyCode = 40
		} else if bytes[2] == 67 {
			// Right
			keyCode = 39
		} else if bytes[2] == 68 {
			// Left
			keyCode = 37
		}
	} else if numRead == 1 {
		ascii = int(bytes[0])
	} else {
		// Two characters read??
	}
	t.Restore()
	t.Close()
	return
}
*/
