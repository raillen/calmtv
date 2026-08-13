package iptv

import (
	"strings"
	"testing"
)

func TestParseM3UExtractsChannelMetadata(t *testing.T) {
	channels, err := ParseM3U(strings.NewReader(`#EXTM3U
#EXTINF:-1 tvg-id="news" tvg-logo="https://logo" group-title="News",Canal News
https://stream.example/live.m3u8
`))
	if err != nil {
		t.Fatal(err)
	}
	if len(channels) != 1 || channels[0].ID != "news" || channels[0].Group != "News" || channels[0].StreamURL == "" {
		t.Fatalf("channels = %#v", channels)
	}
}

func TestParseXMLTVStreamsProgramsToCallback(t *testing.T) {
	var title string
	count := 0
	err := ParseXMLTV(strings.NewReader(`<tv><channel id="news"><display-name>News</display-name></channel><programme channel="news" start="20260813120000 -0300" stop="20260813130000 -0300"><title>Jornal</title><desc>Resumo</desc></programme></tv>`), func(channel EPGChannel) error {
		if channel.Name != "News" {
			t.Fatalf("channel = %#v", channel)
		}
		return nil
	}, func(program Program) error {
		count++
		title = program.Title
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 || title != "Jornal" {
		t.Fatalf("count=%d title=%q", count, title)
	}
}

func TestParseM3UAssignsStableIDWhenMetadataIsMissing(t *testing.T) {
	channels, err := ParseM3U(strings.NewReader("#EXTM3U\n#EXTINF:-1,News\nhttps://example.invalid/news\n"))
	if err != nil || len(channels) != 1 || channels[0].ID != "channel-1" {
		t.Fatalf("channels=%#v err=%v", channels, err)
	}
}
