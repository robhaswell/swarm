package cluster

import "github.com/samalba/dockerclient"

type Container struct {
	dockerclient.Container

	Info dockerclient.ContainerInfo
	node *Node
}

func (c *Container) Node() *Node {
	return c.node
}

func (c *Container) Start() error {
	return c.node.client.StartContainer(c.Id, nil)
}

func (c *Container) Kill(signal string) error {
	return c.node.client.KillContainer(c.Id, signal)
}

func (c *Container) Stop() error {
	return c.node.client.StopContainer(c.Id, 8)
}

func (c *Container) Restart(timeout int) error {
	return c.node.client.RestartContainer(c.Id, timeout)
}

func (c *Container) Pause() error {
	return c.node.client.PauseContainer(c.Id)
}

func (c *Container) Unpause() error {
	return c.node.client.UnpauseContainer(c.Id)
}
