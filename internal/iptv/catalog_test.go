package iptv

import "testing"

func TestCatalogMovesAndKeepsFavorites(t *testing.T) {
	catalog := NewCatalog()
	catalog.Replace([]Channel{{ID: "one", Name: "One", StreamURL: "https://one"}, {ID: "two", Name: "Two", StreamURL: "https://two"}})
	channel, err := catalog.Move(1)
	if err != nil {
		t.Fatal(err)
	}
	if channel.ID != "two" {
		t.Fatalf("selected = %#v", channel)
	}
	if _, err := catalog.ToggleFavorite(); err != nil {
		t.Fatal(err)
	}
	if !catalog.IsFavorite(channel) {
		t.Fatal("channel was not favorited")
	}
}

func TestCatalogSelectsByRemoteNumber(t *testing.T) {
	catalog := NewCatalog()
	catalog.Replace([]Channel{{ID: "one"}, {ID: "two"}})
	channel, err := catalog.SelectNumber(2)
	if err != nil || channel.ID != "two" {
		t.Fatalf("channel=%#v err=%v", channel, err)
	}
}
