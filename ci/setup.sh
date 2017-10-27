#!/bin/sh
fly -t ci set-pipeline --config pipeline.yml --load-vars-from credentials.yml --pipeline shell

