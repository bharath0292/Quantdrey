package scalars

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type BsonId = bson.ObjectID

func MarshalBsonId(id bson.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, `"`+id.Hex()+`"`)
	})
}

func UnmarshalBsonId(v any) (bson.ObjectID, error) {
	s, ok := v.(string)
	if !ok {
		return bson.NilObjectID, fmt.Errorf("BsonId must be a string")
	}
	return bson.ObjectIDFromHex(s)
}
