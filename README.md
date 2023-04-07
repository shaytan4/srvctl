# Pet project

## Simple web service run from systemd

control panel for home nas server

got contol functions from **srvctl.yml** and show buttons on web page

service run on custom port(srvctl.service) - set env var(PORT), or default port 8080

## Docker part (short memo)

```bash
docker build -t kos/srvctl
docker run -d -p 8080:8080 --name srvctl -t kos/srvctl
docker stop srvctl
```


### work in progress ...
