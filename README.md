# PG-sniffer

Capture PostgreSQL SQL-queries.

For traffic capturing you needed administrative privileges (for example, `sudo`)

### Commands
* capture - start listening
* list - network device list

### Flags
| flag      | description                  |
|-----------|------------------------------|
| port      | PostgreSQL port              |
| device    | network device for capturing |
| highlight | highlight SQL syntax         |


## Preparation
The first thing to do is install the dependencies

```bash
sudo apt-get install libpcap-dev
```

## Principle of operation

### Packet capturing
The movement of all data in networks is carried out in the form of packets, which are the unit of data for networks.
The term "packet" is first introduced at the network layer. The main protocols of this layer are `IP` (Internet Protocol).
Further transport layer protocols include `TCP` (Transmission Control Protocol), focused on creating a permanent connection,
UDP (User Datagram Protocol) and the application layer contains many commonly used protocols such as HTTP, FTP, IMAP, SMTP and many others.

Packet capture refers to the collection of data transmitted over a network.
Anytime the NIC receives data, it checks the packet's destination MAC address against its own.
And if the addresses match, an exception is thrown that is used to copy data from the network card buffer to the "decision center".
We look at the packet headers and pass this packet to the appropriate handler. When capturing packets, the driver also sends a copy to us. The usual capture library is `libpcap`.
and that's why we install the dependencies we need.
By the way, this library is also used by `TCPDump` and `Wireshark`, known to all of us.
We can also filter traffic to only receive the packets we need (you can see the filter in the code).
All that remains for us to do is to select a network device and receive a copy of the packets from it and encode them according to the headers.

Packet capture is typically used for network debugging or looking for performance or security anomalies
In our case, this is debugging an application working with a database.