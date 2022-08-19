{{/* WS Reaction control for Rollcall. */}}
{{/* Currently CC 103 on Black Sun Server */}}

{{/* Initializing Variable */}}
{{ $reactMesID := .Reaction.MessageID }}
{{ $reactionID := $.Reaction.Emoji.ID }}
{{ $reactionName := $.Reaction.Emoji.Name }}

{{/* Get user information when they run the command */}}
{{ $userID := .User.ID }}
{{ $userNick := (getMember $userID).Nick }}
{{ $userName := .User.Username }}

{{/* Get database entries for message IDs */}}
{{ $GetmsgID_WS1 := (dbGet 4500 "ws_WS1_msgID").Value }}
{{ $msgID_WS1 := (toString $GetmsgID_WS1) }}
{{ $GetmsgID_WS2 := (dbGet 4500 "ws_WS2_msgID").Value }}
{{ $msgID_WS2 := (toString $GetmsgID_WS2) }}
{{ $GetmsgID_FUN := (dbGet 4500 "ws_FUN_msgID").Value }}
{{ $msgID_FUN := (toString $GetmsgID_FUN) }}

{{/* Start/Set of the variable values. */}}
{{ $args := "" }}
{{ $thumbURL := "" }}
{{ $cName := "" }}
{{ $wsCorpAbbr := "" }}
{{ $wsRole := "" }}
{{ $wsMatchData := sdict }}
{{ $continue := "no"}}
{{ $leaderRole := "void" }}

