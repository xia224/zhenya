# Routing related

## 1. Routing protocol
1.1 SNI (Server Name Indication)
[Proxy server with SNI Doc](https://github.com/apache/pulsar/wiki/PIP-60:-Support-Proxy-server-with-SNI-routing)
![Pulsar proxy with SNI!](/Users/zzy/Desktop/pictures/SNI.png)

*Copy from other place
Figure 1: shows the layer-4 routing network activity diagram between client and application server via proxy-server (eg: ATS). In the figure, the client initiates TLS connection with the proxy-server by passing hostname of application-server in the SNI header. The proxy server examines SNI header, parses hostname of the target application server and creates an outbound connection with that application server host. Once, proxy server successfully completes TLS handshake with the application-server, the proxy creates a TLS tunnel between the client and the application server host to begin further data transactions. Many known proxy server solutions such as ATS, Fabio, Envoy, Nginx, HaProxy support SNI routing to create TLS tunnel between client and server host.