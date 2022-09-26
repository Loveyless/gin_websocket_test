package service

type MessageBasic struct {
	UserIdentity string `json:"user_identity" bson:"user_identity"`
	RoomIdentity string `json:"room_identity" bson:"room_identity"`
	Data         int    `json:"data" bson:"data"`
	CratedAt     int64  `json:"crated_at" bson:"crated_at"`
	UpdatedAt    int64  `json:"updated_at" bson:"updated_t"`
}

func (MessageBasic) CollectionName() string {
	return "user_basic"
}
