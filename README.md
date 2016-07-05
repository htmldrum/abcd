abcd
======
ABC News Australia command line client and web server

## Design
There are 2 executables available:
- abcd - crawling/indexing daemon
- abcli - viewer front-end

Repository follows pattern set out in https://camlistore.googlesource.com/camlistore/+/master
for multiple commands.

## Roadmap
If development continues beyond 1.0, it woud be to include the following features
- A framework for sepparating the executables into a distributed kubernettes
  pod
- Support for more media types
- Support for more interactions with abc media types

## TODO
- Implement client contact acceptance testing for abcd - http://agouti.org/

## DONE
- Implement fs-agnostic testing: https://github.com/spf13/afero
