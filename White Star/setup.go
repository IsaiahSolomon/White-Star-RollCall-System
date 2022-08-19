{{/* This command has a Set and a Get.  1st argument is set to set values, or get to get the values that were set. */}}

{{/* Channel IDs */}}
{{ $wsLounge := 1009948658105602171 }}
{{ $signupsChannel := 1009948698022780948 }}
{{ $testChannel := 1009950153769230347 }}

{{/* Roles for WS */}}
{{ $officerRoleID := 1009948405805629491 }}
{{ $pingRole := 1009948061616832573 }}
{{/* -- WS 1 -- */}}
{{ $ws1RoleID := 1009948131510726667 }}
{{ $ws1OfficerRoleID := 1009948149139386438 }}
{{/* -- WS 2 -- */}}
{{ $ws2RoleID := 1009948197965287435 }}
{{ $ws2OfficerRoleID := 1009948221403050016 }}
{{/* -- Fun Run -- */}}
{{ $funRunID := 1009948253376229457 }}

{{/* ------ Do Not change anything below unless you know what you are doing. ------*/}}

{{ $args := "" }}
{{ $args2 := "" }}
{{ if .CmdArgs }}
	{{ $args = (index .Args 1) }}
    {{ $args2 = (index .Args 2) }}
{{ end }}

{{ if .CmdArgs }}
    {{ if eq $args "set" }} {{/* Setting the values into the DB.  Must run `setup set roles` and `setup set channels`. */}}
        {{ if eq $args2 "channels" }}
            {{/* Enter Channel IDs into the DB */}}
            {{ dbSet 4500 "wsLounge" (toString $wsLounge) }}
            {{ dbSet 4500 "signupsChannel" (toString $signupsChannel) }}
            {{ dbSet 4500 "testChannel" (toString $testChannel) }}
        {{ end }}

        {{ if eq $args2 "roles" }}
            {{/* Enter Role IDs into the DB */}}
            {{ dbSet 4500 "officerRoleID" (toString $officerRoleID) }}
            {{ dbSet 4500 "wsPingRole" (toString $pingRole) }}
            {{/* -- WS 1 -- */}}
            {{ dbSet 4500 "ws1RoleID" (toString $ws1RoleID) }}
            {{ dbSet 4500 "ws1OfficerRoleID" (toString $ws1OfficerRoleID) }}
            {{/* -- WS 2 -- */}}
            {{ dbSet 4500 "ws2RoleID" (toString $ws2RoleID) }}
            {{ dbSet 4500 "ws2OfficerRoleID" (toString $ws2OfficerRoleID) }}
            {{/* -- Fun Run -- */}}
            {{ dbSet 4500 "funRunID" (toString $funRunID) }}
        {{ end }}
        Values Set.
    {{ else if eq $args "get" }} {{/* If you want to get the values from the DB to verify them. */}}
        {{ if eq $args2 "channels" }}
            {{/* Get Channel IDs into the DB */}}
            Lounge Channel: {{ (dbGet 4500 "wsLounge").Value }}
            SignUps Channel: {{ (dbGet 4500 "signupsChannel").Value }}
            Test Channel: {{ (dbGet 4500 "testChannel").Value }}
        {{ end }}
        {{ if eq $args2 "roles" }}
            {{/* Get Role IDs into the DB */}}
            Officer Role: {{ (dbGet 4500 "officerRoleID").Value }}
            WS Ping Role: {{ (dbGet 4500 "wsPingRole").Value }}
            ----------- WS 1 Roles -----------
            WS1 Role: {{ (dbGet 4500 "ws1RoleID").Value }}
            WS1 Officer Role: {{ (dbGet 4500 "ws1OfficerRoleID").Value }}
            ----------- WS 2 Roles -----------
            WS2 Role: {{ (dbGet 4500 "ws2RoleID").Value }}
            WS2 Officer Role: {{ (dbGet 4500 "ws2OfficerRoleID").Value }}
            ----------- Fun Run Role -----------
            FR Role: {{ (dbGet 4500 "funRunID").Value }}
        {{ end }}
    {{ else if eq $args "del" }}
        {{ if eq $args2 "channels" }}
            {{/* Remove IDs from the DB */}}
            {{ dbDel 4500 "wsLounge" }}
            {{ dbDel 4500 "signupsChannel" }}
            {{ dbDel 4500 "testChannel" }}
        {{ end }}

        {{ if eq $args2 "roles" }}
            {{/* Remove Role IDs from the DB */}}
            {{ dbDel 4500 "officerRoleID" }}
            {{ dbDel 4500 "wsPingRole" }}
            {{/* -- WS 1 -- */}}
            {{ dbDel 4500 "ws1RoleID" }}
            {{ dbDel 4500 "ws1OfficerRoleID"}}
            {{/* -- WS 2 -- */}}
            {{ dbDel 4500 "ws2RoleID" }}
            {{ dbDel 4500 "ws2OfficerRoleID" }}
            {{/* -- Fun Run -- */}}
            {{ dbDel 4500 "funRunID" }}
        {{ end }}
        Entries Removed.
    {{ end }}
{{ else }}
    Please enter Set or Get for the arguments.
{{ end }}