package system

import (
	"fmt"

	"github.com/containers/podman-tui/pdcs/registry"
	"github.com/containers/podman-tui/ui/style"
)

// UpdateData retrieves connections list data.
func (sys *System) UpdateData() {
	destinations := sys.connectionListFunc()

	sys.connectionList.mu.Lock()

	sys.connectionList.report = destinations

	sys.connectionList.mu.Unlock()

	sys.udpateConnectionDataStatus()
}

func (sys *System) udpateConnectionDataStatus() {
	sys.connectionList.mu.Lock()
	defer sys.connectionList.mu.Unlock()

	name := registry.ConnectionName()
	status := registry.ConnectionStatus()

	for i := range sys.connectionList.report {
		if sys.connectionList.report[i].Name == name {
			sys.connectionList.report[i].Status = status

			return
		}
	}
}

func (sys *System) getConnectionsData() []registry.Connection {
	sys.connectionList.mu.Lock()
	defer sys.connectionList.mu.Unlock()

	destReport := sys.connectionList.report

	return destReport
}

type connectionItemStatus struct {
	status registry.ConnStatus
}

func (connStatus connectionItemStatus) StatusString() string {
	var status string

	switch connStatus.status {
	case registry.ConnectionStatusConnected:
		status = fmt.Sprintf("%s %s", style.HeavyGreenCheckMark, "connected")
	case registry.ConnectionStatusConnectionError:
		status = fmt.Sprintf("%s %s", style.HeavyRedCrossMark, "connection error")
	}

	return status
}
