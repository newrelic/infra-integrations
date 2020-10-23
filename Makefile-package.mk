DEP_APACHE         = "nri-apache"
DEP_CASSANDRA      = "nri-cassandra"
DEP_MYSQL          = "nri-mysql"
DEP_NGINX          = "nri-nginx"
DEP_REDIS          = "nri-redis"
PACKAGE_TYPES     ?= deb rpm
PROJECT_NAME       = newrelic-infra-integrations
BINS_PREFIX        = nr
PACKAGES_DIR       = $(TARGET_DIR)/packages
VERSION           ?= 0.0.0
RELEASE           ?= dev
LICENSE            = "https://newrelic.com/terms (also see LICENSE.txt installed with this package)"
VENDOR             = "New Relic, Inc."
PACKAGER           = "New Relic Infrastructure Team <infrastructure-eng@newrelic.com>"
PACKAGE_URL        = "https://www.newrelic.com/infrastructure"
SUMMARY            = "New Relic Infrastructure Integrations"

define DESCRIPTION
New Relic Infrastructure Integrations extend the core New Relic\n\
Infrastructure agent\'s capabilities to allow you to collect metric and\n\
live state data from your infrastructure components such as MySQL,\n\
NGINX and Cassandra.\n\
\n\
This is a meta-package that specifies dependencies to the\n\
aforementioned integrations, and installs all of them. You may prefer\n\
to install them individually, selecting only those that you need.
endef

FPM_COMMON_OPTIONS = --verbose -C tmp -s dir -n $(PROJECT_NAME) -v $(VERSION) --iteration $(RELEASE)\
 --prefix "" --license $(LICENSE) --vendor $(VENDOR) -m $(PACKAGER) --url $(PACKAGE_URL)\
 --description "$$(printf $(DESCRIPTION))"

FPM_DEB_OPTIONS    = -t deb -p $(PACKAGES_DIR)/deb/ --replaces "newrelic-infra-integrations-beta (<= 1.0.0)" 
FPM_RPM_OPTIONS    = -t rpm -p $(PACKAGES_DIR)/rpm/ --epoch 0 --rpm-summary $(SUMMARY) --replaces "newrelic-infra-integrations-beta <= 1.0.0"

package: $(PACKAGE_TYPES)

deb:
	@echo "=== Main === [ deb ]: building DEB package..."
	@mkdir -p $(PACKAGES_DIR)/deb
	@fpm $(FPM_COMMON_OPTIONS) $(FPM_DEB_OPTIONS) .

rpm:
	@echo "=== Main === [ rpm ]: building RPM package..."
	@mkdir -p $(PACKAGES_DIR)/rpm
	@fpm $(FPM_COMMON_OPTIONS) $(FPM_RPM_OPTIONS) .

.PHONY: package $(PACKAGE_TYPES)
