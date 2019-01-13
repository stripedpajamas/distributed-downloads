# Roadmap

## Stack
  - Backend/CLI: **Golang**
  - DB: **TBD** or **N/A**
  - Frontend/Web UI: **TBD** or **N/A**

## Version 0.1.0
  - Basic node logic defined and implemented
    - Concurrently serve and download chunks of a file
    - Stitch together all downloaded chunks of file into single complete file
    - Request desired chunks from local peers before remote peers
    - Integrity of chunks are checked upon receipt at all hops
  - Local/remote peer addresses are read in from a configuration file
  - Controller mode: ask local peers to download specific chunks of files from specific remote nodes and to signal controller when done
  - Assist mode: listen for requests from controller node for addresses and chunk ids to download and then serve to controller
  - All actions and decisions made are logged

## Version 0.2.0
  - Nodes discover other nodes on same network
    - Nodes that are not in controller mode (not orchestrating other nodes to assist in downloading a file) and are not already assisting another node are automatically used to assist controller node
  - Nodes persist list of local / remote node addresses so discovery might not be necessary on every run

## Version 0.3.0
  - Adaptive peer selection
    - Nodes keep track of and constantly recompute the cost associated with downloading from each connected peer and choose peers with lowest cost
  - Failure resistance
    - If a peer goes offline, the chunks designated to that peer are reassigned
    - If a the download speed from a given peer declines past a certain threshold, the chunks designated to that peer are reassigned to a faster peer, unless it would make sense to let the slower peer finish


