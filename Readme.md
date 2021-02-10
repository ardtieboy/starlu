# Starlu

Small golang CLI app to resize/crop/edit images for Pragalicious

Made with Cobra: https://github.com/spf13/cobra/blob/master/cobra/README.md

## Building

    go build

    GOOS=windows go build

## Run tests

    cd imageprocessing
    go test -v

## Installation

### Mac/Linux

1. Download the binary `starlu`
2. Add the binary to your path
3. Test the application `starlu --version`
4. Run the binary `starlu teaser test.jpg`

### Windows

1. Download the binary `starlu.exe`
2. [Optional] add `starlu.exe` to your path
3. Open the command line
4. Go to the folder where the binary is situated (not needed if you added it to your path). For example: `cd \Users\joeri\Downloads`
5. Test the application `starlu.exe --version`
6. Run the exe from there `starlu.exe teaser test.jpg`