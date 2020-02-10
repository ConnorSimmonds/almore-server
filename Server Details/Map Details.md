This file is to give a quick overview of the folder structure and map file design.

The Maps folder contains all of the user's maps. Each user has a unique user ID, which is generated and stored client-side (unsure of how to get this ID onto other clients yet - maybe seeded via Google Accounts or the like?).

For example, the User ID 1 will be under `maps/1`, and all of their subsequent maps will be found here.

Each map is a .dng file (short for "dungeon" - which is the same file format used in the client for dungeons). The maps filename is `mapx_y.dng`, where x is the dungeon ID, and y is the floor number.