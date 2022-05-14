#!/bin/bash

grammarVersion="0.0.1"

domain="https://raw.githubusercontent.com"
organization="merideum"
repository="merit"
path="src/main/resources"
grammarFile="Merit.g4"

url="$domain/$organization/$repository/$grammarVersion/$path/$grammarFile"

echo "Pulling $grammarFile from $url"
echo ""
curl ${url} > Merit.g4