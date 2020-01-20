package lib

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Guild struct {
	guild *discordgo.UserGuild
	roles []*Role
}

type Role struct {
	role *discordgo.Role
}

type Discorder struct {
	session *discordgo.Session
	guilds  []Guild
}

func NewDiscorder(token string) (*Discorder, error) {
	session, err := discordgo.New("Bot " + strings.TrimSpace(token))

	if err != nil {
		return nil, err
	}

	d := &Discorder{
		session: session,
	}

	session.AddHandler(d.ready)
	session.AddHandler(d.messageCreate)

	err = session.Open()

	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Discorder) ready(s *discordgo.Session, event *discordgo.Ready) {
	log.Println("ready speghetti")

	guilds, _ := s.UserGuilds(100, "", "")

	for _, guild := range guilds {
		// g, _ := s.Guild(guild.ID)
		log.Printf("Guild: %#v\n", guild)
		roles, _ := s.GuildRoles(guild.ID)

		for _, role := range roles {
			log.Printf("Role: %#v\n", role)

		}
	}
}

func (d *Discorder) messageCreate(s *discordgo.Session, event *discordgo.MessageCreate) {
	log.Printf("Message create: %#v\n", event.Message)
}

func (d *Discorder) Close() {
	d.session.Close()
}
