package ravendb

import (
	"encoding/json"
	"net/http"
)

type PutClientConfigurationOperation struct {
	configuration *ClientConfiguration
	Command       *PutClientConfigurationCommand
}

func NewPutClientConfigurationOperation(configuration *ClientConfiguration) (*PutClientConfigurationOperation, error) {
	if configuration == nil {
		return nil, NewIllegalArgumentException("Configuration cannot be null")
	}

	return &PutClientConfigurationOperation{
		configuration: configuration,
	}, nil
}

func (o *PutClientConfigurationOperation) getCommand(conventions *DocumentConventions) RavenCommand {
	o.Command = NewPutClientConfigurationCommand(conventions, o.configuration)
	return o.Command
}

var (
	_ RavenCommand = &PutClientConfigurationCommand{}
)

type PutClientConfigurationCommand struct {
	*RavenCommandBase

	configuration string
}

func NewPutClientConfigurationCommand(conventions *DocumentConventions, configuration *ClientConfiguration) *PutClientConfigurationCommand {
	panicIf(conventions == nil, "conventions cannot be null")
	panicIf(configuration == nil, "configuration cannot be null")

	d, err := json.Marshal(configuration)
	panicIf(err != nil, "json.Marshal failed with %s", err)
	cmd := &PutClientConfigurationCommand{
		RavenCommandBase: NewRavenCommandBase(),

		configuration: string(d),
	}
	return cmd
}

func (c *PutClientConfigurationCommand) createRequest(node *ServerNode) (*http.Request, error) {
	url := node.getUrl() + "/databases/" + node.getDatabase() + "/admin/configuration/client"
	return NewHttpPut(url, c.configuration)
}
