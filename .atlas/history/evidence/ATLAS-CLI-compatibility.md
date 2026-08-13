# Atlas CLI compatibility note

`atlas validate atlas` and `atlas docs check` pass against the repository.
The installed CLI is a legacy 1.0 build that validates Goal evidence using a
YAML-era evidence schema; it rejects the JSON Goal/evidence contract required
by this checkout even when the JSON files and canonical evidence are valid.
The compatibility YAML projections are kept generated/read-only and are not
used as canonical state. This discrepancy is recorded rather than masking it
by marking incomplete Goals as passed.
