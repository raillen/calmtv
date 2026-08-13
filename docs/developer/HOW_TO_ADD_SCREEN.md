# How-to — Add a Screen

1. Add/reuse GtkBuilder UI structure and design-system components.
2. Do not create a private focus engine.
3. Register focusable elements with the central FocusManager.
4. Bind semantic actions, not physical key codes.
5. Provide loading, empty, error and disabled states.
6. Ensure Back/Home semantics.
7. Add D-pad traversal tests and 720p/1080p visual fixtures.
8. Verify idle screen causes no unnecessary redraw/poll loop.
