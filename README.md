# sshtrap
> A tarpitting SSH server implemented in Go. Based on this post on
> [Nullprogram][1]

This is a basic TCP server that is to act as a replacement for the OpenSSH daemon
on Linux. The objective is to grab a bot/script/scanner and hold onto the initial
connection for as long as possible by sending a random 4 byte string back to the
client for as long as they're willing to hold open the connection.

As per the [Nullprogram][1] post that inspired this code, this daemon just trickles a
perpetual SSH banner back to the client without ever intending to send the
sequence that would conclude this portion of the connection negotiation:

```
SSH-protoversion-softwareversion SP comments CR LF
```

As we never send that string, the client will spend hours or days waiting for
it. Meanwhile, you can run an actual SSH daemon on a higher range port and let
the scripts gum themselves up in your tarpit.

See the [RFC][2] for more information.

## Usage

TBC

## Todo

* Support for logging to syslog.
* Support a configuration file.
* Tracking of SSH clients stuck in the tarpit.
* ~Rewrite the bytes generator. I pinched it off of StackOverflow to PoC.~
* `systemd` unit file to run the server as a service.
* Docker file to build the server in a container.

## Copyright

Copyright &copy; 2019 Paul Stevens. All rights reserved.

## License

Licensed under the MIT license. See LICENSE for details.

[1]: https://nullprogram.com/blog/2019/03/22/
[2]: https://tools.ietf.org/html/rfc4253#section-4.2
