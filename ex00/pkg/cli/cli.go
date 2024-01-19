// package cli
package cli

import (
	"fmt"
	"strconv"
)

func ParseArgs(args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("usage: %s <port>", args[0])
	}
	port := args[1]
	portNum, err := strconv.Atoi(port)

	// DYNAMIC AND/OR PRIVATE PORTS番号でない場合はエラー
	if err != nil || portNum < 1024 || portNum > 65535 {
		return "", fmt.Errorf("invalid port: %s. Please use a port number between 1024 and 65535", port)
	}
	return port, nil
}