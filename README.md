# BQT

- bigquery templating cli app
- reads from `mapping.json`

```bash
bqt file <env>

bqt users.sql live
```

**Usage**
- `file` (required)
- `env` (optional), either "live", "dev", or "staging"

**mapping.json**
- reads from `mapping.json`
- keys are the templated values and values are the what you sub in

