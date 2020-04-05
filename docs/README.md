![Go](https://github.com/henricook/riddler/workflows/Go/badge.svg)

# Introduction

Riddler is a lightweight, performant microservice written in Go.

It stores an in-memory list of the 100,000 most common passwords as issued by the <a href="https://www.ncsc.gov.uk/" target="_new">National Cyber Security Centre</a> and provides a simple API to do whole-word, case-sensitive matching against this list.

## Deploy it

### Docker container

A tiny alpine distribution running the go binary

<ol>
<li><code>git clone git@github.com:henricook/riddler.git</code></li>
<li>Create the files <code>server.crt</code> and <code>server.key</code> in the base directory for HTTPS (compulsory).</li>
<li><code>docker build .</code></li>
<li>Run/deploy as you wish</li>
</ol>

### Go Binary

<ol>
<li>Download from the <a href="https://github.com/henricook/riddler/releases">releases page</a>.</li> 
<li>Unpack the tar ball and select a suitable binary for your system.</li>
<li>Create the files 'server.crt' and 'server.key' in the base directory for HTTPS (compulsory).</li>
<li>Run/deploy as you wish</li>
</ol>

## Simple API

### Check most common passwords
<pre>
POST /check-100k
</pre>

Example Request Payload:

<pre>
{ "value": "string-to-check" }
</pre>

Example Response Payload:

<pre>
{ "common": true }
</pre>

### Ping
<pre>
GET /ping
</pre>

Example Response Payload:

<pre>
{ "response": "PONG" }
</pre>

# Contributors

Are very welcome. Please drop me an email riddler [at] henricook.com
