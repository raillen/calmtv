# Player and Overlays

Player chrome is hidden during steady playback and appears on input.

Required controls:
- timeline/position;
- play/pause;
- audio track;
- subtitle track;
- subtitle offset where supported;
- quality/source when meaningful;
- exit/back.

Quick Settings slides from the right and may expose volume/output, network status, subtitles, display and sleep timer without leaving content.

Overlays must not keep expensive redraw loops alive. Animation ends in a static frame.
