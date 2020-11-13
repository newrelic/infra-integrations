[![Archived header](https://github.com/newrelic/open-source-office/raw/master/examples/categories/images/Archived.png)](https://github.com/newrelic/open-source-office/blob/master/examples/categories/index.md#archived)

# New Relic Infrastructure Integrations

**DEPRECATION NOTICE AND UPGRADE INSTRUCTIONS**

This package was created at a time when New Relic Infrastructure was just
starting with a reduced number of developers. The decision to use a single 
package for installing all our integrations, and a mono-repository to host the
code, was something that made sense at the time; however, now the circumstances
are very different, we have significantly increased the number of integrations 
and the people working on them, which made us switch to an approach based on
individual repositories and packages per integration, this allows us more 
control, faster delivery, it's easier to maintain and causes less errors.

Since version `2.0.0` this package will be empty and no integrations will be 
installed. 

If you're still using this package, we advise you to follow the [NewRelic docs
on how to uninstall integrations][1]; basically:

- Debian/Ubuntu:
	1. `sudo apt-get remove newrelic-infra-integrations`
	1. `sudo apt-get autoremove`
- CentOS/RHEL/Amazon:
	1. `sudo yum remove newrelic-infra-integrations`
	1. `sudo yum autoremove`
- SLES:
	1. `sudo zypper -n remove newrelic-infra-integrations --clean-deps`

Then, follow the [NewRelic docs on how to install integrations][2] to install
the individual integrations you need, keep in mind that the change should be
transparent, your configuration files won't be deleted by uninstalling
this package and will be valid for the individual packages.

**END OF DEPRECATION NOTICE AND UPGRADE INSTRUCTIONS**

New Relic Infrastructure, provided by New Relic, Inc (http://www.newrelic.com),
offers flexible, dynamic server monitoring. This package contains the set of
official integrations supported by New Relic, built to provide the essential
metrics and inventory for monitoring the services. That data will be findable and
usable in New Relic Infrastructure and in New Relic Insights. You can find more
information on how to access and visualize that data on [our docs site](https://docs.newrelic.com/docs/find-use-infrastructure-integration-data).

This repository serves as an aggregator where you can find each of our
integrations as submodules under the `integrations/` directory, you'll find
more information about them in the following section. Previously this was a
mono-repo that hosted the source code of the integrations, but as the number of
integrations grew we move them to their own repositories.

The Integrations Golang SDK is hosted on [github](https://github.com/newrelic/infra-integrations-sdk).

## Compatibility and requirements

Up-to-date [our docs site](https://docs.newrelic.com/docs/compatibility-requirements-infrastructure-integration-sdk).

## Official Integrations

To find more information about our integrations, how to configure and use them
with New Relic, refer to their documentation pages or the README included by
each of them:

| Integration 	| Documentation 																		| Readme  |
| ------------- |:-------------:																		| -----:|
| Cassandra 	| [doc](https://docs.newrelic.com/docs/cassandra-integration-new-relic-infrastructure) 	| [readme.md](https://github.com/newrelic/nri-cassandra/README.md) |
| MySQL 		| [doc](https://docs.newrelic.com/docs/mysql-integration-new-relic-infrastructure) 		| [readme.md](https://github.com/newrelic/nri-mysql/README.md) |
| NGINX 		| [doc](https://docs.newrelic.com/docs/nginx-integration-new-relic-infrastructure) 		| [readme.md](https://github.com/newrelic/nri-nginx/README.md) |
| Redis 		| [doc](https://docs.newrelic.com/docs/redis-integration-new-relic-infrastructure) 		| [readme.md](https://github.com/newrelic/nri-redis/README.md) |
| Apache 		| [doc](https://docs.newrelic.com/docs/apache-integration-new-relic-infrastructure) 	| [readme.md](https://github.com/newrelic/nri-apache/README.md) |

To download this repository with the integrations as submodules run:

```
git clone --recurse-submodules https://github.com/newrelic/infra-integrations.git
```

## Contributing Code

We welcome code contributions (in the form of pull requests) from our user
community. Before submitting a pull request please review [these guidelines](https://github.com/newrelic/infra-integrations/blob/master/CONTRIBUTING.md).

Following these helps us efficiently review and incorporate your contribution
and avoid breaking your code with future changes to the agent.

## Custom Integrations

To extend your monitoring solution with custom metrics, we offer the Integrations
Golang SDK which can be found on [github](https://github.com/newrelic/infra-integrations-sdk).

Refer to [our docs site](https://docs.newrelic.com/docs/infrastructure/integrations-sdk/get-started/intro-infrastructure-integrations-sdk)
to get help on how to build your custom integrations.

We also provided a [command line builder tool](https://github.com/newrelic/nr-integrations-builder)
for creating and scaffolding new integration in Golang. This will generate the
project structure and include the dependencies needed (including the
sdk) for creating your custom integrations.

## Support

You can find more detailed documentation [on our website](http://newrelic.com/docs),
and specifically in the [Infrastructure category](https://docs.newrelic.com/docs/infrastructure).

If you can't find what you're looking for there, reach out to us on our [support
site](http://support.newrelic.com/) or our [community forum](http://forum.newrelic.com)
and we'll be happy to help you.

Find a bug? Contact us via [support.newrelic.com](http://support.newrelic.com/),
or email support@newrelic.com.

New Relic, Inc.

[1]: https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure/installation/uninstall-infrastructure-agent#uninstall-integrations
[2]: https://docs.newrelic.com/docs/integrations/host-integrations/installation/install-host-integrations-built-new-relic