{{/* Assign the values to the Corp variables and get Data from the DB. */}}
{{ if and $msgID_WS1 (eq $msgID_WS1 (toString $reactMesID)) }} {{/* WS Team 1 */}}
	{{ $getMap := (dbGet 4500 (joinStr "" "ws_WS1_data")).Value }}
	{{ $wsMatchData = sdict $getMap }}
	{{ $thumbURL = "https://cdn.discordapp.com/emojis/722136072611430522.png?v=1" }}
	{{ $cName = "WS Team 1" }}
	{{ $wsCorpAbbr = "WS1" }}
	{{ $wsRole = (dbGet 4500 "ws1RoleID").Value }}
	{{ $continue = "go"}}
	{{ $leaderRole = (dbGet 4500 "ws1OfficerRoleID").Value }}
{{ else if and $msgID_WS2 (eq $msgID_WS2 (toString $reactMesID)) }} {{/* WS Team 2 */}}
	{{ $getMap := (dbGet 4500 "ws_WS2_data").Value }}
	{{ $wsMatchData = sdict $getMap }}
	{{ $cName = "WS Team 2" }}
	{{ $thumbURL = "https://cdn.discordapp.com/emojis/722136072611430522.png?v=1" }}
	{{ $wsCorpAbbr = "WS2" }}
	{{ $wsRole = (dbGet 4500 "ws2RoleID").Value }}
	{{ $continue = "go"}}
	{{ $leaderRole = (dbGet 4500 "ws2OfficerRoleID").Value }}
{{ else if and $msgID_FUN (eq $msgID_FUN (toString $reactMesID)) }} {{/* Fun Run */}}
	{{ $getMap := (dbGet 4500 "ws_FUN_data").Value }}
	{{ $wsMatchData = sdict $getMap }}
	{{ $cName = "Fun Run" }}
	{{ $thumbURL = "https://cdn.discordapp.com/emojis/722136072611430522.png?v=1" }}
	{{ $wsRole = (dbGet 4500 "funRunID").Value }}
	{{ $wsCorpAbbr = "FUN" }}
	{{ $continue = "go"}}
{{ end }}	 
{{/* If continue is go, the person is part of a value corp role from above. */}}
{{ if eq $continue "go" }}
	{{/* Has to be one of the already attached emoji - else ignored */}}
	{{ if or (eq $reactionID 623324746599235626) (eq $reactionID 623325695472304147) (eq $reactionID 623325695568904214) (eq $reactionID 623327700756463627) (eq $reactionName "😺") (eq $reactionName "👁️") }}
		{{/* Position selection for what the user reacted with. */}}
		{{ $posSelect := "" }}
		{{ $shipTypeRole := "" }}
		{{ if eq $reactionID 623324746599235626 }}
			{{ $posSelect = "aggressors" }}
			{{ $shipTypeRole = 969346789469720626 }}
		{{ else if eq $reactionID 623325695472304147 }}
			{{ $posSelect = "defenders" }}
			{{ $shipTypeRole = 969346884642693231 }}
		{{ else if eq $reactionID 623325695568904214 }}
			{{ $posSelect = "midrangers" }}
		{{ else if eq $reactionName "😺" }}
			{{ $posSelect = "alts" }}
		{{ else if eq $reactionName "👁️" }}
			{{ $posSelect = "watchers" }}
		{{ end }}
		{{ $addName := "" }}
		{{ if $userNick }}
			{{ $addName = $userNick }}
		{{ else }}
			{{ $addName = $userName }}
		{{ end }}
		{{ $List := "" }}
		{{ $IDList := "" }}
		{{ $userList := cslice.AppendSlice ($wsMatchData.Get $posSelect) }}
		{{ $uIDListInit := cslice.AppendSlice ($wsMatchData.Get "userIDs") }}
		{{ if $.ReactionAdded }} 
			{{/* *****Adding a role bc of reaction added to sign up message.***** */}}
			{{ giveRoleID $.User.ID $wsRole}}
			{{ if $shipTypeRole }}
				{{ giveRoleID $.User.ID $shipTypeRole}}
			{{ end }}

			{{ $userList := $userList.Append $addName }}
			{{ $wsMatchData.Set $posSelect $userList }}
			{{ range $index, $element := $userList }}
				{{ $List = (joinStr "" $List "\n" $element ) }} 
			{{ end }}
			{{ $uIDListInit := $uIDListInit.Append $userID }}
			{{ $wsMatchData.Set "userIDs" $uIDListInit }}
			{{ $officerRole := (dbGet 4500 "officerRoleID").Value }}
			{{ if hasRoleID $officerRole }}
				{{ giveRoleID $userID $leaderRole }}
			{{ end }}
			{{ $wsLounge := (dbGet 4500 "wsLounge").Value }}
			{{ sendMessageNoEscape $wsLounge (joinStr ""  "<@" $userID ">, You have been added to the WS for " $cName) }}
		{{ else }}
			{{/* *****Removing a role bc of reaction removed.****** */}}
			{{ takeRoleID $.User.ID $wsRole }}
			{{ if $shipTypeRole }}
				{{ takeRoleID $.User.ID $shipTypeRole}}
			{{ end }}

			{{ $List = cslice }}
			{{ range $index, $element := $userList }}
				{{ if ne (toString $element) (toString $addName) }}
					{{ $List = $List.Append $element }}
				{{ end }}
			{{ end }}
			{{ $wsMatchData.Set $posSelect $List }}

			{{ $IDList = cslice }}
			{{ range $index, $element := $uIDListInit }}
				{{ if ne (toString $element) (toString $userID) }}
					{{ $IDList = $IDList.Append $element }}
				{{ end }}
			{{ end }}
			{{ $wsMatchData.Set "userIDs" $IDList }}
			{{ $officerRole := (dbGet 4500 "officerRoleID").Value }}
			{{ if hasRoleID $officerRole }}
				{{ takeRoleID $userID $leaderRole }}
			{{ end }}
			{{ $wsLounge := (dbGet 4500 "wsLounge").Value }}
			{{ sendMessageNoEscape $wsLounge (joinStr ""  "<@" $userID ">, You have been removed from the WS for " $cName) }}
		{{ end }}

		{{ dbSet 4500 (joinStr "" "ws_" $wsCorpAbbr "_data") $wsMatchData }}

		{{ $getMap := (dbGet 4500 (joinStr "" "ws_" $wsCorpAbbr "_data")).Value }}
		{{ $wsMatchData = sdict $getMap }}

		{{/* Check through all the lists in the sign up array and build the list to be put into the embed. */}}
		{{ $aggressList := $wsMatchData.Get "aggressors" }}
		{{ $aggrlist := "" }}
		{{ range $index, $element := $aggressList }}
			{{ if ne $element "None" }}
				{{ $aggrlist = (joinStr ""  $aggrlist "\n" $element) }}  
			{{ else if and (eq (len $aggressList) 1) (eq $element "None") }}
				{{ $aggrlist = $element }}
			{{ end }}
		{{ end }}

		{{ $defList := $wsMatchData.Get "defenders" }}
		{{ $defendlist := "" }}
		{{ range $index, $element := $defList }}
			{{ if ne $element "None" }}
				{{ $defendlist = (joinStr ""  $defendlist "\n" $element) }}  
			{{ else if and (eq (len $defList) 1) (eq $element "None") }}
				{{ $defendlist = $element }}
			{{ end }}
		{{ end }}

		{{ $midrList := $wsMatchData.Get "midrangers" }}
		{{ $midlist := "" }}
		{{ range $index, $element := $midrList }}
			{{ if ne $element "None" }}
				{{ $midlist = (joinStr ""  $midlist "\n" $element) }} 
			{{ else if and (eq (len $midrList) 1) (eq $element "None") }}
				{{ $midlist = $element }}
			{{ end }}
		{{ end }}

		{{ $altsList := $wsMatchData.Get "alts" }}
		{{ $altlist := "" }}
		{{ range $index, $element := $altsList }}
			{{ if ne $element "None" }}
				{{ $altlist = (joinStr ""  $altlist "\n" $element) }} 
			{{ else if and (eq (len $altsList) 1) (eq $element "None") }}
				{{ $altlist = $element }}
			{{ end }}
		{{ end }}

		{{ $watchersList := $wsMatchData.Get "watchers" }}
		{{ $watchlist := "" }}
		{{ range $index, $element := $watchersList }}
			{{ if ne $element "None" }}
				{{ $watchlist = (joinStr ""  $watchlist "\n" $element) }} 
			{{ else if and (eq (len $watchersList) 1) (eq $element "None") }}
				{{ $watchlist = $element }}
			{{ end }}
		{{ end }}

		{{/* Add up the number of participants from each list. */}}
		{{ $math := (add (len $aggressList) (len $defList) (len $midrList) (len $altsList)) }}
		{{ $numSignedUp := (sub $math 4) }}

		{{/* Build the embed to be edited to the Sign Up message. */}}
		{{ $signupEm := cembed 
			"title" (joinStr "" $cName " - ***White Star Signup***")
			"description" (joinStr "" "React below to request your preferred role. \nYou can reference builds in <#690954295272538112>. You can also ask for ideas in <#716314530740699186>\n\n - Currently " $numSignedUp " signed up. -")
			"color" 4645612 
			"fields" (cslice 
				(sdict "name" "Aggressor <:destiny:623324746599235626>" "value" $aggrlist "inline" true) 
				(sdict "name" "Planet Def <:blast:623325695472304147>" "value" $defendlist "inline" true) 
				(sdict "name" "Mid-Range <:oshield:623325695568904214>" "value" $midlist "inline" true) 
				(sdict "name" "Alts/ Standby 😺" "value" $altlist "inline" true)
				(sdict "name" "Watch 👁️" "value" $watchlist "inline" true)) 
			"thumbnail" (sdict "url" $thumbURL) 
		}}
		{{- editMessage $.Reaction.ChannelID $.Reaction.MessageID (cembed $signupEm) -}}
	{{ end }}
{{ end }}