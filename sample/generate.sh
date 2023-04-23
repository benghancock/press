#! /bin/sh

# a sample website generation script

OUTPUT_DIR=www
PAGE_TEMPLATE=src/templates/pages.html

mkdir -p $OUTPUT_DIR

for f in src/content/*.html; do
	fn=$(basename $f)
	press -c $f -t $PAGE_TEMPLATE > www/$fn
done
