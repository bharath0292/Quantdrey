package brokerentity

type FlattradeConfig struct {
	ClientId  string `bson:"clientId"`
	ApiKey    string `bson:"apiKey"`
	ApiSecret string `bson:"apiSecret"`
}
