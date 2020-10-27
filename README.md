# githubmetaparser
A project to parse Github Meta information to extract IPs/ CIDRs for whitelisting.


## Building
1. Clone the repository
    ```shell script
    git clone https://github.com/shammishailaj/githubmetaparser.git
    ```
2. Change into the directory

    ```shell script
    cd githubmetaparser
    ```
3. Issue the build command

    ```shell script
    make build
    ```

The output si produced in an nginx compatible format so you can use the binary executable in your cron to automatically refresh the list of IPs. 