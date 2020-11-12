# metaparser
A project to parse meta information to extract IPs/ CIDRs from Github, Cloudflare and AWS (and more) for whitelisting. The output is produced in an nginx compatible format so you can use the binary executable in your cron to automatically refresh the list of IPs.

This project directly utilizes the following open sources projects:
1. [Cobra](https://github.com/spf13/cobra) for implementing the CLI interactions
2. [Go-Resty](https://github.com/go-resty/resty) as its HTTP client library
3. [Govvv](https://github.com/ahmetb/govvv) to add version information during its build process
4. [Logrus](https://github.com/sirupsen/logrus) as its logging library
5. [Dotsql](https://github.com/gchaincl/dotsql) for SQL migrations (not being used currently)
6. [Gopsutil - CPU](https://github.com/shirou/gopsutil/cpu) for CPU information (not being used currently)
7. [Gopsutil - Load](https://github.com/shirou/gopsutil/load) for system load information (not being used currently)
8. [Viper](https://github.com/spf13/viper) for reading configuration files
9. [Times](https://github.com/djherbis/times) for file times (atime, mtime, ctime, btime)

Code tries to conform to the [Golang Standards Project layout](https://github.com/golang-standards/project-layout) template

## Building
0. Install [Govvv](https://github.com/ahmetb/govvv)

1. Clone the repository
    ```shell script
    git clone https://github.com/shammishailaj/metaparser.git
    ```
2. Change into the directory

    ```shell script
    cd metaparser
    ```
3. Issue the build command

    ```shell script
    make build
    ``` 

#### Or, grab the [latest release](https://github.com/shammishailaj/metaparser/releases/latest) from the [releases](https://github.com/shammishailaj/metaparser/releases) page  

## Command Reference

1. `help`

To display help about a command

Invoked By: `metaparser help [command]`

For detailed documentation, use the inbuilt command `docs` to generate the documentation.
For help on using the `docs` command use:

```shell script
metaparser help docs
```
