package auth

type Auth struct {
	AuthKeyId     int64   `bson:"_id" json:"_id"`
	UserId        int64   `bson:"user_id" json:"user_id"`
	AuthKey       []byte  `bson:"auth_key" json:"auth_key"`
	ExpiresIn     int32   `bson:"expires_in" json:"expires_in"`
	ValidSince    int32   `bson:"valid_since" json:"valid_since"`
	ValidUntil    int32   `bson:"valid_until" json:"valid_until"`
	Salt          int64   `bson:"salt" json:"salt"`
	TempAuthKey   int32   `bson:"temp_auth_key" json:"temp_auth_key"`
	Layer         int32   `bson:"layer" json:"layer"`
	Date          int32   `bson:"date" json:"date"`
	PermAuthKeyId int64   `bson:"perm_auth_key_id" json:"perm_auth_key_id"`
	Devices       *Device `bson:"devices" json:"devices"`
}
