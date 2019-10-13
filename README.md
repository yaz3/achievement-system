### Achievements system Osmo

- copy up app
- build app
- optimise image
- start postgres
- start server

From inside directory:
- `docker build --tag=achievements-system .`
- `docker-compose up`

debug
- `docker-compose up -d`
- `docker-compose logs -f -t`


#### Todo list:

- Player -> getAllByGame
- Data form of end game
- How do we create and merge the data from Endgame and generated Statistiques ? (map[string]interface{} seems to be the easier way)
- How to extend our conditions map
- In order to generate the extra statistique we need to iterate on a sorted map since generated stats can be used in the creation of incoming stats, so we need to 'sort' the array take a look a sorted way to iterate over a map since range is random ! 
- Add config file
