# Pet project

## Simple web service run from systemd

web control panel for home nas server

**Config file** : srvctl.yml  set as key **binfile -c srvctl.yml**. Or use systemd srvctl.service unit file.

**Listen TCP port** : default value 8080 , or set env var - PORT

## Docker part (short memo)

```bash
docker build -t kos/srvctl
docker run -d -p 8080:8080 --name srvctl -t kos/srvctl
docker stop srvctl
```

### work in progress ...
