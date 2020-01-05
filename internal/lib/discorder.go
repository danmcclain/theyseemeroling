package lib

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Discorder struct {
	session *discordgo.Session
}

func NewDiscorder(token string) (*Discorder, error) {
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		return nil, err
	}

	d := &Discorder{
		session: session,
	}

	session.AddHandler(d.ready)

	err = session.Open()

	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Discorder) ready(s *discordgo.Session, event *discordgo.Ready) {
	log.Print("ready speghetti")
}

func (d *Discorder) Close() {
	d.session.Close()
}
