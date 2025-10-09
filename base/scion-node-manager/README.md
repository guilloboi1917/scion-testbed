
# SCION AS Container API

**Base URL:**  
`http://<container_ip>:8080/api`

## Packet Capture Endpoints

| Method | Endpoint                | Description                        | Example |
|--------|-------------------------|------------------------------------|---------|
| POST   | `/capture/start`        | Start packet capture on interface  | `{"interface":"eth0"}` |
| POST   | `/capture/stop`         | Stop capture by ID                 | `{"capture_id":"123"}` |
| GET    | `/capture/status`       | Get current capture status         | |
| GET    | `/capture/files`        | List available pcap files          | |
| GET    | `/capture/file/{id}`    | Download specific pcap file        | _Not yet implemented_ |

## Configuration Endpoints

| Method | Endpoint                      | Description                       |
|--------|-------------------------------|-----------------------------------|
| POST   | `/config/scion`               | Upload SCION config file          |
| GET    | `/config/scion/{file}`        | Read SCION config file            |
| POST   | `/config/firewall`            | Update firewall rules             |
| POST   | `/config/scion/path-policy`   | Upload path-policy file           |
| GET    | `/config/scion/path-policy`   | List path-policy files            |
| POST   | `/config/scion/topology`      | Upload topology file              |
| GET    | `/config/scion/topology`      | Get topology file                 |

## Packet Dispatch Endpoints

| Method | Endpoint                      | Description                       | Example |
|--------|-------------------------------|-----------------------------------|---------|
| POST   | `/dispatch/ping/start`        | Start ping to IP                  | `{"dst":"10.100.0.11","count":5}` |
| POST   | `/dispatch/ping/stop`         | Stop ping                         | |
| GET    | `/dispatch/ping/files`        | List ping result files            | |
| POST   | `/dispatch/scionping/start`   | Start SCION ping                  | `{"dst":"17:ffaa:1:1","count":5}` |

---
