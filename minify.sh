#!/bin/bash

# Check if the input directory parameter is provided
if [ -z "$1" ]; then
  echo "Error: Input directory parameter is missing."
  exit 1
fi

# Set the input directory
input_dir="$1"

# Loop through all CSS files in the input directory
for css_file in "$input_dir"/*.css; do
  # Check if the file exists and does not contain "min" in its name
  if [ -f "$css_file" ] && [[ "$css_file" != *min* ]]; then
    # Extract the base name of the file (without the extension)
    base_name=$(basename "$css_file" .css)
    # Set the output file name
    output_file="${input_dir}/${base_name}.min.css"
    # Run the minify command
    minify -o "$output_file" "$css_file"
  fi
done

# Loop through all JS files in the input directory
for js_file in "$input_dir"/*.js; do
  # Check if the file exists and does not contain "min" in its name
  if [ -f "$js_file" ] && [[ "$js_file" != *min* ]]; then
    # Extract the base name of the file (without the extension)
    base_name=$(basename "$js_file" .js)
    # Set the output file name
    output_file="${input_dir}/${base_name}.min.js"
    # Run the minify command
    minify -o "$output_file" "$js_file"
  fi
done