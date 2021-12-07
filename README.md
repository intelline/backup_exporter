# Prometheus xtrabackup exporter

### How to run it:
```bash
// setup and cnnfig pyxbackup to make regular backups
git clone https://github.com/dotmanila/pyxbackup.git > /opt/pyxbackup/

// clone this repo
git clone https://github.com/intelline/backup_exporter.git > opt/backups-collector

// create config
touch hostFile.json
echo '{"hostname": "your_host_name"}' > hostFile.json

// run program
go run collect.go

// or create a service file 
systemctl enable /opt/backups-collector/backups-collector.service 
systemctl restart backups-collector.service
```

### Dependencies:
- https://github.com/dotmanila/pyxbackup
