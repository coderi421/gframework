package net

// IsValidPort check if the port is legal. 0 is considered as a non valid port.
func IsValidPort(port int) bool {
	return port > 0 && port < 65535
}
