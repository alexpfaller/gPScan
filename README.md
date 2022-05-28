
![Logo](https://i.postimg.cc/N07v7BTq/white-gp-Scan-Thumb.png)


# gPScan

gPScan is a simple CLI tool, to check if specific ports are opened or closed.



## Installation

Install my-project with npm

```bash
  git clone https://github.com/alexpfaller/gPScan
  cd gPScan
```
```bash
  sudo ./gPSinstall.sh
```
This will give the gPScan binary execution permissions and will copy it in your /bin/ directory.
    
## Usage

It alway's follows this schema, where you first specify one of the two protocol's (tcp, udp), after that you can insert your destination, which you want to check for open ports(localhost, any IP adress). Last but not least, append the port that should be checked.
```sh
gPS <protocol> <destination> <port>
```

Example:
```sh
gPS tcp localhost 8080
```
To check if port 8080 on your machine is currently open, or closed.
```sh
gPS tcp 132.203.220.211 443
```
To check if port 443 (https) on another machine is open.