#!/bin/bash

# 0. Input validation:

if [ "$(uname)" != "Linux" ]; then
    echo "$0 is for Linux"
    exit 1
fi

if [ "${DIR_NAMES}" == "" ]; then
    echo "DIR_NAMES cannot be empty"
    exit 1
fi


# 1. Iterate on entities:

jsonEntities=""
entityCount=0

IFS=',' read -ra directories <<< "${DIR_NAMES}"
for dirName in "${directories[@]}"; do


    # 2. Data sampling:

    # Get the count of files in the directory
    # find gets the list, wc counts the lines, tr trims whitespace
    fileCount=`find "${dirName}" -type f | wc -l | tr -d ' '`

    # Get the directory size, (note that cut breaks off the first field)
    dirSize=`du "${dirName}" -b -s | cut -f1`

    subfolders=($(find ${dirName}  -maxdepth 1 -type d ))


    # 3. Serialize to JSON:

    # Entity template evaluation
    jsonEntity=`cat ./template/entity.json`

    jsonSubdirs="{"
    for subfolder in "${subfolders[@]}"; do
        separator=""
        if [ "${jsonSubdirs}" != "{" ]; then
            separator=","
        fi
        jsonSubdirs="$jsonSubdirs$separator \"$subfolder\": true"
    done
    jsonSubdirs="$jsonSubdirs }"

    # Replace the values in the JSON
    # The @ in the sed command is a delimiter
    jsonEntity=`echo ${jsonEntity} | sed -e "s@DIR_NAME@${dirName}@"`
    jsonEntity=`echo ${jsonEntity} | sed -e "s@FILE_COUNT@${fileCount}@"`
    jsonEntity=`echo ${jsonEntity} | sed -e "s@DIR_SIZE@${dirSize}@"`
    jsonEntity=`echo ${jsonEntity} | sed -e "s@SUBFOLDERS@${jsonSubdirs}@"`

    separator=""
    entityCount=${entityCount}+1
    if (( ${entityCount} > 1 )); then
        separator=","
    fi

    jsonEntities="${jsonEntities}${separator}${jsonEntity}"
done


# 4. Integration template evaluation:

jsonIntegration=`cat ./template/integration.json`

jsonIntegration=`echo ${jsonIntegration} | sed -e "s@ENTITIES@${jsonEntities}@"`

# Remove whitespaces
jsonIntegration=`printf "${jsonIntegration}" | tr -d [:space:]`


# 5. Output result:

echo "${jsonIntegration}"
