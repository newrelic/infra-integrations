#!/bin/sh

CUSTOM_INTEGRATIONS=/var/db/newrelic-infra/custom-integrations
ETC_CONFIG=/etc/newrelic-infra/integrations.d

# Create these directories if they don't exist
mkdir ${CUSTOM_INTEGRATIONS}/bin

# Copy the config file to the integrations.d directory
cp ./config/*.yaml ${ETC_CONFIG}
echo "All config files copied."

# Copy the configuration file to the custom-integrations directory
cp ./definition/dir-stats-def.yaml ${CUSTOM_INTEGRATIONS}
echo "Definition YAML file copied."

# Copy the shell script the bin directory
cp ./bin/dir-stats.py ${CUSTOM_INTEGRATIONS}/bin/
echo "Python script copied."

# Make sure the executable can be executed
chmod 755 ${CUSTOM_INTEGRATIONS}/bin/dir-stats.py
echo "Python script made into an executable."

# Restart the newrelic-infra agent
service newrelic-infra restart
echo "Infrastructure service restarted"