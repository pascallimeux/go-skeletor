# go-skeletor

go template project
with configuration, logger, tests and docker

## Initialisation du projet

```
go mod init example.com/skeleton
go mod tidy
go mod vendor
```

## Création d'un binaire

`make build`

## Execution des tests

`make test`

## Effacer les binaires

`make clean`

## Création de l'image Docker

`make docker-build`

## Démarrage du conteneur Docker

`make docker-run`