This file gives a basic overview of the SQL database, and how it's set up.

The database is called "labyrinth", and should be contacted on the local network (127.0.0.1) on port 3306. The tables are as follows:

| Database Name | Descriptions
| --- | --- |
| Users |Contains all user IDs, alongside anything else that might be necessary (i.e. Google auth details etc.). It also contains the account status, as well as when it was created. This is only ever accessed once, during user_init.|
| Party Members | Contains all details regarding party members. This includes names, their ID, and current stats (i.e. exp, base stats, level). It also contains references to the equipment. The structure per member is as follows: user ID, internal/client ID,  class ID, EXP, level
| Inventory | Inventory details. This contains things such as any currency, and any items. Items are dictated by their SQL entries. Specifically, there will be a user ID column, item ID, and quantity column.
