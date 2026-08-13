# Media Playback Tests

Fixtures should cover:
- H.264/AAC MP4;
- H.264/H.265 MKV as hardware permits;
- multiple audio tracks/languages;
- subtitles;
- HLS;
- IPTV channel change;
- seek/end-of-file;
- broken/unreachable source;
- network stall/recovery;
- hardware decode enabled/unsupported.

Measure dropped frames and CPU on the reference low-end machine. Do not claim codec support solely because FFmpeg can software-decode it on a development workstation.
