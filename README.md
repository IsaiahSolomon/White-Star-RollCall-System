# White Star RollCall System

This is a Custom Command (CC) set made for the Discord bot YAGPDB at https://yagpdb.xyz/.  The rollcall setup is for Coordinating matches within a corp for the game Hades' Star.
<br>
<br>
Join the following Discord for and example of the setup: https://discord.gg/2yDaqACNRq
<br>
<br>
#White Star rollcall 2.0

![Image of Queue](/images/wsrollcall.PNG)

The rollcall command is used to start the ws sign-up and users sign-up with hitting a reaction.
Currently the rollcall command supports 3 sign-ups.  WS1 WS2 and a Fun Run.

<h1>Setup</h1>

1. On your discord server you will need the following Role setup, you don't need Admin for this bot setup.
<br>

![Image of Server Channels](/images/server_roles.PNG)

2. Then on Discord you will need to setup your channels somthing like the following.
<br>

![Image of Server Channels](/images/server_channels.PNG)


3. The #ws-lounge and #sign-ups are the most important.
<br>The other categories like WS 1, WS 2, Fun Run are used with the roles allowed to only see those channels.  
<br>

4. Copy the channel IDs and the role IDs to the setup.go config area at the top.
<br>

![Image of Server Channels](/images/setup.PNG)


5. You will now need to configure the commands on the YAG website for your discord server.  The custom commands should look something like the following there:
<br>

![Image of Server Channels](/images/ws_cc.PNG)


6. Now run the setup commands:
<br>
<br>
setup set channels
<br>
setup set roles
<br>
<br>
This is done in 2 commands to get around the DB calls limitation that YAG Bot has in place for free servers.  If you pay for a premium server, you can combine them into 1 command.
<br>

7.  You will probably want to limit the rollcall and finish commands to just Commanders or other role, like Officer, so that not everyone in the corp can start or clear things out.