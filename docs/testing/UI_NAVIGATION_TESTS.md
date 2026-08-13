# UI Navigation Tests

Each core screen has a deterministic focus graph fixture.

Required assertions:
- initial focus exists;
- every actionable control is reachable;
- edge navigation is defined;
- disabled/hidden controls are skipped;
- Back closes overlays in correct order;
- focus returns sensibly after modal close;
- Home exits to Home;
- text input has a usable keyboard path;
- 720p/1080p layouts keep focus visible.

Record representative screenshots for visual regression, but do not make screenshots the only behavioral assertion.
