# PG-sniffer

Capture PostgreSQL queries.

This version displays a queries with a delay of a few seconds.

### Commands
* capture - start listening
* list - network device list

### Flags
| flag   | description                  |
|--------|------------------------------|
| port   | PostgreSQL port              |
| device | network device for capturing |


## Preparation
The first thing to do is install the dependencies

```bash
sudo apt-get install libpcap-dev
```