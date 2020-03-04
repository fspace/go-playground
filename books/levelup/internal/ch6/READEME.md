>304 All the Things
Last-Modified headers aren’t just for static files. If you design your service well,
you can use the same process to say that data has yet to be updated since a given
time. If a user has data that’s still valid, you can return a 304 and the page doesn’t
have to be rerendered, the data won’t need to be fetched from the database, and
everything is superfast for the user:
GET /assets/test.txt HTTP/1.1
If-Modified-Since:Mon, 01 Jun 2015 09:50:11 GMT
HTTP/1.1 304 Not Modified
Date: Mon, 01 Jun 2015 09:53:56 GMT
