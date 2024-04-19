# AIP
**Building the Image**
```bash
docker build -t gmsbackend .
```
**run project**
```bash
docker run -d --name gmsbackend -p 8000:8000 gmsbackend
```