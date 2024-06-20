# Disconnectme-pihole
Disconnectme JSON files converted to Pi-Hole compatible blocklists.

# Usage

***DISCLAIMER**: Both entities.txt and services.txt seem to contain large amount of false positives. Use with caution.*

There are generally 2 files in this repository:
```
https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/entities.txt
https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/services.txt
```

If you want to customize **services.txt** file, use below blocklists instead (*it's literally the same services.txt file, but splitted into categories*):
```
https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/services_FingerprintingInvasive.txt
https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/services_Content.txt
https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/services_Anti-fraud.txt
https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/services_Analytics.txt
https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/services_Social.txt
https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/services_FingerprintingGeneral.txt
https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/services_Advertising.txt
https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/services_Cryptomining.txt
https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/services_Email.txt
https://raw.githubusercontent.com/erkexzcx/disconnectme-pihole/master/services_EmailAggressive.txt
```

**Source**: https://github.com/disconnectme/disconnect-tracking-protection