# clamav-proyect

## Prerequisites

- golang 1.18.1 or latest
- docker

## Installation

```bash
docker pull hub.docker.com/ajilaag/clamav-rest:latest
```

## Quick Start

Run clamav-rest docker image:

```bash
docker run -p 9000:9000 -p 9443:9443 -itd --name clamav-rest ajilaag/clamav-rest
```

## Test

```bash
curl -X POST http://localhost:8080/api/scan -H "Content-Type: application/x-www-form-urlencoded" -d "file=/path"
```

## Result

- clean file = no KNOWN infections
```{ "Status": "OK", "Description": "" }```

- INFECTED
```{ "Status": "FOUND", "Description": "Eicar-Test-Signature" }```

- unable to parse file
```{ "Status": "PARSER-ERROR", "Description": "" }```

- ClamAV returned general error for file
```{ "Status": "ERROR", "Description": "" }```

- unknown request
```{ "Status": "UNKNOWN-REQUEST", "Description": "" }```

## Clamav Container Environment Variables

Below is the complete list of available options that can be used to customize your installation.

| Parameter | Description |
|-----------|-------------|
| `MAX_SCAN_SIZE` | Amount of data scanned for each file - Default `100M` |
| `MAX_FILE_SIZE` | Don't scan files larger than this size - Default `25M` |
| `MAX_RECURSION` | How many nested archives to scan - Default `16` |
| `MAX_FILES` | Number of files to scan withn archive - Default `10000` |
| `MAX_EMBEDDEDPE` | Maximum file size for embedded PE - Default `10M` |
| `MAX_HTMLNORMALIZE` | Maximum size of HTML to normalize - Default `10M` |
| `MAX_HTMLNOTAGS` | Maximum size of Normlized HTML File to scan- Default `2M` |
| `MAX_SCRIPTNORMALIZE` | Maximum size of a Script to normalize - Default `5M` |
| `MAX_ZIPTYPERCG` | Maximum size of ZIP to reanalyze type recognition - Default `1M` |
| `MAX_PARTITIONS` | How many partitions per Raw disk to scan - Default `50` |
| `MAX_ICONSPE` | How many Icons in PE to scan - Default `100` |
| `PCRE_MATCHLIMIT` | Maximum PCRE Match Calls - Default `100000` |
| `PCRE_RECMATCHLIMIT` | Maximum Recursive Match Calls to PCRE - Default `2000` |
| `SIGNATURE_CHECKS` | Check times per day for a new database signature. Must be between 1 and 50. - Default `2` |

## See also

- <https://hub.docker.com/r/ajilaag/clamav-rest>
