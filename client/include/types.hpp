#pragma once

#include <string>
#include <cstdint>
#include <chrono>
#include <nlohmann/json.hpp>

namespace fastgox {

using json = nlohmann::json;
using namespace std::chrono_literals;

// Nano协议常量
namespace protocol {
    // Package类型
    enum class PackageType : uint8_t {
        HANDSHAKE = 0x01,
        HANDSHAKE_ACK = 0x02,
        HEARTBEAT = 0x03,
        DATA = 0x04,
        KICK = 0x05
    };
    
    // Message类型
    enum class MessageType : uint8_t {
        REQUEST = 0x00,
        NOTIFY = 0x01,
        RESPONSE = 0x02,
        PUSH = 0x03
    };
    
    constexpr size_t PKG_HEAD_BYTES = 4;
    constexpr size_t MSG_FLAG_BYTES = 1;
    constexpr uint8_t MSG_COMPRESS_ROUTE_MASK = 0x01;
    constexpr uint8_t MSG_TYPE_MASK = 0x07;
}

// 认证相关结构体
struct LoginRequest {
    std::string Token;
    
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(LoginRequest, Token)
};

struct LoginResponse {
    int Code{0};
    std::string Message;
    std::string UserID;
    std::string Nickname;
    
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(LoginResponse, Code, Message, UserID, Nickname)
};

struct HeartBeatRequest {
    std::string ClientTime;
    
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(HeartBeatRequest, ClientTime)
};

struct HeartBeatResponse {
    std::string ServerTime;
    int64_t Timestamp{0};
    
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(HeartBeatResponse, ServerTime, Timestamp)
};

// 房间相关结构体
struct JoinRoomRequest {
    std::string RoomID;
    
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(JoinRoomRequest, RoomID)
};

struct LeaveRoomRequest {
    std::string RoomID;
    
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(LeaveRoomRequest, RoomID)
};

struct GetRoomInfoRequest {
    std::string RoomID;
    
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(GetRoomInfoRequest, RoomID)
};

struct RoomInfoResponse {
    int Code{0};
    std::string Message;
    std::string RoomID;
    std::vector<std::string> Players;
    int MaxPlayers{0};
    bool IsActive{false};
    
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(RoomInfoResponse, Code, Message, RoomID, Players, MaxPlayers, IsActive)
};

// 回调函数类型
template<typename T>
using ResponseCallback = std::function<void(const T&)>;

using ErrorCallback = std::function<void(const std::string&)>;
using ConnectCallback = std::function<void()>;
using DisconnectCallback = std::function<void()>;

} // namespace fastgox
