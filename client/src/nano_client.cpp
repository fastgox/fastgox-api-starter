#include "nano_client.hpp"
#include <iostream>
#include <chrono>
#include <Poco/Timespan.h>

namespace fastgox {

NanoClient::NanoClient() 
    : socket_(std::make_unique<TcpSocket>())
    , connected_(false)
    , running_(false)
    , receiveBuffer_(4096) {
    std::cout << "🎮 FastGox Nano客户端初始化完成" << std::endl;
}

NanoClient::~NanoClient() {
    stop();
    disconnect();
    std::cout << "🎮 FastGox Nano客户端已销毁" << std::endl;
}

bool NanoClient::connect(const std::string& host, uint16_t port) {
    try {
        host_ = host;
        port_ = port;
        
        std::cout << "🔗 连接服务器: " << host << ":" << port << std::endl;
        
        // 创建服务器地址
        Poco::Net::SocketAddress addr(host, port);
        
        // 连接服务器
        socket_->connect(addr);
        
        // 设置socket选项
        socket_->setNoDelay(true);  // 禁用Nagle算法，降低延迟
        socket_->setKeepAlive(true); // 启用keep-alive
        
        connected_ = true;
        std::cout << "✅ 连接成功!" << std::endl;
        
        // 触发连接回调
        if (onConnect_) {
            onConnect_();
        }
        
        return true;
        
    } catch (const Poco::Exception& e) {
        std::cerr << "❌ 连接失败: " << e.displayText() << std::endl;
        connected_ = false;
        
        if (onError_) {
            onError_("连接失败: " + e.displayText());
        }
        
        return false;
    }
}

void NanoClient::disconnect() {
    if (!connected_) return;
    
    try {
        std::cout << "🔌 断开连接..." << std::endl;
        
        connected_ = false;
        socket_->close();
        
        // 触发断开连接回调
        if (onDisconnect_) {
            onDisconnect_();
        }
        
        std::cout << "✅ 已断开连接" << std::endl;
        
    } catch (const Poco::Exception& e) {
        std::cerr << "⚠️ 断开连接时出错: " << e.displayText() << std::endl;
    }
}

bool NanoClient::isConnected() const noexcept {
    return connected_;
}

void NanoClient::setConnectCallback(ConnectCallback callback) {
    onConnect_ = std::move(callback);
}

void NanoClient::setDisconnectCallback(DisconnectCallback callback) {
    onDisconnect_ = std::move(callback);
}

void NanoClient::setErrorCallback(ErrorCallback callback) {
    onError_ = std::move(callback);
}

void NanoClient::login(const LoginRequest& request, ResponseCallback<LoginResponse> callback) {
    this->request("AuthComponent.Login", request, callback);
}

void NanoClient::heartbeat(const HeartBeatRequest& request, ResponseCallback<HeartBeatResponse> callback) {
    this->request("AuthComponent.HeartBeat", request, callback);
}

void NanoClient::joinRoom(const JoinRoomRequest& request, ResponseCallback<RoomInfoResponse> callback) {
    this->request("RoomComponent.Join", request, callback);
}

void NanoClient::leaveRoom(const LeaveRoomRequest& request, ResponseCallback<RoomInfoResponse> callback) {
    this->request("RoomComponent.Leave", request, callback);
}

void NanoClient::getRoomInfo(const GetRoomInfoRequest& request, ResponseCallback<RoomInfoResponse> callback) {
    this->request("RoomComponent.GetRoomInfo", request, callback);
}

void NanoClient::run() {
    if (running_ || !connected_) {
        return;
    }
    
    running_ = true;
    
    // 启动网络线程
    networkThread_ = std::thread([this]() {
        std::cout << "🔄 网络线程启动" << std::endl;
        
        while (running_ && connected_) {
            if (!receiveData()) {
                std::this_thread::sleep_for(std::chrono::milliseconds(10));
            }
        }
        
        std::cout << "🔄 网络线程退出" << std::endl;
    });
}

void NanoClient::stop() {
    if (!running_) return;
    
    std::cout << "🛑 停止客户端..." << std::endl;
    running_ = false;
    
    if (networkThread_.joinable()) {
        networkThread_.join();
    }
    
    std::cout << "✅ 客户端已停止" << std::endl;
}


bool NanoClient::sendData(const std::vector<uint8_t>& data) {
    if (!connected_) {
        if (onError_) {
            onError_("未连接到服务器");
        }
        return false;
    }
    
    try {
        int sent = socket_->sendBytes(data.data(), static_cast<int>(data.size()));
        
        if (sent == static_cast<int>(data.size())) {
            std::cout << "📤 发送数据: " << data.size() << " 字节" << std::endl;
            return true;
        } else {
            std::cerr << "⚠️ 数据发送不完整: " << sent << "/" << data.size() << std::endl;
            return false;
        }
        
    } catch (const Poco::Exception& e) {
        std::cerr << "❌ 发送数据失败: " << e.displayText() << std::endl;
        
        if (onError_) {
            onError_("发送数据失败: " + e.displayText());
        }
        
        // 发送失败可能意味着连接断开
        connected_ = false;
        return false;
    }
}

bool NanoClient::receiveData() {
    if (!connected_) {
        return false;
    }
    
    try {
        // 检查是否有数据可读
        if (!socket_->poll(Poco::Timespan(0, 10000), Poco::Net::Socket::SELECT_READ)) {
            return false; // 没有数据
        }
        
        // 接收数据
        int received = socket_->receiveBytes(receiveBuffer_.data(), 
                                           static_cast<int>(receiveBuffer_.size()));
        
        if (received > 0) {
            std::cout << "📥 接收数据: " << received << " 字节" << std::endl;
            
            // 添加到消息缓冲区
            messageBuffer_.insert(messageBuffer_.end(), 
                                receiveBuffer_.begin(), 
                                receiveBuffer_.begin() + received);
            
            // 处理完整的消息包
            processMessageBuffer();
            return true;
            
        } else if (received == 0) {
            // 连接已关闭
            std::cout << "🔌 服务器关闭了连接" << std::endl;
            connected_ = false;
            
            if (onDisconnect_) {
                onDisconnect_();
            }
            
            return false;
        }
        
    } catch (const Poco::Exception& e) {
        std::cerr << "❌ 接收数据失败: " << e.displayText() << std::endl;
        
        if (onError_) {
            onError_("接收数据失败: " + e.displayText());
        }
        
        connected_ = false;
        return false;
    }
    
    return false;
}

void NanoClient::processMessageBuffer() {
    while (messageBuffer_.size() >= 4) { // 至少需要包头
        // 尝试解析包
        auto packageResult = protocol::ProtocolHandler::decodePackage(
            std::span<const uint8_t>(messageBuffer_));
        
        if (!packageResult) {
            // 解析失败，可能数据不完整
            break;
        }
        
        auto [type, body] = packageResult.value();
        
        // 计算整个包的大小
        size_t packageSize = 4 + body.size();
        
        // 处理消息
        handleMessage(body);
        
        // 从缓冲区移除已处理的数据
        messageBuffer_.erase(messageBuffer_.begin(), 
                           messageBuffer_.begin() + packageSize);
    }
}

void NanoClient::handleMessage(const std::vector<uint8_t>& data) {
    // 解码消息层
    auto messageResult = protocol::ProtocolHandler::decodeMessage(
        std::span<const uint8_t>(data));
    
    if (!messageResult) {
        std::cerr << "❌ 消息解码失败: " << messageResult.error() << std::endl;
        return;
    }
    
    const json& message = messageResult.value();
    
    // 获取元信息
    if (message.contains("_meta")) {
        const auto& meta = message["_meta"];
        const int msgType = meta["type"];
        const std::string route = meta["route"];
        const uint32_t messageId = meta.value("messageId", 0);
        
        if (msgType == static_cast<int>(protocol::MessageType::RESPONSE)) {
            // 处理响应
            handleResponse(messageId, message);
        } else if (msgType == static_cast<int>(protocol::MessageType::PUSH)) {
            // 处理推送
            handlePush(route, message);
        }
    }
}

void NanoClient::handleResponse(uint32_t requestId, const json& data) {
    std::lock_guard<std::mutex> lock(requestMutex_);
    
    auto it = pendingRequests_.find(requestId);
    if (it != pendingRequests_.end()) {
        // 调用回调函数
        it->second(data);
        
        // 移除已处理的请求
        pendingRequests_.erase(it);
        
        std::cout << "✅ 处理响应: 请求ID " << requestId << std::endl;
    } else {
        std::cerr << "⚠️ 收到未知请求ID的响应: " << requestId << std::endl;
    }
}

void NanoClient::handlePush(const std::string& route, const json& data) {
    std::lock_guard<std::mutex> lock(pushMutex_);
    
    auto it = pushHandlers_.find(route);
    if (it != pushHandlers_.end()) {
        it->second(data);
        std::cout << "✅ 处理推送消息: " << route << std::endl;
    } else {
        std::cout << "📢 收到推送消息: " << route << " -> " << data.dump() << std::endl;
    }
}

void NanoClient::sendMessage(const std::vector<uint8_t>& data) {
    sendData(data);
}

uint32_t NanoClient::generateRequestId() {
    return nextRequestId_.fetch_add(1);
}

} // namespace fastgox
