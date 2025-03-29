# Insult Connector 2000

A Go network dialer implementation with a sole purpose in life: insult you!

## Description

<img src="docs/insult-connector-250px.jpeg" align="right" >
The Insult Connector 2000 (or ICU2K for short) is a critical piece of networking
infrastructure. It was written with the specific purpose of insulting you if you
try to open a proxied network connections.


ICU2K adds to Go's networking ecosystem by providing a new
[Dialer](https://pkg.go.dev/net#Dial) implementation that never connects but
is guaranteed to return a random insult in an error every time you try to connect.

It is important to note that the connector's Dialer never opens a network
connection or handles JWTs in any way (wink wink ;-)  

## Jokes Aside...

The true purpose of this repository is to demonstrate how to enable a sustainable
VEX process and lifecycle to manage explitability data in a repository.

The project hosted here is especially crafted to show as vulnerable to
[CVE-2025-22870](https://nvd.nist.gov/vuln/detail/CVE-2025-22870) and [CVE-2025-27144](https://nvd.nist.gov/vuln/detail/CVE-2025-27144) 
when a scanner looks at it. But if you inspect the code you'll
notice it is not affected by those vulnerabilities.

### VEX To the Rescue

This is where VEX (Vulnerabilityu Exploitability EXchange) and 
[vexflow](https://github.com/carabiner-dev/vexflow) come to the rescue.

<img src="https://avatars.githubusercontent.com/u/121361164?s=200&v=4" align="left">

Setting up an [OpenVEX](http://github.com/openvex) flow for the project lets you suppress
non exploitable vulnerabilities in the scanner output allowing vulnerability 
policies to pass.

Vexflow lets maintainers submit assessments about the
impact that vulnerabilities have on your software through the familiar
GitHub issues interface. Vexflow handles the generation, attesting and signing of OpenVEX data, this lab teaches you how to it set up in your
repositories. 

## Contributions Welcome!

This lab is open source. We have carefully set up a branch structure for the purposes of learning, but if you think of ways we can improve it
we're always happy to get help!

This lab is Copyright &copy; 2025 by Carabiner Systems, Inc. The code is
freely available under the Apache 2.0 license, the lab contents are released under the Creative Commons BY-SA 4.0 license, meaning you can use it ditribute it build of it, even for commercial purposes as long as you credit the original authors. Have fun remixing!
