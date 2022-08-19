{{ deleteTrigger 1 }}
{{ $embed := cembed
"title" "White Star Help"
    "description" "There are no WS Commands for Members to run.  An Officer/Comander will start `!rollcall` and a sign up message will be posted into <#731494882350333993>.  Select one of the Reactions <:destiny:623324746599235626><:blast:623325695472304147><:oshield:623325695568904214> on the message to assign yourself a role.  You can Select Any for just joining, but you are open to doing whatever in the WS."
     "color" 16448250
    "thumbnail" (sdict "url" "https://cdn.discordapp.com/emojis/722136072611430522.png?v=1")
    "fields" (cslice
        (sdict "name" "**Officer/Commander Commands**" "value" "----------------------")
        (sdict "name" "**`!rollcall ws1`**" "value" "This will start a rollcall for the first WS. " "inline" false )
        (sdict "name" "**`!rollcall ws2`**" "value" "Start a rollcall for a second WS match." "inline" false )
        (sdict "name" "**`!rollcall fun`**" "value" "Start a rollcall for the Fun Run. This is a Sactuary/ Loss Rewards run. Much more laid back from even the normal matches." "inline" false )
        (sdict "name" "**`!finish`**" "value" "Same detection of Corp role, like !rollcall, applies here.  This will clear out the values of a !rollcall.  `!finish ws1` and `!finish fun` for the respective rollcall." "inline" false ) )
}}

{{ sendMessage nil $embed }}