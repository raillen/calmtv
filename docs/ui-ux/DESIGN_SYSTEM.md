# Design System — Calm TV UI

## Intent

A dark, calm, content-first 10-foot interface: fewer elements than Google TV/Kodi/Big Picture, but equally explicit focus and remote navigation.

## Visual principles
- strong hierarchy at 2–4 m viewing distance;
- graphite/dark-neutral surfaces, not absolute black everywhere;
- one restrained accent color;
- large typography and predictable spacing;
- little decoration and no dependence on blur/compositor effects;
- artwork enhances content but never blocks navigation/loading.

## Provisional tokens

```text
color.bg         #0B0D10
color.surface    #15181D
color.surface2   #1C2026
color.text       #F4F4F2
color.muted      #A8ADB5
color.accent     #E6A23C   (working accent; brand final pending)
radius.card      12px
stroke.focus     3px
space.base       8px
```

Typography: Inter or Noto Sans baseline; final packaged font decision must consider Debian availability/licensing and rendering quality.

## Scale at 1080p
- metadata/body: 18–24 px;
- card title: 24–28 px;
- section: 30–36 px;
- page title: 40–48 px;
- hero title: 48–64 px.

These are design targets, not GTK hardcoded global constants; scaling must be validated at 720p and overscan-safe TVs.
