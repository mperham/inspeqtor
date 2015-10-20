# Inspeqtor Changelog

Inspeqtor is stable and ready for production usage.
Try it out and let us know how we can improve it for you.

## 1.0.0

- Add tls\_port option for SMTP servers. [exploid, #79]
- Add gometalinter. [#83]
- Show silenced until date in status output. [#101]

## 0.8.1-1

- Alerts could still fire during deploy due to a race condition in the
  Inspeqtor's silence window. [#76]

## 0.8.0-1

- Support for real-time memory monitoring for Go daemons **Pro** [#65]
  ![memory monitor](https://cloud.githubusercontent.com/assets/2911/5670572/fde0b112-9735-11e4-8161-6df283d090bc.png)
- Fix issue resolving Down services with Upstart [#73]
- Add grace period for cron job checks **Pro** [#61]
- Add grace period for deploy window alerts [#69]
- **LOTS** of code cleanup based on **golint**, **go vet** and
  **errcheck** static analysis tools. [#64, #63, #30]

## 0.7.0-2

- Welcome new committer, @sorentwo!
- Add daemon-specific metrics for PostgreSQL! [#26]
- Fix bugs in init.d support, better logging
- Add new `reload` action which will send the HUP signal to a service [sorentwo, #38]
- Support unauthenticated port 25 SMTP [felixbuenemann, #48]
- Add memory:total\_rss metric for the total memory consumed by a tree of
  processes (e.g. useful if child processes are bloating) [#10]
- Prefix all statsd metrics with hostname [#42] **Pro**


## 0.6.0-2

- Fix missing /etc/inspeqtor/host.inq in packaging.


## 0.6.0

- Add [init.d](https://github.com/mperham/inspeqtor/wiki/Initd) support [#14]
- Add new "export" command to emit current metric values as JSON [#2]
- Send metrics to [Statsd](https://github.com/mperham/inspeqtor/wiki/Pro-Statsd) [#2] **Pro**
- Monitor [cron job execution](https://github.com/mperham/inspeqtor/wiki/Pro-Recurring-Jobs) [#5] **Pro**
- Support per-second rates for thresholds [#11]
```bash
  # Use the "/sec" suffix. You can use the k,m,g modifier letters also.
  if mysql:Queries > 1k/sec then alert
  if mysql:Slow_queries > 1/sec then alert
```
- Changed the /etc/inspeqtor layout a bit [#22]
- inspeqtorctl no longer requires sudo for members of the 'adm' group [#27]


## 0.5.0

- Initial Release
- Yay, shipped it!
