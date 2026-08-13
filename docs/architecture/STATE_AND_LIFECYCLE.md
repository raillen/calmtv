# State and Lifecycle

## Persistent user state
SQLite stores application/runtime user state such as history, favorites, media positions, configured sources and lightweight library metadata.

## App restore state
Each app exposes a compact restorable state:
- route/screen;
- selected item;
- active content ID;
- playback position when applicable;
- user filter/search state where useful.

Do not serialize entire widget trees or opaque process memory.

## Lifecycle states

`STOPPED → STARTING → FOREGROUND → BACKGROUND_ALLOWED | SAVING → STOPPING → STOPPED`

Crash/timeout paths transition to a recoverable error state and release resources.

## System sleep/reboot
AppManager requests state saves before system power operations when time permits. logind remains the authority for actual suspend/reboot/poweroff.
