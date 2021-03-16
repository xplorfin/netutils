package netutils

import "github.com/phayes/freeport"

// FreePortStack stack allows you to locally store a list of autogenerated ports
// it is stateful to make sure ports aren't reused
type FreePortStack struct {
	ports []int
}

// NewFreeportStack is a helper method to initalize a freeport stack
func NewFreeportStack() FreePortStack {
	return FreePortStack{}
}

func (s *FreePortStack) containsPort(port int) bool {
	for _, testPort := range s.ports {
		if testPort == port {
			return true
		}
	}
	return false
}

// GetFreePort gets a free port that has not already been used in the stack
// returns errors if they exist
func (s *FreePortStack) GetFreePort() (port int, err error) {
	port, err = freeport.GetFreePort()
	if err != nil {
		return 0, err
	}
	// TODO we need to add a loop iterator here
	if s.containsPort(port) {
		return s.GetFreePort()
	}
	s.ports = append(s.ports, port)
	return port, err
}

// GetPort gets a free port that has not already been used in the stack
// panics on error
func (s *FreePortStack) GetPort() (port int) {
	port, err := s.GetFreePort()
	if err != nil {
		panic(err)
	}
	return port
}