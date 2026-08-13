package library

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type Item struct{ Path, Title, Kind string }

var supported = map[string]string{
	".mkv": "video", ".mp4": "video", ".webm": "video", ".avi": "video",
	".mp3": "audio", ".flac": "audio", ".ogg": "audio", ".wav": "audio",
}

func Scan(root string) ([]Item, error) {
	var items []Item
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		extension := strings.ToLower(filepath.Ext(entry.Name()))
		kind, ok := supported[extension]
		if !ok {
			return nil
		}
		items = append(items, Item{Path: path, Title: strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name())), Kind: kind})
		return nil
	})
	return items, err
}
