# Valuation
1. Latency (P99)
2. Detector

## Architecture
Client -> TCP :9999 -> Lb -> SCM_RIGHTS -> Api 1 e Api 2

- communication via unix.socket between lb, api1, api2

### Load balancer
- listen()
- accept()

- choose worker (round robin)

- sendmsg(SCM_RIGHTS)
- close(fd local)

### Worker (APIs)
- recebe fd
- net.FileConn()
- http.Server
- handler
- response
- close

### Flow

1. Lb -> Listener TCP -: :9999
2. Workers (APIs) -> Unix socket to Lb
3. Connection start:
    - accept()
    - get fd
    - choose worker
    - send fd
4. worker receive fd
5. worker convert fd to net.conn
6. flow continuous

