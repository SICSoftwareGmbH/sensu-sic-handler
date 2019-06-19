[![GitHub release](https://img.shields.io/github/tag/SICSoftwareGmbH/sensu-sic-handler.svg?label=latest)](https://github.com/SICSoftwareGmbH/sensu-sic-handler/releases)
[![Travis](https://img.shields.io/travis/SICSoftwareGmbH/sensu-sic-handler/master.svg)](https://travis-ci.org/SICSoftwareGmbH/sensu-sic-handler)
[![License](https://img.shields.io/github/license/SICSoftwareGmbH/sensu-sic-handler.svg)](./LICENSE)

# Sensu Go SIC Handler

The Sensu SIC handler is a [Sensu Event Handler][1] that is used by SIC! Software for dispatching events based on entity annotations.

## Installation

Download the latest version of the sensu-sic-handler from [releases][2],
or create an executable script from this source.

From the local path of the sensu-sic-handler repository:
```
go build -o /usr/local/bin/sensu-sic-handler main.go
```

## Configuration

### Example:

```yaml
etcd-endpoints: http://etcd:2379

redmine-url: https://redmine.example.com
redmine-token: foobar

annotation-prefix: com.example
smtp-address: smtp.example.com:25
mail-from: sensu@example.com
slack-webhook-url: https://hooks.slack.com/services/foo/bar/foobar
slack-username: sensu
slack-icon-url: http://s3-us-west-2.amazonaws.com/sensuapp.org/sensu.png
xmpp-server: jabber.example.com
xmpp-username: sensu@jabber.example.com
xmpp-password: foobar
```

## Usage examples

      $ sensu-sic-handler redmine import -c sic-handler.yml
      $ sensu-sic-handler event -c sic-handler.yml --outputs='mail'
      $ sensu-sic-handler event -c sic-handler.yml --outputs='slack'
      $ sensu-sic-handler event -c sic-handler.yml --outputs='xmpp'

[1]: https://docs.sensu.io/sensu-go/5.0/reference/handlers/#how-do-sensu-handlers-work
[2]: https://github.com/SICSoftwareGmbH/sensu-sic-handler/releases
