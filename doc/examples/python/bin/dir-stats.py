#!/usr/bin/env python
import argparse
import random
import json
import logging
import os


# Protocol v2 integration with remote entities
class Integration:
    def __init__(self, name, integration_version):
        self.name = name
        self.protocol_version = "2"
        self.integration_version = integration_version
        self.data = []

    def add_entity(self, type, name):
        e = Entity(type, name)
        self.data.append(e)
        return e

    def to_json(self, pretty):
        indentation = None if pretty is False else 4

        return json.dumps(self, default=lambda o: o.__dict__,
                          sort_keys=pretty, indent=indentation)


class Entity:
    def __init__(self, type, name):
        self.entity = {
            'type': type,
            'name': name,
        }
        self.inventory = {}
        self.metrics = []
        self.events = []

    def add_inventory(self, item, key, value):
        self.inventory.setdefault(item, {})[key] = value

    def add_metric(self, metric_dict):
        self.metrics.append(metric_dict)


# Setup the integration's command line parameters
def parse_arguments():
    parser = argparse.ArgumentParser()
    parser.add_argument('-v', default=False, dest='verbose', action='store_true',
                        help='Print more information to logs')
    parser.add_argument('-p', default=False, dest='pretty', action='store_true',
                        help='Print pretty formatted JSON')

    args, unknown = parser.parse_known_args()
    return args


# Setup logging, redirect logs to stderr and configure the log level.
def get_logger():
    logger = logging.getLogger("infra")
    logger.addHandler(logging.StreamHandler())
    if args.verbose:
        logger.setLevel(logging.DEBUG)
    else:
        logger.setLevel(logging.INFO)

    return logger


def folder_size(start_path='.'):
    total_size = 0
    seen = {}
    for dirpath, dirnames, filenames in os.walk(start_path):
        for f in filenames:
            fp = os.path.join(dirpath, f)
            try:
                stat = os.stat(fp)
            except OSError:
                continue

            try:
                seen[stat.st_ino]
            except KeyError:
                seen[stat.st_ino] = True
            else:
                continue

            total_size += stat.st_size

    return total_size


if __name__ == "__main__":
    folders = os.getenv("DIR_NAMES")
    if folders is None:
        print("Environment variable DIR_NAMES is required with format: \"/tmp,/var\"")
        exit(1)

    args = parse_arguments()

    logger = get_logger()

    # Integration
    i = Integration("com.myorg.dir-stats", "1.0.0")

    for folder in folders.split(","):
        # Entities
        e = i.add_entity("directory", folder)

        # Inventory
        folderContents = os.listdir(folder)
        for subFolder in folderContents:
            if os.path.isdir(os.path.join(folder, subFolder)):
                e.add_inventory("subfolders", subFolder, True)
                logger.debug("Added inventory item on subfolders: %s", subFolder)

        # Metrics
        metric = {
            "event_type": "DirectorySample",
            "fileCount": len(folderContents),
            "dirSize": folder_size(folder),
        }
        e.add_metric(metric)
        logger.debug("Added metric %s", metric)

    # Output
    print i.to_json(args.pretty)
