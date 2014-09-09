# inspeqtor

Next generation host and process monitoring for server applications.  Inspired
by a decade of using monit but a complete rethink of what's necessary for modern
applications.

What it does:

* Monitor upstart-, runit-, systemd- or launchd-managed services
* Monitor process memory and CPU usage
* Monitor host CPU, load, swap and disk usage
* Alert if processes disappear or change PID
* Alert if processes or host goes over defined RAM or CPU utilization
* Signal deploy start/stop to silence alerts during deploy

What it doesn't:

* monitor arbitrary processes, services must be init-managed
* have *any* runtime or 3rd party dependencies at all, not even libc.


## Installation

See the [Inspeqtor wiki](https://github.com/mperham/inspeqtor/wiki) for complete documentation.


## Requirements

Linux 3.0+.  It will run on OS X.  FreeBSD is untested.  It uses about 5-10MB of RAM at runtime.


## Upgrade

[Inspeqtor Pro](http://contribsys.com/inspeqtor) is the commercial version of Inspeqtor and offers more
features, official support and a non-GPL license:

* Monitor legacy sysvinit services with PID files
* Monitor daemon-specific metrics (e.g. redis, memcached, mysql, nginx...)
* Send alerts to Slack or other team chat rooms

See the [wiki documentation](https://github.com/mperham/inspeqtor/wiki#inspeqtor-pro) for
in-depth documentation around each Pro feature.


## License

Licensed under GPLv3.


# Author

inspeqtor is written by [Mike Perham](http://twitter.com/mperham) of [Contributed Systems](http://contribsys.com).  We build awesome open source-based infrastructure to help you build awesome apps.

We also develop [Sidekiq](http://sidekiq.org) and sell [Sidekiq Pro](http://sidekiq.org/pro), the best Ruby background job processing system.
