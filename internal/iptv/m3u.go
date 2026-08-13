package iptv

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Channel struct {
	ID        string
	Name      string
	Group     string
	Logo      string
	Number    string
	StreamURL string
}

var attributePattern = regexp.MustCompile(`([A-Za-z0-9_-]+)="([^"]*)"`)

func ParseM3U(reader io.Reader) ([]Channel, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	var channels []Channel
	var pending *Channel
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.EqualFold(line, "#EXTM3U") {
			continue
		}
		if strings.HasPrefix(line, "#EXTINF:") {
			channel := Channel{}
			for _, match := range attributePattern.FindAllStringSubmatch(line, -1) {
				switch strings.ToLower(match[1]) {
				case "tvg-id":
					channel.ID = match[2]
				case "tvg-name":
					channel.Name = match[2]
				case "tvg-logo":
					channel.Logo = match[2]
				case "tvg-chno":
					channel.Number = match[2]
				case "group-title":
					channel.Group = match[2]
				}
			}
			if index := strings.LastIndex(line, ","); index >= 0 && channel.Name == "" {
				channel.Name = strings.TrimSpace(line[index+1:])
			}
			pending = &channel
			continue
		}
		if strings.HasPrefix(line, "#") || pending == nil {
			continue
		}
		pending.StreamURL = line
		if pending.Name == "" {
			pending.Name = pending.StreamURL
		}
		if pending.ID == "" {
			pending.ID = fmt.Sprintf("channel-%d", len(channels)+1)
		}
		channels = append(channels, *pending)
		pending = nil
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read M3U: %w", err)
	}
	return channels, nil
}
