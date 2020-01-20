This file is just a loose guideline for detailing packet details. It contains the following:

- Packet Name
- Packet Description
- Packet OPcode
- Packet Arguments (if any)

Packet payload are done as follows, to ensure that future expansion is possible. Each packet length is variable - however, the server will expect n arguments, depending on the packet. Each argument will be a specific byte length, which is hard-coded in the client.

|8 Bytes|? bytes
|-------|------|
|Op Code|Arguments

For Example, a ping packet will simply be "1" with no other arguments. If it receives a 1, the server will discard any further information from it. 

| Packet Name | Packet Description | OP Code (Int8)| Arguments
|-------------|--------------------|---------|----------
|Init. Connection|Initializes the connection between the server and client.|0|User ID (8 bytes?)
|Ping|Pings the server.|1|n/a
|Quit|Closes the connection.|2|n/a
|Map Update|Updates the map.|10|