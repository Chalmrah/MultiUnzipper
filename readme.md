# MultiUnzipper

MultiUnzipper is a Golang program designed to calculate or extract all compressed files in multiple 7zip archives.

## Features

- List the contents of a folder.
- Extract archives to a specified destination.

## Installation

1. Download the latest file from the releases page.

## Usage

### Flags

- `-folder` (string): The path to the folder that contains the files to process. (Required)
- `-extract` (bool): Flag to indicate whether to extract files from archives. (Default: `false`)
- `-destination` (string): The path where extracted files should be stored. (Required if `-extract` is set to true)
- `-version` (bool): Show the version information of MultiUnzipper. (Default: `false`)

### Commands

#### Get uncompressed size of all archives in the folder

To get the compressed and uncompressed size of the folder (without extracting), use the `-folder` flag.

```bash
./multiunzipper -folder /path/to/folder
```

This will output the sum of compressed and uncompressed sizes in the provided folder.

#### Extract Files

To extract all files from a folder to a specific destination, use the `-extract` and `-destination` flags:

```bash
./multiunzipper -folder /path/to/folder -extract -destination /path/to/destination
```