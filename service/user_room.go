package service

type UserRoom struct {
	UserIdentity    string `json:"user_identity" bson:"user_identity"`
	RoomIdentity    string `json:"room_identity" bson:"room_identity"`
	MessageIdentity int    `json:"message_identity" bson:"message_identity"`
	CratedAt        int64  `json:"crated_at" bson:"crated_at"`
	UpdatedAt       int64  `json:"updated_at" bson:"updated_t"`
}

func (UserRoom) CollectionName() string {
	return "user_basic"
}
