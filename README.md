# DNSEscapist
DNS exfiltration tool written in go that uses nslookup to send queries to a name server
It first turns the file into a base64 string that is then split into 63 pieces each to fit inside a subdomain.
