# kodi_exporter

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/nlamirault%2Fkodi_exporter.svg)](https://badge.fury.io/gh/nlamirault%2Fkodi_exporter)

Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/kodi_exporter/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/kodi_exporter/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/kodi_exporter/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/kodi_exporter/tree/develop)

## Installation

You can download the binaries :

* Architecture i386 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_linux_386) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_openbsd_386) / [windows](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.2.0_netbsd_arm) ]


## Usage

First, test your Kodi API using bash:

    $ curl -X POST -H "Content-Type: application/json" \
            -d '{"jsonrpc":"2.0","method":"GUI.ShowNotification","params":{"title":"add you title here","message":"add your message here"},"id":1}'  http://192.168.1.10:8080/jsonrpc
    {"id":1,"jsonrpc":"2.0","result":"OK"}



Then Launch the Prometheus exporter :

    $ kodi_exporter -log.level=debug -kodi.server 192.168.1.10 -kodi.port 8080


## Debug

You could try your Kodi API :

* Artists:

        $ curl': curl -X POST -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"AudioLibrary.GetArtists","params":{},"id":1}'  http://10.10.10.10:8080/jsonrpc

* Albums:

        $ curl': curl -X POST -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"AudioLibrary.GetAlbums","params":{},"id":1}'  http://10.10.10.10:8080/jsonrpc

* Songs:

        $ curl': curl -X POST -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"AudioLibrary.GetSongs","params":{},"id":1}'  http://10.10.10.10:8080/jsonrpc

## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test


## Local Deployment

* Launch Prometheus using the configuration file in this repository:

        $ prometheus -config.file=prometheus.yml

* Launch exporter:

        $ kodi_exporter -log.level=debug -kodi.server 192.168.1.10 -kodi.port 8080

* Check that Prometheus find the exporter on `http://localhost:9090/targets`


## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).

Some Kodi API calls :


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
