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
	TopicAll    = "all"
	PUBLISH     = "publish"
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
)

func (p *PubSub) publish(message Message) {
	for _, t := range p.Topic {
		if (t.Topic == message.Topic) {
			t.Client.Conn.WriteMessage(1, []byte(message.Message))
		}
	}
}

func (p *PubSub) add(client Client) *PubSub {
	p.Client = append(p.Client, client)
	p.addTopic(client, TopicAll)
	return p
}

func (p *PubSub) addTopic(client Client, topic string) {
	for _, t := range p.Topic {
		if (t.Client.ID == client.ID && t.Topic == topic) {
			return
		}
	}

	p.Topic = append(p.Topic, Topic{
		Topic:  topic,
		Client: &client,
	})
}

func (p *PubSub) unTopic(client Client, topic string) {
	for i, t := range p.Topic {
		if (t.Client.ID == client.ID && t.Topic == topic) {
			ps.Topic = append(ps.Topic[:i], ps.Topic[i+1:]...)
		}
	}
}

func (p *PubSub) remove(client Client) {
	for i, c := range p.Client {
		if (c.ID == client.ID) {
			ps.Topic = append(ps.Topic[:i], ps.Topic[i+1:]...)
			ps.Client = append(ps.Client[:i], ps.Client[i+1:]...)
		}
	}
}

func (p *PubSub) HandleReceiveMessage(client Client, message Message) {
	switch message.Action {
	case PUBLISH:
		p.publish(message)
		break
	case SUBSCRIBE:
		p.addTopic(client, message.Topic)
		break
	case UNSUBSCRIBE:
		p.unTopic(client, message.Topic)
		break
	}
}
