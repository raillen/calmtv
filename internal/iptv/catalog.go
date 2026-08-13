package iptv

import "fmt"

type Catalog struct {
	channels  []Channel
	selected  int
	favorites map[string]bool
}

func NewCatalog() *Catalog { return &Catalog{favorites: make(map[string]bool)} }
func (c *Catalog) Replace(channels []Channel) {
	c.channels = append([]Channel(nil), channels...)
	c.selected = 0
}
func (c *Catalog) Channels() []Channel { return append([]Channel(nil), c.channels...) }
func (c *Catalog) Selected() (Channel, error) {
	if len(c.channels) == 0 {
		return Channel{}, fmt.Errorf("nenhum canal IPTV importado")
	}
	return c.channels[c.selected], nil
}
func (c *Catalog) Move(delta int) (Channel, error) {
	if len(c.channels) == 0 {
		return Channel{}, fmt.Errorf("nenhum canal IPTV importado")
	}
	c.selected = (c.selected + delta) % len(c.channels)
	if c.selected < 0 {
		c.selected += len(c.channels)
	}
	return c.channels[c.selected], nil
}

func (c *Catalog) SelectNumber(number int) (Channel, error) {
	if len(c.channels) == 0 {
		return Channel{}, fmt.Errorf("nenhum canal IPTV importado")
	}
	if number < 1 || number > len(c.channels) {
		return Channel{}, fmt.Errorf("número de canal inválido: %d", number)
	}
	c.selected = number - 1
	return c.channels[c.selected], nil
}
func (c *Catalog) ToggleFavorite() (Channel, error) {
	channel, err := c.Selected()
	if err != nil {
		return Channel{}, err
	}
	c.favorites[channel.ID+"\x00"+channel.StreamURL] = !c.favorites[channel.ID+"\x00"+channel.StreamURL]
	return channel, nil
}
func (c *Catalog) IsFavorite(channel Channel) bool {
	return c.favorites[channel.ID+"\x00"+channel.StreamURL]
}
