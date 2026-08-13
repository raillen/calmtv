# Motion and UI Performance

Allowed global motion primitives:
- focus;
- fade;
- slide;
- expand/collapse.

Provisional durations:
- focus: 100–140 ms;
- fade: 140–180 ms;
- slide: 160–220 ms;
- modal: 180–240 ms.

Avoid:
- continuous idle animation;
- autoplay carousels;
- permanent 60 FPS redraw;
- blur/compositor dependence;
- particle backgrounds;
- large animated video heroes.

Home at rest should be almost entirely event-driven.
