# Flashback

## Project Description 

The project goal is to have a system processing my pictures, extracting meta-data information from the files (gps location, labels, when the picture as taken, ...). With this information I would to be able to process event. My first goal is to send once a week a chat message with a list of pictures taken at the same week but in previous years. Ideally have a generic system to be able to send this information by email or other mechanism.

## Architecture 

The application will be wrote in GoLang and run with docker.

In the first iteration :
    * pictures will be store locally  (keep in mind: find a mechanism to be able to process external source)
    * The album information will be store in memory (keep in mind: in the future the storage will be externalise)
    * Use google auth for authentication (Pictures *must* not be freely available on the net)
    * Chat system will be RocketChat (In the future we will have to support more system)
    * Provide a web interface to visualise pictures of the day

For the first version I do not expect support multi-user

Global information about the data flow : 

![](./docs/architecture.png)

## Requirement

* Unittest
* Pull Request 
* follow golang standart syntax

## LICENCE

GPL v3 see [LICENSE](./LICENSE) for detail.
