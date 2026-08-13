# Explanation — Privacy and Local AI

TV Shell has no telemetry by default.

The Smart Organizer uses deterministic local evidence before AI. When the optional local model is used:
- it is loaded on demand;
- receives only the minimum filenames/metadata necessary;
- returns a structured plan;
- cannot execute arbitrary shell commands;
- can access only approved user-media directories through the executor;
- terminates after the task.

External cloud AI is not required for normal product operation.
