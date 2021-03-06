This file is just a loose guideline for detailing packet details. It contains the following:

- Packet Name
- Packet Description
- Packet OP Code
- Packet Arguments (if any)

Packet payload are done as follows, to ensure that future expansion is possible. Each packet length is variable - however, the server will expect n arguments, depending on the packet. Each argument will be a specific byte length, which is hard-coded in the client.

|8 Bytes|? bytes
|-------|------|
|Op Code|Arguments

For Example, a ping packet will simply be "1" with no other arguments. If it receives a 1, the server will discard anything past that. 

#Receiving
| Packet Name | Packet Description | OP Code (Int8)| Arguments
|-------------|--------------------|---------|----------
|Init. Connection|Initializes the connection between the server and client. First part is initialized from the client - after the initial connection, it "logs in" with the User ID. If no user ID is provided, then it sends a user ID.|0|User ID (uint32)
|Ping|Pings the server.|1|n/a
|Quit|Closes the connection.|2|n/a
|Map Update|Updates the map.|10|X (uint8) , Y (uint8), Updated Value (uint8)
|Map Send|Gets the map file from the client. This will be a byte stream.|11|Map File
|Map Request|Requests the map file that is stored on the server.|12|n/a
|Map Open|Tells th server to open the map file which the player is currently in, or in the arguments given.|13|(optional): Dungeon ID (uint16), Map Number (unint16)
|Map Create|Tells the server the size of the map. We do not need to resend the  dungeon or floor, as this will always follow Map Open.|14|X (unint8), Y (unint8)
|Dungeon Update|Updates the current dungeon and floor the player is in.|20|Dungeon ID  (unint16), Map ID  (unint16)|
|Floor Update|Updates the current floor the player is in.|21|Map ID (uint16)
|End Dungeon|Changes dungeon ID and floor ID to 0 to reset both.|22|n/a
|Database Init|Tells the server that to start up the database.|30|n/a
|Party Member Init Changes|Tells the server to start caching changes to any party member.|31|n/a
|Party Member Create|Creates a level 1 party member in the database.|32|Name, Class, Portrait
|Party Member Update|Updates the party member details in the database.|33|Internal ID, current level, current total exp
|Party Member Finish|Tells the server to apply all changes to the party member(s).|38|n/a
|Database Finish|Tells the server to close the database.|39|n/a
#Sending
| Packet Name | Packet Description | OP Code (Int8)| Arguments
|-------------|--------------------|---------|----------
|Init. Connection|Finishes the initialization. Confirms that the client is properly logged in.|0|Error Code
|Ping|Pings the client.|1|n/a
|Quit|Closes the connection.|2|n/a
|User ID|Gives the client the user ID to use (if none was given)|3|uint32 id
|Map Update|Updates the map.|10|X, Y, Updated Value
|Map Send|Sends the map file to the client.|11|Map File
|Map Create|Requests the map details so it may create the file|12|n/a
|Map Error|Tells the client that an error has occurred while requesting the map|13|(Error Code)