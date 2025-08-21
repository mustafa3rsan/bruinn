
// Add to Connections struct:
Fluxx    []FluxxConnection    `yaml:"fluxx,omitempty" json:"fluxx,omitempty" mapstructure:"fluxx"`

// Add to AddConnection switch statement:
case "fluxx":
	var conn FluxxConnection
	if err := mapstructure.Decode(creds, &conn); err != nil {
		return fmt.Errorf("failed to decode credentials: %w", err)
	}
	conn.Name = name
	env.Connections.Fluxx = append(env.Connections.Fluxx, conn)

// Add to DeleteConnection switch statement:
case "fluxx":
	env.Connections.Fluxx = removeConnection(env.Connections.Fluxx, connectionName)

// Add to MergeFrom method:
mergeConnectionList(&c.Fluxx, source.Fluxx)
