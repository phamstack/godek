```
graph
  - generated: auto generated by gqlgen
  - resolvers:
    ~ mutation.resolvers.go: mutations
    ~ query.resolvers.go: queries
    ~ resolver.go: root resolver
  - schema: graphql schemas - types: new schemas that matches to @models
    ~ \*.graphql: predefined graphql boilerplates
lib: helpers and middlewares
models
  - services.go: \*Services umbrella singleton for backend
  - *.go: microservices under *Services singleton
server.go: init services, database, and middlewares
gqlgen.yml: gqlgen rules within application
```
