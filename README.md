Wormhole
========

This code should not be used in another project, it's only here for testing purpose. 

It partially implements a STUN client as described in the [RFC 5389](https://tools.ietf.org/html/rfc5389). This RFC defines the second version of STUN (Session Traversal Utilities for NAT).

The original purpose of this project is to implement a NAT traversal solution in Go using UDP and have 2 clients behind a NAT communicating.

Currently, this solution can send a datagram to a STUN server.

