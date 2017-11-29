# New Relic Infrastructure Integration for Apache

New Relic Infrastructure Integration for Apache captures critical performance
metrics and inventory reported by Apache web server.

Inventory data is obtained using `httpd` command in RedHat family distributions
and `apache2ctl` command in Debian family distributions.
Metrics data is obtained doing HTTP requests to `/status` endpoint, provided by
`mod_status` Apache module.

## Usage
This is the description about how to run the Apache Integration with New Relic
Infrastructure agent, so it is required to have the agent installed
(see
[agent installation](https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure/installation/install-infrastructure-linux)).

In order to use the Apache Integration it is required to configure
`apache-config.yml.sample` file. Firstly, rename the file to
`apache-config.yml`. Then, depending on your needs, specify all instances that
you want to monitor with correct arguments.

## Integration development usage

Assuming that you have the source code and Go tool installed you can build and run the Apache Integration locally.
* After cloning this repository, go to the directory of the Apache Integration and build it
```bash
$ make
```
* The command above will execute the tests for the Apache Integration and build an executable file called `nr-apache` under `bin` directory. Run `nr-apache`:
```bash
$ ./bin/nr-apache
```
* If you want to know more about usage of `./bin/nr-apache` check
```bash
$ ./bin/nr-apache -help
```

For managing external dependencies [govendor tool](https://github.com/kardianos/govendor) is used. It is required to lock all external dependencies to specific version (if possible) into vendor directory.
