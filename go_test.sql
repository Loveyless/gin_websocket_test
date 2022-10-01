/*
 Navicat Premium Data Transfer

 Source Server         : go聊天室test
 Source Server Type    : MongoDB
 Source Server Version : 40406
 Source Host           : 101.200.243.101:27017
 Source Schema         : go_test

 Target Server Type    : MongoDB
 Target Server Version : 40406
 File Encoding         : 65001

 Date: 01/10/2022 23:34:14
*/


// ----------------------------
// Collection structure for message_basic
// ----------------------------
db.getCollection("message_basic").drop();
db.createCollection("message_basic");

// ----------------------------
// Documents of message_basic
// ----------------------------
db.getCollection("message_basic").insert([ {
    _id: ObjectId("633325947e29000087006d65"),
    "user_identity": "用户唯一标识",
    "room_identity": "房间唯一标识",
    data: "发送的数据",
    "created_at": 1,
    "updated_at": 1
} ]);
db.getCollection("message_basic").insert([ {
    _id: ObjectId("63347119125b650283da3f78"),
    "user_identity": "81974471",
    "room_identity": "888",
    data: "你好",
    "crated_at": NumberLong("1664381209"),
    "updated_t": NumberLong("1664381209")
} ]);
db.getCollection("message_basic").insert([ {
    _id: ObjectId("6334711e125b650283da3f79"),
    "user_identity": "81974471",
    "room_identity": "888",
    data: "你好",
    "crated_at": NumberLong("1664381214"),
    "updated_t": NumberLong("1664381214")
} ]);
db.getCollection("message_basic").insert([ {
    _id: ObjectId("63347121125b650283da3f7a"),
    "user_identity": "81974471",
    "room_identity": "888",
    data: "你好",
    "crated_at": NumberLong("1664381217"),
    "updated_t": NumberLong("1664381217")
} ]);
db.getCollection("message_basic").insert([ {
    _id: ObjectId("63347121125b650283da3f7b"),
    "user_identity": "81974471",
    "room_identity": "888",
    data: "你好",
    "crated_at": NumberLong("1664381217"),
    "updated_t": NumberLong("1664381217")
} ]);
db.getCollection("message_basic").insert([ {
    _id: ObjectId("6334712f125b650283da3f7c"),
    "user_identity": "123",
    "room_identity": "888",
    data: "我不好",
    "crated_at": NumberLong("1664381231"),
    "updated_t": NumberLong("1664381231")
} ]);
db.getCollection("message_basic").insert([ {
    _id: ObjectId("63347132125b650283da3f7d"),
    "user_identity": "123",
    "room_identity": "888",
    data: "呵呵",
    "crated_at": NumberLong("1664381234"),
    "updated_t": NumberLong("1664381234")
} ]);
db.getCollection("message_basic").insert([ {
    _id: ObjectId("63347137125b650283da3f7e"),
    "user_identity": "81974471",
    "room_identity": "888",
    data: "哼",
    "crated_at": NumberLong("1664381239"),
    "updated_t": NumberLong("1664381239")
} ]);
db.getCollection("message_basic").insert([ {
    _id: ObjectId("633477146886e3eca251f7a8"),
    "user_identity": "81974471",
    "room_identity": "888",
    data: "不理你了",
    "crated_at": NumberLong("1664382740"),
    "updated_t": NumberLong("1664382740")
} ]);

// ----------------------------
// Collection structure for room_basic
// ----------------------------
db.getCollection("room_basic").drop();
db.createCollection("room_basic");

// ----------------------------
// Documents of room_basic
// ----------------------------
db.getCollection("room_basic").insert([ {
    _id: ObjectId("632ee04de6000000ea001845"),
    number: "房间号",
    name: "房间名",
    info: "房间简介",
    "user_identity": "房间创建者的唯一标识",
    "created_at": 1,
    "updated_at": 1
} ]);
db.getCollection("room_basic").insert([ {
    _id: ObjectId("633185eb120a0000ca006232"),
    number: "11",
    name: "test",
    info: "房间简介",
    "user_identity": "124",
    "created_at": 1,
    "updated_at": 1
} ]);

