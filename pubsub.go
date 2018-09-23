package main

type PubSub struct {
	Client []Client
	Topic  []Topic
}

type Topic struct {
	Topic  string
	Client *Client
}

type Message struct {
	Action  string `json:"action"`
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

const (
	TopicAll = "all"
)

func (p *PubSub) add(client Client) *PubSub {
	p.Client = append(p.Client, client)
	p.Topic = append(p.Topic, Topic{
		Topic:  TopicAll,
		Client: &client,
	})

	return p
}

func (p *PubSub) remove(client Client) {
	for i, c := range p.Client {
		if (c.ID == client.ID) {
			ps.Topic = append(ps.Topic[:i], ps.Topic[i+1:]...)
			ps.Client = append(ps.Client[:i], ps.Client[i+1:]...)
		}
	}
}

func (p PubSub) HandleReceiveMessage(client Client) {

}
