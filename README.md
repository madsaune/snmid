# ServiceNow MID Server Downloader

## Installation

```bash
go install github.com/madsaune/snmid
```

## Quick Start

```bash
Usage: snmid <buildstamp> [flags]

  --url-only    only print the download url
  -p            target platform (windows, linux)
  -t            installer filetype (msi, zip)
  -o            output file
```

To find the buildstamp for your ServiceNow instance you must go to `https://<instance_name>.service-now.com/stats.do` and look for `MID buildstamp`. Depending on how your instance is configured, this may require that your are logged in.

```bash
# download windows installer
snmid sandiego-12-22-2021__patch9a-hotfix1-01-31-2023_02-01-2023_1625 -o ./installer.msi

# download windows installer as zip
snmid sandiego-12-22-2021__patch9a-hotfix1-01-31-2023_02-01-2023_1625 -t zip -o ./mid_server.zip

# download linux installer
snmid sandiego-12-22-2021__patch9a-hotfix1-01-31-2023_02-01-2023_1625 -p linux -t zip -o ./mid_server.zip

# get the only the download url
snmid sandiego-12-22-2021__patch9a-hotfix1-01-31-2023_02-01-2023_1625 --url-only
```
