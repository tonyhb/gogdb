#!/bin/bash

# . /blink/devtools/scripts/CodeGeneratorFrontend.py ./blink/devtools/protocol.json --output_js_dir=./blink/devtools/front_end
# ./blink/devtools/scripts/generate_supported_css.py ./blink/core/css/CSSProperties.in ./blink/devtools/front_end/SupportedCSSProperties.js

cp blink/devtools/protocol.json blink/devtools/front_end/protocol.json
"$GOPATH/bin/go-bindata" -pkg assets -o gui/assets/assets.go -prefix blink/devtools/front_end blink/devtools/front_end/...
rm blink/devtools/front_end/protocol.json
