# Redacted

**This software is still under active development and should be considered beta quality**

Redacted monitors your application infrastructure.  It gathers and verifies key metrics
from all the moving parts in your application and alerts you when something
looks wrong.  It understands the application deployment workflow so it
won't bother you during a deploy.

What it does:

* Monitor systemd-, upstart-, runit- or launchd-managed services
* Monitor process memory and CPU usage
* Monitor daemon-specific metrics (e.g. redis, memcached, mysql, nginx...)
* Monitor and alert based on host CPU, load, swap and disk usage
* Alert or restart a process if a rule threshold is breached
* Alert if a process disappears or changes PID
* Signal deploy start/stop to silence alerts during deploy

What it doesn't:

* monitor or control arbitrary processes, services must be init-managed
* have *any* runtime dependencies at all, not even libc.


## Installation

See the [Redacted wiki](https://github.com/mperham/redacted/wiki) for complete documentation.


## Requirements

Linux 3.0+.  It will run on OS X.  FreeBSD is untested.  It uses about 5-10MB of RAM at runtime.


## Upgrade

[Redacted Pro](http://contribsys.com/redacted) is the commercial upgrade for Redacted and offers more
features, official support and a non-GPL license:

* Monitor legacy /etc/init.d services with PID files
* Route alerts to different teams or individuals
* Send alerts to Slack or other team chat rooms
* More to come...

To buy, go to the [homepage](http://contribsys.com/redacted) and select
the right plan for your organization.

See the [wiki documentation](https://github.com/mperham/redacted/wiki#redacted-pro) for
in-depth documentation around each Pro feature.


## License

Licensed under GPLv3.


# Author

Redacted is written by [Mike Perham](http://twitter.com/mperham) of [Contributed Systems](http://contribsys.com).  We build awesome open source-based infrastructure to help you build awesome apps.

We also develop [Sidekiq](http://sidekiq.org) and sell [Sidekiq Pro](http://sidekiq.org/pro), the best Ruby background job processing system.
