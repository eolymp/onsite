# Onsite Proxy

This basic SOCK5 proxy server is meant to help limit access to the Internet during programming competitions. It's easy to use but provides a limited filtering, which can be circumvented by an experienced attacker.

# Installation

Onsite Proxy should be installed on a server with access to the LAN and should be accessible from participant's workstations. 

1. Set up a dedicated server on your local network which will work as a proxy. This server must have access to the Internet and should be accessible from participant's workstations.
2. Download the latest release of Onsite Proxy from [release page](https://github.com/eolymp/onsite/releases) for your OS (Windows/Linux/macOS).
3. Download an example configuration file from [release page](https://github.com/eolymp/onsite/releases) called `config.yaml.dist`.
4. Rename `config.yaml.dist` to `config.yaml` and change the configuration to match your requirements.
5. Start Onsite Proxy
6. Configure your participant's workstations to use Onsite Proxy Server as their proxy server (see below).

For Onsite Proxy to be effective, participant's workstations should not have access to the Internet: either by being in a LAN which does not have access to the Internet, or restricts all traffic (incl. DNS) for participant's workstations. 
If the participant manages to disable proxy configuration, their workstation should not allow them to access the Internet.

## Configuration

The goal of Onsite Proxy is to be easy and lightweight, so it does not inspect traffic and only applies filtering rules on DNS resolution and network levels.

### DNS filtering

Onsite Proxy acts as a DNS resolver, it will receive DNS requests, check domain names against the allow-list and, if allowed, will proxy the request to the default DNS server. You can specify which domain names are allowed in the `config.yaml`.

```yaml
allowed_domains:
  - "*.eolymp.com"
forbidden_domains:
  - "basecamp.eolymp.com"
```

Domain names can be fully specified, or may be set as patterns with the wildcard symbol `*`. The wildcard symbol will match a single, entier, segment of domain name separated by `.`. For example, pattern `*.xyz` will match any first level subdomain of `xyz`, except for `xyz` itself or subdomains of second+ level.

If `allowed_domains` is an empty array or not specified at all, all domains except these listed in `forbidden_domains` will be resolved.

It's recommended to specify `allowed_domains` to explicitly list resources which will be accessible during the competition.

### Network filtering

Network filtering allows specifying IP addresses and ports which will be available to the participants. You can use following configuration:

```yaml
# Filter destination IPs
allowed_ip: ["1.1.1.1", "8.8.8.8"]
forbidden_ip: ["127.0.0.1"]

# Filter destination ports
allowed_ports: [80, 443]
forbidden_ports: [22]
```

These options do not support patterns and will be matched as is.

If `allowed_ip` is empty, all addresses will be allowed except these specified in `forbidden_ip`. On the other hand if `allowed_ip` is not empty, only these IP addresses will be allowed, so there is no need to specify `forbidden_ip`. Similarly, if `allowed_ports` is empty, all ports will be allowed except these specified in `forbidden_ports`.

Specifying exact IP addresses of servers might be a tedious and error-prone process, so it's possible to automatically populate allowed_ip with IP addresses received during domain name resolution. You can enable this feature by setting the option `allow_resolved_ips`.

```yaml
allow_resolved_ips: true
```

This configuration will automatically add IP addresses returned by DNS resolver to `allowed_ip`.

### Recommended

The recommended way of configuring proxy is to use these options together: `allowed_domains`, `allowed_ports` and `allow_resolved_ips` set to true. This way, the proxy server will only resolve requests to `allowed_domains` and allow traffic only to IP addresses corresponding to the allowed domains.

## Setup proxy on workstations

Each workstation should be configured to use Onsite Proxy as proxy server. Proxy server type should be set to `SOCK5`, the address should correspond to the IP address of the server where Onsite Proxy is installed and port should be set to `8000` (or other as specified in `config.yaml`). If configuration has option "Use proxy for DNS resolution", this option should be enabled.

You should be able to configure proxy server in your OS network settings, although some browsers allow to manually override proxy configuration in their own settings (eg. [Firefox](https://support.mozilla.org/en-US/kb/connection-settings-firefox)).

<p align="center">
  <img width="675" alt="Screenshot 2024-03-19 at 15 53 05" src="https://github.com/eolymp/onsite/assets/576301/3b4ea976-91d1-48d5-9511-a1e2596abf26">

<p align="center">
  <i>Firefox Proxy Configuration</i>


