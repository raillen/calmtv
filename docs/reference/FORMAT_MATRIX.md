# Format Matrix

| Purpose | Canonical format |
|---|---|
| Human/project knowledge | Markdown |
| Atlas/project config | JSON |
| Goals | JSON |
| Workforce manifests | JSON |
| Project Intelligence | JSON |
| Runtime media/search state | SQLite |
| Working context | runtime/SQLite, ephemeral |
| IPTV playlist | M3U/M3U8 input |
| EPG | XMLTV input normalized to SQLite |
| UI layout | GtkBuilder XML implementation asset, not canonical project knowledge |
| UI styling | GTK CSS implementation asset |
| Legacy Atlas YAML | read/migration only |

Do not introduce YAML as new maintained Atlas/project configuration.
