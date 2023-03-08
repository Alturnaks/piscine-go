#!/bin/bash
curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"kdyussyu\"}}){id}}"}'| sed 's/[^0-9]*//g'

