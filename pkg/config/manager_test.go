
// Add test case to TestConfig_AddConnection tests array:
{
	name:     "Add Fluxx connection",
	envName:  "default",
	connType: "fluxx",
	connName: "fluxx-conn",
	creds: map[string]interface{}{
		"instance": "test-instance",
		"client_id": "test-client-id",
		"client_secret": "test-client-secret",
	},
	expectedErr: false,
},

// Add test assertion case:
case "fluxx":
	assert.Len(t, env.Connections.Fluxx, 1)
	assert.Equal(t, tt.connName, env.Connections.Fluxx[0].Name)
	assert.Equal(t, tt.creds["instance"], env.Connections.Fluxx[0].Instance)
	assert.Equal(t, tt.creds["client_id"], env.Connections.Fluxx[0].ClientId)
	assert.Equal(t, tt.creds["client_secret"], env.Connections.Fluxx[0].ClientSecret)
