# Focus and Remote Navigation

Focus is the most important UI state.

## Central FocusManager
- owns directional navigation;
- guarantees a visible focused element;
- remembers logical focus per screen/row where useful;
- handles disabled/hidden items;
- defines entry/exit targets for overlays;
- exposes focus transitions for automated tests.

No standard screen may implement an independent D-pad graph.

## Focus appearance
Selected content should combine:
- modest scale increase (~1.04–1.06);
- strong 2–3 px outline/accent;
- slight elevation/shadow that does not require a compositor;
- label/metadata emphasis.

Color alone is insufficient.

## Back behavior
Back follows a strict stack:
1. close transient overlay/menu;
2. exit nested detail;
3. return to parent;
4. from root destination, reveal Home/rail behavior rather than quitting the Shell.

## Focus safety
Every screen must pass:
- no focus trap;
- no invisible focus;
- predictable edge behavior;
- correct restoration after modal close;
- keyboard-only completion.
