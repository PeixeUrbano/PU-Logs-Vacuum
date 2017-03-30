<h1 align="center">
  <br>
  Logs Vacuum Cleaner
  <br>
</h1>

<h4 align="center">A small and configurable logs cleaner tool.</h4>

### Features

**Logs Vacuum Cleaner** is a small and flexible (i.e. configurable) tool created to remove old log files, allowing
you to keep them compressed in a single file and/or permanently remove them.

Generally, the **Logs Vacuum Cleaner** is applicable whenever:

* Your application(s) generates multiple log files and there isn't any post process to deal with them
* You want to be able to compress a set of log files using a compression algorithm of choice.
* You have more than one directory storing logs and you want a centralized way to purge them.
* You want to discard old, non-compressed files.

### Configuration

Getting right to the point, configurations are stored at ./config/vacuums.json dir. 
This is a complete configuration example:

```json
{
    "vacuums" : [
        {
            "path" : "/some/log/path/",
            "filesPrefix" : "server*",
            "filesSuffix" : "*.log",
            "removeLogs" : true,
            "compressor" : "tar.gz",
            "outputPath" : "/some/backup/path/",
            "outputName" : "logs.2006-01.tag.gz",
            "updateOutput" : true
        }
    ]
}
```

*Where:*

#### - path
The Base path where the logs are stored. *Required.*

#### - filesPrefix
Files prefix. Accepts wildcard \*. Under \*nix systems you can
specify more than one pattern using braces. Ex: 
{log-\*,access-\*,error-\*} *Not required.*

#### - filesSuffix
Files suffix. Accepts wildcard \*. Under \*nix systems you can
specify more than one pattern using braces. Ex: 
{\*.log,\*.txt,\*.output} *Not required.*

#### - removeLogs
Should the Vacuum remove found logs after merging and compress them? *Not required.*

#### - compressor
Compression algorithm. Only tar.gz is supported at the moment. *Required.*

#### - outputPath
Output path where the compressed file will be stored. *Required.*

#### - outputName
Output file name where the compressed file will be stored. Accepts [Go Time format parameters](https://golang.org/pkg/time/#pkg-examples). *Required.*

#### - updateOutput
If the vacuum finds a previously generated output, should it just update the output file (by appending new log files) ? *Not Required, default false.*

### Develop using Docker

To build a docker image, you should just simply type: 

`$ docker build --rm -t peixeurbano/logs-vacuum .`

In order to run it, simply do: 

`$ docker run peixeurbano/logs-vacuum` 
 
### Running tests

Tests were write using [Go testing](https://golang.org/pkg/testing/). In order to run them, just type:

`$ go test` 

### Generating builds

Under *scripts/* dir you can regenerate builds using *generate-builds.sh* script file. Currently, we are exporting LogsVacuum to the following systems/archs:

* linux/386
* linux/amd64
* linux/arm
* macos (darwin)/386
* macos (darwin)/amd64
* windows/386
* windows/amd64

The latest builds can be found at /dist dir.

### License

MIT. Copyright (c) [Peixeu Urbano](http://www.peixeurbano.com.br). 
