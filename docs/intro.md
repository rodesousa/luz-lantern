# Intro

## What ?

Luz-lantern or lantern is a tool which makes it possible to perform infra/soft tests quickly in a cluster/server. 

## Pourquoi lantern ?

( It is my opinion )

We must test something :

    - You already have a tool :
        - We must to add a test ... it allows to do ... must have permision ... remember test DSL ... have good lucky
        - Easy... after 15 tickets and 3 weeks 

    - no tool :

        - installing a tool ... need dependancies ... learn new DSL ... must have permision ...
        - " ssh id@server .... " ... into 15 servers

Don't forget environmental subtleties :

- env prod => + monitoring - test
- env inf  => - monitoring + test

## How lantern resolve this ?

The lantern is in Go => No dependancies ! 
Test configuration is in yaml... not too bad

Constituted :

- client lambda
[Client](http://)

- server with API
[Server](http://)

- Super Lantern mode which allowing to play test without agent(server) into other servers
[Super Lantern](http://)
