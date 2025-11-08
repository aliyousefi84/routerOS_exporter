# MikroTik RouterOS Exporter

![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)
![Kubernetes](https://img.shields.io/badge/Kubernetes-Deploy-success.svg)
![Prometheus](https://img.shields.io/badge/Prometheus-Exporter-orange.svg)

A simple Prometheus exporter for collecting metrics from MikroTik RouterOS devices.

## 📊 Features

- Real-time monitoring of MikroTik routers
- Comprehensive metrics collection
- Docker and Kubernetes ready
- Easy integration with Prometheus
- Lightweight and efficient

## 🎯 Collected Metrics

| Metric | Description | Type |
|--------|-------------|------|
| `mik_cpu_load` | CPU usage percentage | Gauge |
| `mik_mem_free` | Available memory in bytes | Gauge |
| `mik_hard_free` | Free disk space in bytes | Gauge |
| `mikrotik_in_traffic` | Interfaces incoming traffic | Counter |
| `mik_out_traffic` | Interfaces outgoing traffic | Counter |
| `mik_user_traffic` | User traffic | Counter |


## 🚀 Quick Start

### Prerequisites
- Docker and Docker Compose
- Access to MikroTik router with API enabled

### Using Docker Compose

1. Create `docker-compose.yml`:

```yaml
services:
     exporter:
       container_name: exporter
       restart: always
       image: aliyousefi84/mikrotik-exporter:v1.20
       environment:
         - ROUTEROS_ADDRESS=<your-mikrotik-address>
         - ROUTEROS_USER=<your-mikrotik-username>
         - ROUTEROS_PASSWORD=<your-mikrotik-password>
       ports:
         - "9200:9200"
```
### Using docker cli

    docker run -d -p 9200:9200 -e ROUTEROS_ADDRESS=<your-mikrotik-address> \ 
           -e ROUTEROS_USER=<your-mikrotik-username> \
           -e ROUTEROS_PASSWORD=<your-mikrotik-password> \
           --name mikrotik \
           aliyousefi84/mikroti-exporter:v1.20
## Deploy in kubernetes 

    kubectl apply -f kubernetes.yaml
    
