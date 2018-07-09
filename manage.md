## Examples
### A few quick examples for Windows and Unix, using the command line tool curl:

Get a list of vhosts:
:: Windows
```
C:\> curl -i -u guest:guest http://localhost:15672/api/vhosts
```
# Unix
```
$ curl -i -u guest:guest http://localhost:15672/api/vhosts
```

HTTP/1.1 200 OK
Server: MochiWeb/1.1 WebMachine/1.10.0 (never breaks eye contact)
Date: Mon, 16 Sep 2013 12:00:02 GMT
Content-Type: application/json
Content-Length: 30

[{"name":"/","tracing":false}]
Get a list of channels, fast publishers first, restricting the info items we get back:
:: Windows
C:\> curl -i -u guest:guest "http://localhost:15672/api/channels?sort=message_stats.publish_details.rate&sort_reverse=true&columns=name,message_stats.publish_details.rate,message_stats.deliver_get_details.rate"

# Unix
$ curl -i -u guest:guest 'http://localhost:15672/api/channels?sort=message_stats.publish_details.rate&sort_reverse=true&columns=name,message_stats.publish_details.rate,message_stats.deliver_get_details.rate'

HTTP/1.1 200 OK
Server: MochiWeb/1.1 WebMachine/1.10.0 (never breaks eye contact)
Date: Mon, 16 Sep 2013 12:01:17 GMT
Content-Type: application/json
Content-Length: 219
Cache-Control: no-cache

[{"message_stats":{"publish_details":{"rate" ... (remainder elided)
Create a new vhost:
:: Windows
C:\> curl -i -u guest:guest -H "content-type:application/json" ^
      -XPUT http://localhost:15672/api/vhosts/foo

# Unix
$ curl -i -u guest:guest -H "content-type:application/json" \
   -XPUT http://localhost:15672/api/vhosts/foo

HTTP/1.1 204 No Content
Server: MochiWeb/1.1 WebMachine/1.10.0 (never breaks eye contact)
Date: Mon, 16 Sep 2013 12:03:00 GMT
Content-Type: application/json
Content-Length: 0
Note: you must specify application/json as the mime type.

Note: the name of the object is not needed in the JSON object uploaded, since it is in the URI. As a virtual host has no properties apart from its name, this means you do not need to specify a body at all!

Create a new exchange in the default virtual host:
:: Windows
C:\> curl -i -u guest:guest -H "content-type:application/json" ^
       -XPUT -d"{""type"":""direct"",""durable"":true}" ^
       http://localhost:15672/api/exchanges/%2F/my-new-exchange

# Unix
$ curl -i -u guest:guest -H "content-type:application/json" \
    -XPUT -d'{"type":"direct","durable":true}' \
    http://localhost:15672/api/exchanges/%2F/my-new-exchange

HTTP/1.1 204 No Content
Server: MochiWeb/1.1 WebMachine/1.10.0 (never breaks eye contact)
Date: Mon, 16 Sep 2013 12:04:00 GMT
Content-Type: application/json
Content-Length: 0
Note: we never return a body in response to a PUT or DELETE, unless it fails.

And delete it again:
:: Windows
C:\> curl -i -u guest:guest -H "content-type:application/json" ^
       -XDELETE http://localhost:15672/api/exchanges/%2F/my-new-exchange

# Unix
$ curl -i -u guest:guest -H "content-type:application/json" \
    -XDELETE http://localhost:15672/api/exchanges/%2F/my-new-exchange

HTTP/1.1 204 No Content
Server: MochiWeb/1.1 WebMachine/1.10.0 (never breaks eye contact)
Date: Mon, 16 Sep 2013 12:05:30 GMT
Content-Type: application/json
Content-Length: 0
