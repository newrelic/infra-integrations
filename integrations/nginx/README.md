# New Relic Infrastructure Integration for NGINX
New Relic Infrastructure Integration for NGINX captures critical performance metrics and inventory reported by NGINX server. There is an open source and a commercial version of NGINX, both supported by this integration.

Inventory data is obtained from the configuration files and metrics from the status modules.

<!---
See [metrics]() or [inventory]() for more details about collected data and review [dashboard]() in order to know how the data is presented.
--->

## Configuration
* Depending on which NGINX edition you use please update your configuration enabling
  * [HTTP stub status module](http://nginx.org/en/docs/http/ngx_http_stub_status_module.html) for NGINX Open Source
  * [HTTP status module](http://nginx.org/en/docs/http/ngx_http_status_module.html) for NGINX Plus

## Installation
* download an archive file for the NGINX Integration
* extract `nginx-definition.yml` and `/bin` directory into `/var/db/newrelic-infra/newrelic-integrations`
* add execute permissions for the binary file `nr-nginx` (if required) 
* extract `nginx-config.yml.sample` into `/etc/newrelic-infra/integrations.d`

## Usage
This is the description about how to run the NGINX Integration with New Relic Infrastructure agent, so it is required to have the agent installed (see [agent installation](https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure/installation/install-infrastructure-linux)).

In order to use the NGINX Integration it is required to configure `nginx-config.yml.sample` file. Firstly, rename the file to `nginx-config.yml`. Then, depending on your needs, specify all instances that you want to monitor. Once this is done, restart the Infrastructure agent.

You can view your data in Insights by creating your own custom NRQL queries. To do so use the **NginxSample** event type.

## Integration development usage
Assuming that you have source code you can build and run the NGINX Integration locally.

* Go to directory of the NGINX Integration and build it
```bash
$ make
```
* The command above will execute tests for the NGINX Integration and build an executable file called `nr-nginx` in `bin` directory.
```bash
$ ./bin/nr-nginx
```
* If you want to know more about usage of `./nr-nginx` check
```bash
$ ./bin/nr-nginx -help
```

For managing external dependencies [govendor tool](https://github.com/kardianos/govendor) is used. It is required to lock all external dependencies to specific version (if possible) into vendor directory.
