# metaparser
A project to parse Github Meta information to extract IPs/ CIDRs for whitelisting.


## Building
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

The output is produced in an nginx compatible format so you can use the binary executable in your cron to automatically refresh the list of IPs. 
