inspeqtor
=========

Next generation process monitoring.

What it does:

* Monitor service memory and CPU usage
* Monitor system CPU, swap and disk usage
* Notify if services disappear or change PID outside of a deploy window
* Notify if services or system goes over defined utilization
* Define applications and deploy windows
* Notification schemes: email, webhook, script
* Extremely developer friendly:
  - Test configuration
  - Test notifications

What it doesn't:

* support PID files, which are one big race condition
* monitor arbitrary processes, processes must be system-managed
* know how to start/stop services.  Defers to upstart/systemd/init.d/runit/launchctl
* complain about multiple operations taking place in parallel
* have *any* runtime dependencies at all, not even libc.

The system is scanned every 30 seconds, this is called a *cycle*.

The default monitoring rules out of the box perform basic health checks:

* / partition is > 95% full
* Swap is more than 20% utilized
* CPU(user) is 95% for more than 2 cycles
* load(5min) is > 10
