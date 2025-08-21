
// Add import:
"github.com/bruin-data/bruin/pkg/fluxx"

// Add to Manager struct:
Fluxx    map[string]*fluxx.Client

// Add connection method:
func (m *Manager) AddFluxxConnectionFromConfig(connection *config.FluxxConnection) error {
	m.mutex.Lock()
	if m.Fluxx == nil {
		m.Fluxx = make(map[string]*fluxx.Client)
	}
	m.mutex.Unlock()

	client, err := fluxx.NewClient(&fluxx.Config{
		Instance: connection.Instance,
		ClientId: connection.ClientId,
		ClientSecret: connection.ClientSecret,
	})
	if err != nil {
		return err
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.Fluxx[connection.Name] = client
	m.availableConnections[connection.Name] = client
	m.AllConnectionDetails[connection.Name] = connection
	return nil
}

// Add to NewManagerFromConfig:
processConnections(cm.SelectedEnvironment.Connections.Fluxx, connectionManager.AddFluxxConnectionFromConfig, &wg, &errList, &mu)
