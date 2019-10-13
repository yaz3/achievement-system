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
- How to extend our conditions map
- Add config file