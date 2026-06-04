# Dataset Loading Optimization

## V1

### Flow

1. Open file
2. Create JSON decoder
3. Decode JSON stream
4. Populate structs

### Libraries

* encoding/json

### Results

```text
2026/06/04 11:38:21 loaded 3000000 references in 6.855227725s
2026/06/04 11:38:21 load normalization in 49.16µs
2026/06/04 11:38:21 load mccRisk in 20.3µs
2026/06/04 11:38:21 server starting on 9999
```

---

## V2

### Flow

1. Open file
2. Read entire file into memory (`[]byte`)
3. Unmarshal JSON
4. Populate structs

### Libraries

* github.com/bytedance/sonic

### Results

```text
2026/06/04 11:38:40 loaded 3000000 references in 1.851967151s
2026/06/04 11:38:40 load normalization in 1.188963ms
2026/06/04 11:38:40 load mccRisk in 736.772µs
2026/06/04 11:38:40 server starting on 9999
```

---

# Performance Comparison

## References Dataset

| Version | Time   |
| ------- | ------ |
| V1      | 6.855s |
| V2      | 1.852s |

### Difference

```text
6.855s - 1.852s = 5.003s
```

### Improvement

```text
1.852 / 6.855 = 0.27
```

V2 requires only 27% of the original loading time.

```text
73% reduction in loading time
```

or

```text
~3.7x faster
```

---

## Normalization

| Version | Time             |
| ------- | ---------------- |
| V1      | 49.16µs          |
| V2      | 1188µs (1.188ms) |

### Difference

```text
1188µs - 49µs ≈ 1139µs
```

The operation became slower, but the absolute difference is approximately 1ms and has no practical impact on startup time.

---

## MCC Risk

| Version | Time    |
| ------- | ------- |
| V1      | 20.3µs  |
| V2      | 736.7µs |

### Difference

```text
736µs - 20µs ≈ 716µs
```

The operation became slower, but the absolute difference is less than 1ms and is negligible during application startup.

---

# Technical Analysis

Previously, JSON was decoded directly from the file stream:

```text
File
 ↓
JSON Decoder
 ↓
Structs
```

After the change:

```text
File
 ↓
ReadAll
 ↓
[]byte
 ↓
Sonic Unmarshal
 ↓
Structs
```

The startup now performs an additional memory allocation step, but Sonic processes the JSON significantly faster than the standard library decoder.
For small configuration files (`normalization.json` and `mccRisk.json`) this additional step outweighs the parsing gains.
For the large dataset (`3,000,000 references`), the faster parser more than compensates for the extra memory allocation, reducing load time by approximately 73%.

Current evidence suggests that JSON parsing is no longer the primary bottleneck. The remaining startup cost is likely dominated by:
* Gzip decompression
* Memory allocations for millions of structs
* Data copying during dataset construction
