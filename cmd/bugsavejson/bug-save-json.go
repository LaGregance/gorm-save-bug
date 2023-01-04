package main

import (
	"github.com/LaGregance/gorm-save-bug/database"
	"github.com/google/uuid"
	"log"
	"math/big"
)

func main() {
	db, err := database.OpenDB()
	if err != nil {
		log.Fatalln("Can't open the DB: ", err)
	}

	// 1. Create User
	id, _ := uuid.NewRandom()
	user := &database.User{
		ID:         id,
		SecurityID: big.NewInt(123456789),
		Age:        17,
		Metadata: &database.UserMetadata{
			Hello: "World",
			World: 42,
			Big:   big.NewInt(42),
		},
	}

	err = db.Save(user).Error
	if err != nil {
		log.Fatalln("Can't save user: ", err)
		// I get the error:
		// failed to encode args[1]: unable to encode &schema.serializer{Field:(*schema.Field)(0xc000000780), Serializer:schema.SerializerInterface(nil),
		// SerializeValuer:schema.JSONSerializer{}, Destination:reflect.Value{typ:(*reflect.rtype)(0x1584e60), ptr:(unsafe.Pointer)(0xc000223470), flag:0x199},
		// Context:(*context.emptyCtx)(0xc00002a0c8), value:interface {}(nil), fieldValue:(*database.UserMetadata)(0xc0002074e0)}
		// into text format for jsonb (OID 3802): json: unsupported value: encountered a cycle via *schema.Schema
	}

	// 2. Update user
	/*user.SecurityID = big.NewInt(987654321)
	user.Age = 19
	user.Metadata.Hello = "Bonjour"
	user.Metadata.Big = big.NewInt(123456)

	err = db.Save(user).Error
	if err != nil {
		log.Fatalln("Can't save user: ", err)
	}

	// 3. Check the user is well updated
	fetched := new(database.User)
	err = db.First(fetched, fmt.Sprintf("id='%s'", id.String())).Error
	if err != nil {
		log.Fatalln("Can't fetch user: ", err)
	}

	fmt.Printf("SecurityID: %s / %s\n", user.SecurityID, fetched.SecurityID)
	fmt.Printf("Age: %d / %d\n", user.Age, fetched.Age)
	fmt.Printf("Major: %v / %v\n", user.IsMajor, fetched.IsMajor)*/
}
