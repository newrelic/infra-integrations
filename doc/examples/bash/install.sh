#!/bin/sh

CUSTOM_INTEGRATIONS=/var/db/newrelic-infra/custom-integrations
ETC_CONFIG=/etc/newrelic-infra/integrations.d

# Create these directories if they don't exist
mkdir ${CUSTOM_INTEGRATIONS}/bin
mkdir ${CUSTOM_INTEGRATIONS}/template

# Copy the config file to the integrations.d directory
cp ./config/*.yaml ${ETC_CONFIG}
echo "All config files copied."

# Copy the configuration file to the custom-integrations directory
cp ./definition/dir-stats-def.yaml ${CUSTOM_INTEGRATIONS}
echo "Definition YAML file copied."

# Copy the template file to a template directory
cp ./template/*.json ${CUSTOM_INTEGRATIONS}/template/
echo "Template JSON files copied."

# Copy the shell script the bin directory
cp ./bin/dir-stats.sh ${CUSTOM_INTEGRATIONS}/bin/
echo "Bash script copied."

# Make sure the executable can be executed
chmod 755 ${CUSTOM_INTEGRATIONS}/bin/dir-stats.sh
echo "Bash script made into an executable."

# Restart the newrelic-infra agent
service newrelic-infra restart
echo "Infrastructure service restarted"