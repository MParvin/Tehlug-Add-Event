---
title: \"{{ .EVENT_TITLE}}\"
date: {{ .CURRENT_DATE }}
layout: event
type: event
eventNumber: {{ .EVENT_NUMBER }}
eventDate: {{ .EVENT_DATE }}
startTime: {{ .START_TIME }}
endTime: {{ .END_TIME }}
address: "{{ .ADDRESS }}"
registerLink: {{ .REGISTER_LINK }}
speakers: [{{ .SPEAKERS }}]
topics: [{{ .TOPICS }}]
lat: {{ .LAT }}
lng: {{ .LNG }}
locations: [{{ .LOCATION }}]
---

{{ .DESCRIPTION }}