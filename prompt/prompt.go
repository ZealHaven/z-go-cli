package prompt

import (
	"bufio"
	"fmt"
	"os"
)

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label)
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}

	return s
}