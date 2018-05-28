# Dir Stats On Host Integration

> Based on the previous work https://github.com/kenahrens/nri-bash
> This adds support for remote-entities.

In this sample each directory will be mapped to a different entity.
Each entity will report directory-size and file-count metrics.

## Configuration

This script expects you to supply 1 variable:
* DIR_NAMES - list of directories from which to collect extra stats.

In order to run this you do the following
* Copy `config/dir-stats-config.COPYME` to `config/dir-stats-config.yaml`
* Edit your YAML file, add an instance for each directory you want to track

```yaml
integration_name: com.myorg.dir-stats

instances:
  - name: Instance 1
    command: dir-stats
    arguments:
      dir_names:
        - /tmp
        - /var
```

Now that you have your own config file you run the installer (with sudo). It will put all the files in the correct places and will restart the newrelic-infra agent. After you run it for the first time you'll see those `File exists` warnings which are no problem.

```
dir-stats$ sudo sh install.sh 
mkdir: cannot create directory ‘/var/db/newrelic-infra/custom-integrations/bin’: File exists
mkdir: cannot create directory ‘/var/db/newrelic-infra/custom-integrations/template’: File exists
All config files copied.
dir-stats-definition YAML file copied.
dir-stats-template JSON files copied.
dir-stats script copied.
dir-stats made into an executable.
Infrastructure service restarted
```

## Viewing Data

You will see new attributes in  `DirectorySample` called `dirName`, `dirSize` and `fileCount`.

> dirNames should be available on this query once the agent supports entity-names for remote-entities.

Example query used:
```
SELECT dirSize, fileCount FROM DirectorySample SINCE 15 minutes ago
```
