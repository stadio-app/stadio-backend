# Stadio Backend

## Clone repo
```
$ git@github.com:stadio-app/stadio-backend.git
```

## Docker
We run our development Postgres database in a Docker container, so make sure docker is installed (see [Docker desktop](https://docs.docker.com/desktop/)).

### Start docker container
```
$ make docker-container
```
This will also map the default `5432` port to a new `5433` port to avoid collisions. Check to see if the new postgres db is running on port `5433`.

## Atlas
We currently use Atlas to handle DB migrations. Make sure you have atlas fully installed on your machine (see [Atlas installation docs](https://atlasgo.io/getting-started#installation)).

## Run server
After completing all the steps from above, we can finally start our server.

```
$ make run
```

This command will simply generate all the entity files, graphql models, etc. Note that this will also make all the necessary migrations listed in `ent/migrate/migrations` (so make sure the docker container is still running!) and start the server on port `8080`.

In order to run the server in watch mode run `make watch` instead.


## Entities
When trying to add a new entity run
```
$ make ent-create entity=NameOfYourEntity
```
this will create a new file under `ent/schema` with the name of your entity. Read more about it [here](https://entgo.io/docs/schema-fields).
After you've added all the fields for the new entity run
```
$ make generate
```
this will generate a bunch of files in the `ent` and `graph` directories.

Note that running `make run` will always do this before running the server or any migrations.

## Migrations
For creating an entity migration run
```
$ make atlas-create entity=name_of_your_migration
```
this will create a new file under `ent/migrate/migrations` with a timestamp, and also update the `atlas.sum` file. The generated migration file will contain all the changes that were made to the entities (You can change the migration file if there is something inaccurate too).

Also note that after applying the migration a new schema will be created named `atlas_schema_revisions` which will hold the `atlas_schema_revisions` table listing all the successful migrations.

### Verification
We can check if all the migration files are valid and were not added manually by running
```
$ make atlas-validate
```

### Applying migration
There are two ways to apply these migrations, the first and the most simple way is to run the server using `make run`, `make watch` or `make build`. This will apply all the migrations. However, if you only want to apply the migration go with the second option as shown in the [Atlas docs](https://atlasgo.io/versioned/apply#existing-databases).
