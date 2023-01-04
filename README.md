# Setup
You need a postgresql database in order to test this.
You have to setup following environments variables:

```bash
export DB_HOST="url_of_the_db"
export DB_PORT=5432
export DB_USERNAME="username"
export DB_PASSWORD="password"
export DB_DATABASE="name_of_the_db"
```

Then you can create the user schema using the script `go run cmd/dbmigrate/db-migrate.go`

# Run the error

Run the error with `go run cmd/bugsavejson/bug-save-json.go` and enjoy the following error:
```bash
UPDATE "users" SET "securityID"='123456789',"metadata"='{"Hello":"World","World":42}',"age"=17,"isMajor"=false WHERE "id" = 'f7a872e8-d3f1-4586-8a3f-1c54d2ed7d2e'
Can't save user: failed to encode args[1]: unable to encode &schema.serializer{Field:(*schema.Field)(0xc0002b65a0), Serializer:schema.SerializerInterface(nil), 
SerializeValuer:schema.JSONSerializer{}, Destination:reflect.Value{typ:(*reflect.rtype)(0x1584920), ptr:(unsafe.Pointer)(0xc000245470), flag:0x199}, 
Context:(*context.emptyCtx)(0xc00009c000), value:interface {}(nil), fieldValue:(*database.UserMetadata)(0xc00027acf0)} into text format for jsonb (OID 3802): 
json: unsupported value: encountered a cycle via *schema.Schema
```