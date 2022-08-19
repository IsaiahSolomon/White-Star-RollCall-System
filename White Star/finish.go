{{ $wsLounge := (dbGet 4500 "wsLounge").Value }}
{{ $signupsChannel := (dbGet 4500 "signupsChannel").Value }}

{{/* Get arg data if command is run by a user */}}
{{ $wsCorpAbbr := "" }}
{{ $wsRole := "" }}
{{ $wsOfficerRole := "" }}
{{ $args := "" }}
{{ $wsMatchData := sdict }}
{{ $pingRole := (dbGet 4500 "wsPingRole").Value }}
{{ if .CmdArgs }}
	{{ $args = (lower (index .Args 1)) }}
{{ end }}

{{ if eq $args "fun" }} {{/* Fun Run */}}
	{{ $wsCorpAbbr = "FUN" }}
	{{ $wsRole = (dbGet 4500 "funRunID").Value }}
{{ else if eq $args "ws2" }} {{/* WS 2 */}}
	{{ $wsCorpAbbr = "WS2" }}
	{{ $wsRole = (dbGet 4500 "ws2RoleID").Value }}
	{{ $wsOfficerRole = (dbGet 4500 "ws2OfficerRoleID").Value }}
{{ else if eq $args "ws1"  }} {{/* WS 1 */}}
	{{ $wsCorpAbbr = "WS1" }}
	{{ $wsRole = (dbGet 4500 "ws1RoleID").Value }}
	{{ $wsOfficerRole = (dbGet 4500 "ws1OfficerRoleID").Value }}
{{ end }}

{{ $getMap := (dbGet 4500 (joinStr "" "ws_" $wsCorpAbbr "_data")).Value }}
{{ if $getMap }}
	{{ $wsMatchData = sdict $getMap }}
{{ end }}

{{ if $wsMatchData }}
	{{ $GetmsgID := (dbGet 4500 (joinStr "" "ws_" $wsCorpAbbr "_msgID")).Value }}
	{{ $msgID := (toString $GetmsgID) }}
	{{ deleteMessage $signupsChannel $msgID 0}}

	{{ dbDel 4500 (joinStr "" "ws_" $wsCorpAbbr "_msgID") }}
	{{ dbDel 4500 (joinStr "" "ws_" $wsCorpAbbr "_data") }}
	{{ dbDel 4500 (joinStr "" $wsCorpAbbr "_WS_Rollcall_Lock") }}

	{{/* Send a message to WS Lounge that the match is cleared out. */}}
	{{ sendMessageNoEscape $wsLounge (joinStr "" "<@&" $wsRole "> - WS Cleared. What did you like/dislike? Anything we should change or do differently?") }}

	{{ $userIDListInit := cslice.AppendSlice ($wsMatchData.Get "userIDs") }}
	{{ range $index, $element := $userIDListInit }}
		{{ takeRoleID $element $wsRole }}
		{{ takeRoleID $element $wsOfficerRole }}
	{{ end }}
{{ else }}
	{{ $mID := sendMessageRetID nil "No current WS To clear." }}
	{{ deleteMessage nil $mID 5 }}
{{ end }}
{{ deleteTrigger 0 }}