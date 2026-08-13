package iptv

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
	"time"
)

type EPGChannel struct {
	ID   string
	Name string
}
type Program struct {
	ChannelID   string
	Title       string
	Description string
	Start       time.Time
	End         time.Time
}

// ParseXMLTV consumes tokens incrementally so a large guide never becomes one
// in-memory XML tree. The callback owns persistence and may discard programs.
func ParseXMLTV(reader io.Reader, onChannel func(EPGChannel) error, onProgram func(Program) error) error {
	decoder := xml.NewDecoder(reader)
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("read XMLTV: %w", err)
		}
		start, ok := token.(xml.StartElement)
		if !ok {
			continue
		}
		switch start.Name.Local {
		case "channel":
			var value struct {
				ID          string `xml:"id,attr"`
				DisplayName string `xml:"display-name"`
			}
			if err := decoder.DecodeElement(&value, &start); err != nil {
				return fmt.Errorf("decode XMLTV channel: %w", err)
			}
			if onChannel != nil {
				if err := onChannel(EPGChannel{ID: value.ID, Name: strings.TrimSpace(value.DisplayName)}); err != nil {
					return err
				}
			}
		case "programme":
			var value struct {
				ChannelID   string `xml:"channel,attr"`
				Start       string `xml:"start,attr"`
				Stop        string `xml:"stop,attr"`
				Title       string `xml:"title"`
				Description string `xml:"desc"`
			}
			if err := decoder.DecodeElement(&value, &start); err != nil {
				return fmt.Errorf("decode XMLTV programme: %w", err)
			}
			program := Program{ChannelID: value.ChannelID, Title: strings.TrimSpace(value.Title), Description: strings.TrimSpace(value.Description)}
			program.Start, err = parseXMLTVTime(value.Start)
			if err != nil {
				return fmt.Errorf("decode XMLTV start: %w", err)
			}
			program.End, err = parseXMLTVTime(value.Stop)
			if err != nil {
				return fmt.Errorf("decode XMLTV stop: %w", err)
			}
			if onProgram != nil {
				if err := onProgram(program); err != nil {
					return err
				}
			}
		}
	}
}

func parseXMLTVTime(value string) (time.Time, error) {
	value = strings.TrimSpace(value)
	for _, layout := range []string{"20060102150405 -0700", "20060102150405Z0700", "20060102150405"} {
		if parsed, err := time.Parse(layout, value); err == nil {
			return parsed, nil
		}
	}
	return time.Time{}, fmt.Errorf("invalid XMLTV time %q", value)
}
