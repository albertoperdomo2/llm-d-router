# EPP pprof artifacts

The `simulator` and `live` directories contain the raw CPU and heap profiles,
per-pod `capture.json`, and run-level `capture-summary.json` from the 100%
sampling diagnostic runs. Each capture covered the only ready EPP replica for
30 seconds and completed successfully. `SHA256SUMS` records the raw profile
checksums.

The profiles can be inspected with the pprof version pinned by llm-d-router:

```sh
go run github.com/google/pprof@v0.0.0-20260402051712-545e8a4df936 -top simulator/cpu.pprof
go run github.com/google/pprof@v0.0.0-20260402051712-545e8a4df936 -top live/cpu.pprof
```

For the simulator profile, 1.17 of 26.81 sampled CPU-seconds, or 4.36%, had at
least one router tracing or OpenTelemetry frame on the stack. The tracer handle
lookup had 0.28 cumulative CPU-seconds, the batch processor had 0.40, and the
exporter had 0.36.

For the live profile, 1.16 of 36.80 sampled CPU-seconds, or 3.15%, matched the
same filter. The tracer handle lookup had 0.35 cumulative CPU-seconds, the batch
processor had 0.48, and the exporter had 0.40.

The union measurement can be reproduced with:

```sh
go tool pprof -top \
  -focus='go.opentelemetry.io/otel|pkg/common/observability/tracing' \
  simulator/cpu.pprof
```

The union counts each matching sample once. The individual function values are
cumulative call-stack attributions and must not be added. These measurements do
not assign all runtime and garbage-collection cost caused by trace allocations
to OpenTelemetry.
