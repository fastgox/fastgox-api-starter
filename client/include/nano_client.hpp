#pragma once

#include "types.hpp"
#include "protocol.hpp"
// FastGox Nano C++客户端 - 基于POCO网络库
#include <Poco/Net/TCPSocket.h>
#include <Poco/Net/SocketAddress.h>
#include <Poco/Net/NetException.h>
#include <memory>
#include <atomic>
#include <unordered_map>
#include <mutex>

namespace fastgox {

class NanoClient {
public:
    using TcpSocket = Poco::Net::TCPSocket;
    
    explicit NanoClient();
    ~NanoClient();
    
    // 禁用拷贝，允许移动
    NanoClient(const NanoClient&) = delete;
    NanoClient& operator=(const NanoClient&) = delete;
    NanoClient(NanoClient&&) = default;
    NanoClient& operator=(NanoClient&&) = default;
    
    // 连接管理
    [[nodiscard]] bool connect(const std::string& host, uint16_t port);
    void disconnect();
    [[nodiscard]] bool isConnected() const noexcept;
    
    // 事件回调设置
    void setConnectCallback(ConnectCallback callback);
    void setDisconnectCallback(DisconnectCallback callback);
    void setErrorCallback(ErrorCallback callback);
    
    // 认证相关
    void login(const LoginRequest& request, ResponseCallback<LoginResponse> callback);
    void heartbeat(const HeartBeatRequest& request, ResponseCallback<HeartBeatResponse> callback);
    
    // 房间相关
    void joinRoom(const JoinRoomRequest& request, ResponseCallback<RoomInfoResponse> callback);
    void leaveRoom(const LeaveRoomRequest& request, ResponseCallback<RoomInfoResponse> callback);
    void getRoomInfo(const GetRoomInfoRequest& request, ResponseCallback<RoomInfoResponse> callback);
    
    // 通用请求方法
    template<typename RequestType, typename ResponseType>
    void request(const std::string& route, 
                const RequestType& req, 
                ResponseCallback<ResponseType> callback);
    
    // 通知方法
    template<typename RequestType>
    void notify(const std::string& route, const RequestType& req);
    
    // 启动客户端事件循环
    void run();
    void stop();
    
private:
    // 内部网络操作
    bool sendData(const std::vector<uint8_t>& data);
    bool receiveData();
    void processMessageBuffer();
    
    // 消息处理
    void handleMessage(const std::vector<uint8_t>& data);
    void handleResponse(uint32_t requestId, const json& data);
    void handlePush(const std::string& route, const json& data);
    
    // 发送消息
    void sendMessage(const std::vector<uint8_t>& data);
    
    // 生成请求ID
    uint32_t generateRequestId();
    
private:
    // POCO TCP Socket
    std::unique_ptr<TcpSocket> socket_;
    std::string host_;
    uint16_t port_;
    std::atomic<bool> connected_{false};
    std::atomic<bool> running_{false};
    
    // 网络缓冲区
    std::vector<uint8_t> receiveBuffer_;
    std::vector<uint8_t> messageBuffer_;
    std::thread networkThread_;
    
    // 回调函数
    ConnectCallback onConnect_;
    DisconnectCallback onDisconnect_;
    ErrorCallback onError_;
    
    // 请求管理
    std::atomic<uint32_t> nextRequestId_{1};
    std::unordered_map<uint32_t, std::function<void(const json&)>> pendingRequests_;
    std::mutex requestMutex_;
    
    // 推送消息处理
    std::unordered_map<std::string, std::function<void(const json&)>> pushHandlers_;
    std::mutex pushMutex_;
    
    // 异步操作
    std::thread networkThread_;
};

// 模板方法实现
template<typename RequestType, typename ResponseType>
void NanoClient::request(const std::string& route, 
                        const RequestType& req, 
                        ResponseCallback<ResponseType> callback) {
    if (!connected_) {
        if (onError_) {
            onError_("Client not connected");
        }
        return;
    }
    
    const uint32_t requestId = generateRequestId();
    const json reqJson = req;
    
    // 编码消息
    auto encoded = protocol::ProtocolHandler::encodeRequest(route, reqJson, requestId);
    if (!encoded) {
        if (onError_) {
            onError_("Failed to encode request: " + encoded.error());
        }
        return;
    }
    
    // 注册回调
    {
        std::lock_guard<std::mutex> lock(requestMutex_);
        pendingRequests_[requestId] = [callback](const json& response) {
            try {
                ResponseType resp = response.get<ResponseType>();
                callback(resp);
            } catch (const std::exception& e) {
                // TODO: 错误处理
            }
        };
    }
    
    // 发送消息
    sendMessage(encoded.value());
}

template<typename RequestType>
void NanoClient::notify(const std::string& route, const RequestType& req) {
    if (!connected_) {
        if (onError_) {
            onError_("Client not connected");
        }
        return;
    }
    
    const json reqJson = req;
    
    // 编码通知消息
    auto encoded = protocol::ProtocolHandler::encodeNotify(route, reqJson);
    if (!encoded) {
        if (onError_) {
            onError_("Failed to encode notify: " + encoded.error());
        }
        return;
    }
    
    // 发送消息
    sendMessage(encoded.value());
}

} // namespace fastgox
