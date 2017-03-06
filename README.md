# New Relic Infrastructure Integrations

New Relic Infrastructure, provided by New Relic, Inc (http://www.newrelic.com),
offers flexible, dynamic server monitoring. This package contains the set of
official integrations supported by New Relic, built to provide the essential
metrics and inventory for monitoring the services. That data will be findable and
usable in New Relic Infrastructure and in New Relic Insights. You can find more
information on how to access and visualize that data on [our docs site](https://docs.newrelic.com/docs/find-use-infrastructure-integration-data).

 The New Relic Infrastructure Integrations are hosted on [github](https://github.com/newrelic/infra-integrations),
 and the Integrations Golang SDK is hosted on [github](https://github.com/newrelic/infra-integrations-sdk).


## Compatibility and requirements

Up-to-date [our docs site](https://docs.newrelic.com/docs/compatibility-requirements-infrastructure-integration-sdk).


## Contributing Code

We welcome code contributions (in the form of pull requests) from our user
community. Before submitting a pull request please review [these guidelines](https://github.com/newrelic/infra-integrations/blob/master/CONTRIBUTING.md).

Following these helps us efficiently review and incorporate your contribution
and avoid breaking your code with future changes to the agent.


## Official Integrations

Currently the set of official integrations includes three services, to find more
information about each one of them, how to configure and use them with New Relic,
refer to their documentation pages or the README included by each of them:

* Cassandra:
  - [docs site](https://docs.newrelic.com/docs/cassandra-integration-new-relic-infrastructure)
  - [README.md](integrations/cassandra/README.md)

* MySQL:
  - [docs site](https://docs.newrelic.com/docs/mysql-integration-new-relic-infrastructure)
  - [README.md](integrations/mysql/README.md)

* NGINX:
  - [docs site](https://docs.newrelic.com/docs/nginx-integration-new-relic-infrastructure)
  - [README.md](integrations/nginx/README.md)


## Custom Integrations

To extend your monitoring solution with custom metrics, we offer the Integrations
Golang SDK which can be found on [github](https://github.com/newrelic/infra-integrations-sdk).

Refer to [our docs site](https://docs.newrelic.com/docs/infrastructure/integrations-sdk/get-started/intro-infrastructure-integrations-sdk)
to get help on how to build your custom integrations.


## Build and Test

The project includes a set of makefiles to help with building and running the unit
tests for the official integrations.

Running the unit tests for all the integrations or a specific set of them is as
easy as calling the following commands from the project root:

```bash
$ make test
$ INTEGRATIONS="nginx mysql" make test
```

Similarly, to build all the integrations or a specific set of them can be done
by calling the following commands from the project root:

```bash
$ make build
$ INTEGRATIONS="nginx mysql" make build
```

## Support

You can find more detailed documentation [on our website](http://newrelic.com/docs),
and specifically in the [Infrastructure category](https://docs.newrelic.com/docs/infrastructure).

If you can't find what you're looking for there, reach out to us on our [support
site](http://support.newrelic.com/) or our [community forum](http://forum.newrelic.com)
and we'll be happy to help you.

Find a bug? Contact us via [support.newrelic.com](http://support.newrelic.com/),
or email support@newrelic.com.

New Relic, Inc.
