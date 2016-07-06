# kodi_exporter

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/nlamirault%2Fkodi_exporter.svg)](https://badge.fury.io/gh/nlamirault%2Fkodi_exporter)

Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/kodi_exporter/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/kodi_exporter/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/kodi_exporter/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/kodi_exporter/tree/develop)

## Installation

You can download the binaries :

* Architecture i386 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_linux_386) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_openbsd_386) / [windows](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/kodi_exporter-0.1.0_netbsd_arm) ]


## Usage

Test your Kodi API using bash:

    $ curl -X POST -H "Content-Type: application/json" \
            -d '{"jsonrpc":"2.0","method":"GUI.ShowNotification","params":{"title":"add you title here","message":"add your message here"},"id":1}'  http://192.168.1.10:8080/jsonrpc
    {"id":1,"jsonrpc":"2.0","result":"OK"}

Then launch the Prometheus exporter :

    $ kodi_exporter -log.level=debug -kodi.server http://192.168.1.10:8080/jsonrpc

## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
