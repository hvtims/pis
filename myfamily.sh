#!/bin/bash

curl https://learn.zone01oujda.ma/assets/superhero/all.json | jq --argjson id "$HERO_ID" '. [] | select(.id == $id) | .connections.relatives' | sed 's/^"\(.*\)"$/\1/'