# Introduction

Riddler is a lightweight, performant microservice written in Go.

It stores an in-memory list of the 100,000 most common passwords as issued by the <a href="https://www.ncsc.gov.uk/" target="_new">National Cyber Security Centre</a> and provides a simple API to do whole-word, case-sensitive matching against this list.

## Deploy

### Docker container

A tiny alpine distribution running the go binary

<code>
	- git clone git@github.com:henricook/riddler.git
	- Create the files 'server.crt' and 'server.key' in the base directory for HTTPS (compulsory).
	- docker build .
	- Run/deploy as you wish
</code>

### Go Binary

<code>
	- Download from the <a href="https://github.com/henricook/riddler/releases">releases page</a>. 
	- Unpack the tar ball and select a suitable binary for your system.
	- Create the files 'server.crt' and 'server.key' in the base directory for HTTPS (compulsory).
	- Run/deploy as you wish
</code>

## Simple API

<code>
	POST /check-100k
</code>

Example Request Payload:
<code>
	{ "value": "string-to-check" }
</code>

Example Response Payload:
<code>
	{ "common": true }
</code>


# Contributors

Are very welcome. Please drop me an email riddler [at] henricook.com
