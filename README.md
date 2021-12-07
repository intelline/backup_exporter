# Prometheus xtrabackup exporter

### How to run it:
```bash
//login to your server
git clone https://github.com/dotmanila/pyxbackup.git > /opt/pyxbackup/
// create config
touch hostFile.json
echo '{"hostname": "your_host_name"}' > hostFile.json

// run program
go run collect.go


```

### Dependencies:
- https://github.com/dotmanila/pyxbackup
