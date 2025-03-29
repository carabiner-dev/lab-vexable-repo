# Insult Connector 2000

A Go dialer iplementation with a sole purpose in life: insult you.

## Description

<img src="docs/insult-connector-250px.jpeg" align="right" >
The Insult Connector 2000 or ICU2K for short is a critical piece of networking
infrastructure. It was written with the specific purpose of insulting you if you
try to open a proxied network connections.


ICU2K adds to Go's networking ecosystem by providing a new
[Dialer](https://pkg.go.dev/net#Dial) implementation that never connects but
is guaranteed to return a random insult in an error every time yuo try to connect. 
