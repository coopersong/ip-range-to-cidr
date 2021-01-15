# ip-range-to-cidr

## Build

```bash
sh build.sh
```

## Usage

### Show help

```bash
./irtc -h
```

### Convert IP range to CIDRs

* #### IPv4

```bash
./irtc -begin 192.168.1.0 -end 192.168.1.255
```

```bash
./irtc -begin 192.168.1.3 -end 192.168.1.254
```

* #### IPv6

```bash
./irtc -begin 2408:874f:2000:100::fff -end 2408:874f:2000:1ff:ffff:ffff:ffff:ffff
```