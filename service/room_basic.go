package service

type RoomBasic struct {
	Number       string `json:"number" bson:"number"`               //房间号
	Name         string `json:"name" bson:"name"`                   //房间名
	Info         int    `json:"info" bson:"info"`                   //房间信息
	UserIdentity string `json:"user_identity" bson:"user_identity"` //房主id
	CratedAt     int64  `json:"crated_at" bson:"crated_at"`
	UpdatedAt    int64  `json:"updated_at" bson:"updated_t"`
}

func (RoomBasic) CollectionName() string {
	return "room_basic"
}
