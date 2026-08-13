# Deprecation and Support Lifecycle

Before removing a user-visible capability, config key or data format:
1. document replacement/migration;
2. mark deprecation in release notes;
3. preserve data where practical;
4. define minimum supported migration source;
5. update support/compatibility docs.

Architecture-only internal interfaces may evolve faster before V1, but accepted ADRs must be amended rather than silently ignored.
