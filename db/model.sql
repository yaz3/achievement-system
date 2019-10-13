create extension pgcrypto;

drop table if exists game;
drop table if exists player;
drop TYPE if exists game_state;

CREATE TYPE game_state AS ENUM ('loading', 'started', 'ended');

create table game (
    id serial primary key,
	team1 integer[],
	team2 integer[],
	nb_players integer,
    state game_state,
    CHECK (nb_players > 2 AND nb_players < 6)
);
comment on column game.team1 is 'a list of players ids';
comment on column game.team2 is 'a list of players ids';

create table player (
    id serial primary key,
	name varchar(20) not null UNIQUE,
	wins integer,
	losses integer,
    time_played integer,
	achievements json
);
