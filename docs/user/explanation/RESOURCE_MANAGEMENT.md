# Explanation — Why Apps Do Not Stay Open

Calm TV treats 2 GB RAM as a hard product constraint. The default policy allows one heavy foreground app plus a small number of explicitly permitted background tasks such as downloads or audio.

When switching apps, the current app may:
1. save state;
2. terminate;
3. release memory;
4. restore that state when reopened.

This can look like multitasking while avoiding the RAM cost of keeping Firefox, RetroArch, IPTV and media apps alive simultaneously. A stopped process is preferred over swapping heavily to an old HDD.
