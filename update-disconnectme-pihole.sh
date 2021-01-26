#!/bin/bash
#
# Copy *only* this script to your system's permanent storage and edit accordingly.
#

# Clone disconnectme-pihole repo if doesn't exist on RAM disk
if ! [ -d "/dev/shm/disconnectme-pihole" ]; then
    git clone git@github.com:erkexzcx/disconnectme-pihole.git /dev/shm/disconnectme-pihole
fi

# Navigate to RAM disk
cd /dev/shm/disconnectme-pihole

# Update files
go run update.go
if [[ $? -ne 0 ]]; then
    echo "update.go script failed"
    exit 1
fi

# Update README.md 
echo \#\ Disconnectme-pihole > README.md
echo Disconnectme\ JSON\ files\ converted\ to\ Pi-Hole\ compatible\ blocklists. >> README.md
echo  >> README.md
echo \#\ Usage >> README.md
echo  >> README.md
echo \*\*\*DISCLAIMER\*\*:\ Both\ entities.txt\ and\ services.txt\ seem\ to\ contain\ large\ amount\ of\ false\ positives.\ Use\ with\ caution.\* >> README.md
echo  >> README.md
echo There\ are\ generally\ 2\ files\ in\ this\ repository: >> README.md
echo \`\`\` >> README.md
echo https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/entities.txt >> README.md
echo https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/services.txt >> README.md
echo \`\`\` >> README.md
echo >> README.md
echo If\ you\ want\ to\ customize\ \*\*services.txt\*\*\ file,\ use\ below\ blocklists\ instead\ \(\*it\'s\ literally\ the\ same\ services.txt\ file,\ but\ splitted\ into\ categories\*\): >> README.md
echo \`\`\` >> README.md
for f in services_*.txt; do
    echo "https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/"$(basename "$f") >> README.md
done
echo \`\`\` >> README.md
echo  >> README.md
echo \*\*Source\*\*:\ https://github.com/disconnectme/disconnect-tracking-protection >> README.md

# Exit if no changes
if ! [[ `git status --porcelain` ]]; then
    exit 0
fi

# Push changes
git add -A .
git commit -m "update $(date +%Y-%m-%d)"
git push
