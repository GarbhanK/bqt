# BQT

- bigquery templating cli app
- reads from a `mapping.json` file


### Setup
- create a folder `${HOME}/Documents/bqt`
- create a `mappings.json` file tith template names as keys and the desired unput as values

```bash
bqt <file> <env>

bqt users.sql live
```

**Usage**
- `file` (required)
- `env` (optional), either "live", "dev", or "staging"
- `isTest` (optional), read from the `mapping.json` in main package

**mapping.json**
- reads from `mapping.json`
- keys are the templated values and values are the what you sub in

