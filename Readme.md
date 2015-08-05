# gather

gather is a library that collects info on how user is using their VM. This libaray contains the abuse and analytics scripts which is ultimately distributed via a binary for others to run.

## Scripts

The scripts themselves are binary encoded into a Go binary using `go-bindata`. The main reason for this is to obuscate the scripts from the user. Some of the scripts check for abuse, if these are accesible in clear text, the user can easily circumvent it. The scripts need to be in bash and return output in format:

```
{
  "name"  : "<metric name>",
  "type"  : "<type of value>",
  "value" : <actual value>
}
```

Due to encoding, we need to do a trick to avoid duplication. `gatherers/common` contains the reusable bash functions; any function added to this file will be appened in memory to the scripts file before the combined string is executed by Go using `bash -C` command. Any file in a folder will be considered as runnable.

Even though the scripts are binary encoded and the program shared as a binary, the binary itself might contain enough info to tip people about the scripts. This is why the name of the scripts are only 3 lettes in length.

There are currently two types of scripts: abuse in `gatherers/ab/` and analytics in `gatherers/an/`. Which script to run is determined via `GATHER` env variable. The seperation is so that abuse scripts are run more often than analytics scripts.

## Building

`./build.sh` will cross compile the `gather` binary. You'll need cross compilation enabled in Go (GOOS=linux GOARCH=386). It'll also upload the script to S3.

## Running

    # to run analytics scripts
    ./run.sh analytics`

    # to run abuse scripts
    ./run.sh abuse`

## Tests

`go test`
