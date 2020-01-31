> Hostname Matching
  Patterns can also begin with a hostname, where only URLs for that hostname will
  match. A hostname match has a higher precedence than a pattern without a
  hostname.
  
  > Byte Arrays
    Byte arrays and slices are easily converted to strings with the string function,
    since a string is essentially a slice of bytes. This is a bit of an oversimplification,
    so if you’d like to understand more about how strings work in Go, check out this
    excellent post on the Go blog.6
