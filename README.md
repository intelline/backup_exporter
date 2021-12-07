# Prometheus xtrabackup exporter

### How to run it:
```bash
touch hostFile.json
echo '{"hostname": "your_host_name"}' > hostFile.json

go run collect.go
```

### Dependencies:
- https://github.com/dotmanila/pyxbackup