// ----------------------------
// Collection structure for user_basic
// ----------------------------
db.getCollection("user_basic").drop();
db.createCollection("user_basic");

// ----------------------------
// Documents of user_basic
// ----------------------------
db.getCollection("user_basic").insert([ {
    _id: ObjectId("632ede35e6000000ea001843"),
    identity: "123",
    username: "账号1",
    password: "密码1",
    nicname: "用户名",
    sex: 1,
    email: "邮箱",
    avatar: "头像",
    "created_at": 1,
    "updated_at": 1
} ]);
db.getCollection("user_basic").insert([ {
    _id: ObjectId("632feac9c26c000003001175"),
    identity: "81974471",
    username: "123",
    password: "123",
    nicname: "123",
    sex: 1,
    email: "1@qq.com",
    avatar: "123",
    "created_at": 1,
    "updated_at": 1
} ]);
db.getCollection("user_basic").insert([ {
    _id: ObjectId("63367fa922d475c77ef82958"),
    identity: "c77be1a5-6045-4204-bfaa-67daa6ba7e88",
    username: "1234",
    password: "",
    nicname: "",
    sex: NumberInt("0"),
    email: "59771463@qq.com",
    avatar: "",
    "created_at": NumberLong("1664516009"),
    "updated_at": NumberLong("1664516009")
} ]);
db.getCollection("user_basic").insert([ {
    _id: ObjectId("633732cf4e45d197af416467"),
    identity: "663821cb-b837-480a-88dd-174d76479d82",
    username: "2222",
    password: "",
    nicname: "",
    sex: NumberInt("0"),
    email: "loveyless@126.com",
    avatar: "",
    "created_at": NumberLong("1664561871"),
    "updated_at": NumberLong("1664561871")
} ]);

// ----------------------------
// Collection structure for user_room
// ----------------------------
db.getCollection("user_room").drop();
db.createCollection("user_room");

// ----------------------------
// Documents of user_room
// ----------------------------
db.getCollection("user_room").insert([ {
    _id: ObjectId("632ee153e6000000ea001846"),
    "user_identity": "用户的唯一标识",
    "room_identity": "房间的唯一标识",
    "message_identity": "消息的唯一标识",
    "room_type": "2",
    "created_at": 1,
    "updated_at": 1
} ]);
db.getCollection("user_room").insert([ {
    _id: ObjectId("633319707e29000087006d62"),
    "user_identity": "81974471",
    "room_identity": "888",
    "message_identity": "000",
    "room_type": "2",
    "created_at": 1,
    "updated_at": 1
} ]);
db.getCollection("user_room").insert([ {
    _id: ObjectId("63331c547e29000087006d63"),
    "user_identity": "81974471",
    "room_identity": "房间的唯一标识",
    "message_identity": "消息的唯一标识",
    "room_type": "2",
    "created_at": 1,
    "updated_at": 1
} ]);
db.getCollection("user_room").insert([ {
    _id: ObjectId("633320637e29000087006d64"),
    "user_identity": "123",
    "room_identity": "888",
    "message_identity": "000",
    "room_type": "2",
    "created_at": 1,
    "updated_at": 1
} ]);
db.getCollection("user_room").insert([ {
    _id: ObjectId("63372b769771000008004122"),
    "user_identity": "81974471",
    "room_identity": "980",
    "message_identity": "000",
    "room_type": 1,
    "created_at": 1,
    "updated_at": 1
} ]);
db.getCollection("user_room").insert([ {
    _id: ObjectId("63372b899771000008004123"),
    "user_identity": "c77be1a5-6045-4204-bfaa-67daa6ba7e88",
    "room_identity": "980",
    "message_identity": "000",
    "room_type": 1,
    "created_at": 1,
    "updated_at": 1
} ]);
