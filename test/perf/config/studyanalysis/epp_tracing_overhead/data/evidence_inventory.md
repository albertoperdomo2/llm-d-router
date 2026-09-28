# Evidence inventory

The study uses the following benchmark cohorts. The disposition states whether
a cohort contributes to the primary result, supplies workload sensitivity, or
was superseded.

| Cohort | Scope | Disposition |
|---|---|---|
| EPP-only broad matrix | Five workload shapes and six tracing settings | Failed high-pressure cells were replaced with clean reruns; used for non-streaming workload sensitivity |
| Streaming EPP-only matrix | Streaming equivalent of the broad matrix | Used after excluding the 800 requests/s and concurrency-128 stages at every tracing ratio |
| Fixed simulator matrix | 400 requests/s, 200/100 tokens, off/1%/10%/100%, plus pprof | Used as the primary isolated EPP result |
| Initial fixed simulator matrix | Same fixed simulator workload | Superseded because Prometheus query coverage was incomplete |
| Fixed live matrix | Eight Qwen3-0.6B model-server replicas | Used as the primary live-inference check and for live pprof attribution |
| AIPerf sampling sweep | Qwen3.6-35B-A3B, concurrency 32, off and 5% through 100%, 900-second replay | Supporting only; every run contained request errors and model execution dominated the measured response |

The MLflow CLI audit checked the status, parameters, scalar metrics, and
artifacts of every available run ID. All audited runs finished. Acceptance was
decided at the benchmark-stage level because a multi-stage run can have clean
low-pressure stages and a failed final stage.
