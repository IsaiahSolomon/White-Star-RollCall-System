{{/* WS Rollcall command. */}}
{{/* Currently CC 97 on Black Sun Server */}}

{{ $wsLounge := (dbGet 4500 "wsLounge").Value }}
{{ $signupChan := (dbGet 4500 "signupsChannel").Value }}
{{ $pingRole := (dbGet 4500 "wsPingRole").Value }}

{{/* Get arg data if command is run by a user */}}
{{ $wsThumbURL := "" }}
{{ $corpName := "" }}
{{ $wsCorpAbbr := "" }}
{{ $args := "" }}

{{ if .CmdArgs }}
	{{ $args = (index .Args 1) }}
{{ end }}

{{ if eq $args "fun" }}
	{{ $args = (index .Args 1) }}
	{{ $corpName = "Fun Run" }}
	{{ $wsCorpAbbr = "FUN" }}
	{{ $wsThumbURL = "https://cdn.discordapp.com/emojis/722136072611430522.png?v=1" }}
{{ else if eq $args "ws2" }}
	{{ $args = (index .Args 1) }}
	{{ $corpName = "WS Team 2" }}
	{{ $wsCorpAbbr = "WS2" }}
	{{ $wsThumbURL = "https://cdn.discordapp.com/emojis/722136072611430522.png?v=1" }}
{{ else if eq $args "ws1" }}
	{{ $corpName = "WS Team 1" }}
	{{ $wsCorpAbbr = "WS1" }}
	{{ $wsThumbURL = "https://cdn.discordapp.com/emojis/722136072611430522.png?v=1" }}
{{ end }}

{{ if eq $args "help" }}
	{{ $embed := cembed 
		"title" "***Rollcall Help***"
		"description" "!rollcall command help "
		"color" 4645612 
		"fields" (cslice 
			(sdict "name" "Command Usage:" "value" "!rollcall \nThis will start a rollcall for the Corp that the !rollcall starter is in." "inline" true)
			) 
		"thumbnail" (sdict "url" "https://cdn.discordapp.com/emojis/722136072611430522.png?v=1") 
	}}
	{{ sendMessage nil $embed }}
{{ else if or (eq $args "ws1") (eq $args "ws2") (eq $args "fun") }}
	{{ $numSignedUp := 0 }}
	
	{{ $aggressorList := cslice "None" }}
	{{ $defenderList := cslice "None"}}
	{{ $midRangeList := cslice "None" }}
	{{ $altList := cslice "None"}}
	{{ $watchList := cslice "None"}}
	{{ $userIDList := cslice "None"}}

	{{ $signupEmbed := cembed 
		"title" (joinStr "" $corpName " - ***White Star Signup***")
		"description" (joinStr "" "React below to request your preferred role. \nYou can reference builds in <#690954295272538112>. You can also ask for ideas in <#716314530740699186>\n\n Currently " $numSignedUp " signed up.")
		"color" 4645612 
		"fields" (cslice 
			(sdict "name" "Aggressor <:destiny:623324746599235626>" "value" "None" "inline" true) 
			(sdict "name" "Planet Def <:blast:623325695472304147>" "value" "None" "inline" true) 
			(sdict "name" "Mid-Range <:oshield:623325695568904214>" "value" "None" "inline" true) 
			(sdict "name" "Alts/ Standby 😺" "value" "None" "inline" true)
			(sdict "name" "Watch 👁️" "value" "None" "inline" true)) 
		"thumbnail" (sdict "url" $wsThumbURL) 
	}}
	{{ $msgID := sendMessageRetID $signupChan $signupEmbed }}
	{{ $stringMsgID := (toString $msgID) }}

	{{ addMessageReactions $signupChan $msgID "destiny:623324746599235626" "blast:623325695472304147" "oshield:623325695568904214" "😺" "👁️" }}

	{{ $setDBMap := sdict "corp" $wsCorpAbbr "messageID" $stringMsgID "aggressors" $aggressorList "defenders" $defenderList "midrangers" $midRangeList "alts" $altList "watchers" $watchList "userIDs" $userIDList }}

	{{ dbSet 4500 (joinStr "" "ws_" $wsCorpAbbr "_msgID") $stringMsgID }}
	{{ dbSet 4500 (joinStr "" "ws_" $wsCorpAbbr "_data") $setDBMap }}

	{{ sendMessageNoEscape $wsLounge (joinStr "" "<@&" $pingRole "> - **Roll Call for " $corpName " this week .** Posted in <#731494882350333993>" ) }}

{{ else }}
	{{ $mID := sendMessageRetID nil "Specify which WS you are starting." }}
	{{ deleteMessage nil $mID 5 }}
{{ end }}
{{ deleteTrigger 0 }}