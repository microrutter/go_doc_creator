# go_doc_creator

## General

This little application currently will scan your cypress files and will record the following:

- The text in describe and will use it as the main title
- The text in it and will use that as a subtitle
- any comments using the denominator given in the config

With this data it will then create a notion page (details of which again should be recorded in the config) with those details
so you will have something similar to this [notion page](https://panoramic-sugar-fb9.notion.site/metric-journey-cy-ts-3ed6725db25840d3a0b4f67813415c0e)

This can then keep your documentation about your tests dynamically up to date

### Config File

The config file should be a yaml file similar to the test_conf.yaml file found in test_files
The below is a description on the fields

- maintitle, what identifier do you want to use to get the main title of the doc 
- split, what is the starting point and end point of the useful text within the title elements
- subtitle, like maintitle but this is the subtitle element that may split your tests up
- comment, what identifys a comment
- output, currenlty only notion is supported but if requested and popular may add more in future
-- type, the type of output currently only notion
-- secret, the bearer token to talk to notion
-- startingpage, the main page the app has been given permission to (32 digit number at end of url)
-- imageurl, a image you would like to use as a cover for your pages

### Use

If you are running on linux just download the go_doc_creator executable then run:

./go_doc_creator -conffile=pathtoconffile -topdir=pathtotests

for other systems download the repo and either run

go build

and then follow the directions above or

go run main.go -conffile=pathtoconffile -topdir=pathtotests

