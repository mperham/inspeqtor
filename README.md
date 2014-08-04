# inspeqtor

Next generation service monitoring.  Inspired by a decade of
using monit but a complete rethink of what's necessary for modern
server-side applications.

What it does:

* Monitor upstart-, runit-, systemd- or launchctl-managed services
* Monitor process memory and CPU usage
* Monitor host CPU, swap and disk usage
* Notify if processes disappear or change PID
* Notify if processes or host goes over defined RAM or CPU utilization
* Email notification
* As developer friendly as possible:
  - Test configuration
  - Test notifications

What it doesn't:

* support PID files, which are racy and error-prone
* monitor arbitrary processes, services must be init-managed
* know how to start/stop services.  Defers to your OS's init system
* have *any* runtime or 3rd party dependencies at all, not even libc.

The default monitoring rules out of the box perform basic health checks:

* / partition is > 90% full
* Swap is more than 20% utilized
* CPU(user) is > 90% for more than 4 cycles
* load(5) is > 10
* load(1) is > 20


## Platforms

inspeqtor's platform target is Linux 3.0+.  Other platforms (OSX,
FreeBSD) aren't as well-supported but I welcome help to improve it.
Non-Unix platforms, e.g. Windows, aren't supported at this time, mostly
because I don't have a Windows machine and so have no way of testing on it.


## Requirements

inspeqtor has no third-party dependencies.  It uses about 5-10MB of RAM at runtime.


## License

Licensed under GPLv3.


# Author

inspeqtor is written by [Mike Perham](http://twitter.com/mperham) of [Contributed Systems](http://contribsys.com).  We build awesome open source-based products.

We also develop [Sidekiq](http://sidekiq.org) and sell [Sidekiq Pro](http://sidekiq.org/pro), the best Ruby background job processing system.
