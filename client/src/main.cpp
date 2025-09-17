#include "nano_client.hpp"
#include <iostream>
#include <thread>
#include <chrono>

using namespace fastgox;

int main() {
    std::cout << "🎮 FastGox Nano C++客户端启动" << std::endl;
    std::cout << "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" << std::endl;
    
    // 创建客户端
    NanoClient client;
    
    // 设置回调函数
    client.setConnectCallback([]() {
        std::cout << "🎉 连接成功回调触发!" << std::endl;
    });
    
    client.setDisconnectCallback([]() {
        std::cout << "👋 断开连接回调触发!" << std::endl;
    });
    
    client.setErrorCallback([](const std::string& error) {
        std::cerr << "❌ 错误回调: " << error << std::endl;
    });
    
    // 连接服务器
    if (!client.connect("127.0.0.1", 3250)) {
        std::cerr << "❌ 无法连接到服务器，请确保Go服务器正在运行" << std::endl;
        return 1;
    }
    
    // 启动客户端事件循环
    client.run();
    
    // 等待连接稳定
    std::this_thread::sleep_for(std::chrono::seconds(1));
    
    // 测试1: 用户登录
    std::cout << "\n🚀 测试1: 用户登录" << std::endl;
    LoginRequest loginReq;
    loginReq.Token = "poco-client-token-12345";
    
    client.login(loginReq, [](const LoginResponse& response) {
        std::cout << "✅ 登录响应:" << std::endl;
        std::cout << "   Code: " << response.Code << std::endl;
        std::cout << "   Message: " << response.Message << std::endl;
        std::cout << "   UserID: " << response.UserID << std::endl;
        std::cout << "   Nickname: " << response.Nickname << std::endl;
    });
    
    // 等待响应
    std::this_thread::sleep_for(std::chrono::seconds(2));
    
    // 测试2: 心跳检测
    std::cout << "\n💓 测试2: 心跳检测" << std::endl;
    HeartBeatRequest heartbeatReq;
    heartbeatReq.ClientTime = std::to_string(
        std::chrono::duration_cast<std::chrono::milliseconds>(
            std::chrono::system_clock::now().time_since_epoch()
        ).count()
    );
    
    client.heartbeat(heartbeatReq, [](const HeartBeatResponse& response) {
        std::cout << "✅ 心跳响应:" << std::endl;
        std::cout << "   ServerTime: " << response.ServerTime << std::endl;
        std::cout << "   Timestamp: " << response.Timestamp << std::endl;
    });
    
    // 等待响应
    std::this_thread::sleep_for(std::chrono::seconds(2));
    
    // 测试3: 加入房间
    std::cout << "\n🏠 测试3: 加入房间" << std::endl;
    JoinRoomRequest joinReq;
    joinReq.RoomID = "poco-room-001";
    
    client.joinRoom(joinReq, [](const RoomInfoResponse& response) {
        std::cout << "✅ 加入房间响应:" << std::endl;
        std::cout << "   Code: " << response.Code << std::endl;
        std::cout << "   Message: " << response.Message << std::endl;
        std::cout << "   RoomID: " << response.RoomID << std::endl;
        std::cout << "   Players: ";
        for (const auto& player : response.Players) {
            std::cout << player << " ";
        }
        std::cout << std::endl;
        std::cout << "   MaxPlayers: " << response.MaxPlayers << std::endl;
        std::cout << "   IsActive: " << (response.IsActive ? "true" : "false") << std::endl;
    });
    
    // 等待响应
    std::this_thread::sleep_for(std::chrono::seconds(2));
    
    // 测试4: 获取房间信息
    std::cout << "\n📋 测试4: 获取房间信息" << std::endl;
    GetRoomInfoRequest roomInfoReq;
    roomInfoReq.RoomID = "poco-room-001";
    
    client.getRoomInfo(roomInfoReq, [](const RoomInfoResponse& response) {
        std::cout << "✅ 房间信息响应:" << std::endl;
        std::cout << "   Code: " << response.Code << std::endl;
        std::cout << "   Message: " << response.Message << std::endl;
        std::cout << "   RoomID: " << response.RoomID << std::endl;
        std::cout << "   Players: ";
        for (const auto& player : response.Players) {
            std::cout << player << " ";
        }
        std::cout << std::endl;
    });
    
    // 等待响应
    std::this_thread::sleep_for(std::chrono::seconds(2));
    
    // 测试5: 离开房间
    std::cout << "\n🚪 测试5: 离开房间" << std::endl;
    LeaveRoomRequest leaveReq;
    leaveReq.RoomID = "poco-room-001";
    
    client.leaveRoom(leaveReq, [](const RoomInfoResponse& response) {
        std::cout << "✅ 离开房间响应:" << std::endl;
        std::cout << "   Code: " << response.Code << std::endl;
        std::cout << "   Message: " << response.Message << std::endl;
        std::cout << "   RoomID: " << response.RoomID << std::endl;
    });
    
    // 等待响应
    std::this_thread::sleep_for(std::chrono::seconds(2));
    
    std::cout << "\n🎯 所有测试完成!" << std::endl;
    std::cout << "⏰ 等待5秒后自动退出..." << std::endl;
    std::this_thread::sleep_for(std::chrono::seconds(5));
    
    // 停止客户端
    client.stop();
    client.disconnect();
    
    std::cout << "👋 FastGox Nano C++客户端退出" << std::endl;
    return 0;
}
