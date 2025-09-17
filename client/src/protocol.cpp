#include "protocol.hpp"
#include <iostream>
#include <bit>

namespace fastgox::protocol {

// 编码Package层
EncodeResult ProtocolHandler::encodePackage(PackageType type, std::span<const uint8_t> body) {
    std::vector<uint8_t> package;
    package.reserve(PKG_HEAD_BYTES + body.size());
    
    // Type (1 byte)
    package.push_back(static_cast<uint8_t>(type));
    
    // Length (3 bytes, big-endian)
    const uint32_t length = static_cast<uint32_t>(body.size());
    if (length > 0xFFFFFF) {
        return std::unexpected("Package body too large");
    }
    
    package.push_back((length >> 16) & 0xFF);
    package.push_back((length >> 8) & 0xFF);
    package.push_back(length & 0xFF);
    
    // Body
    package.insert(package.end(), body.begin(), body.end());
    
    return package;
}

// 解码Package层
std::expected<std::pair<PackageType, std::vector<uint8_t>>, std::string> 
ProtocolHandler::decodePackage(std::span<const uint8_t> data) {
    if (data.size() < PKG_HEAD_BYTES) {
        return std::unexpected("Package data too short");
    }
    
    // 解析Type
    const auto type = static_cast<PackageType>(data[0]);
    
    // 解析Length (big-endian)
    const uint32_t length = (static_cast<uint32_t>(data[1]) << 16) |
                           (static_cast<uint32_t>(data[2]) << 8) |
                           static_cast<uint32_t>(data[3]);
    
    if (data.size() < PKG_HEAD_BYTES + length) {
        return std::unexpected("Package data incomplete");
    }
    
    // 提取Body
    std::vector<uint8_t> body(data.begin() + PKG_HEAD_BYTES, 
                             data.begin() + PKG_HEAD_BYTES + length);
    
    return std::make_pair(type, std::move(body));
}

// 编码Message层
EncodeResult ProtocolHandler::encodeMessage(MessageType type, 
                                           const std::string& route, 
                                           const json& data,
                                           uint32_t requestId) {
    std::vector<uint8_t> message;
    
    // Flag byte: [type:3bits][compress:1bit][reserved:4bits]
    uint8_t flag = (static_cast<uint8_t>(type) & MSG_TYPE_MASK);
    message.push_back(flag);
    
    // Message ID (仅对Request/Response类型)
    if (type == MessageType::REQUEST || type == MessageType::RESPONSE) {
        auto idBytes = encodeVarint(requestId);
        message.insert(message.end(), idBytes.begin(), idBytes.end());
    }
    
    // Route
    auto routeBytes = encodeRoute(route);
    message.insert(message.end(), routeBytes.begin(), routeBytes.end());
    
    // Data (JSON)
    const std::string jsonStr = data.dump();
    const auto jsonBytes = std::as_bytes(std::span{jsonStr});
    message.insert(message.end(), 
                  reinterpret_cast<const uint8_t*>(jsonBytes.data()),
                  reinterpret_cast<const uint8_t*>(jsonBytes.data()) + jsonBytes.size());
    
    return message;
}

// 解码Message层
DecodeResult ProtocolHandler::decodeMessage(std::span<const uint8_t> data) {
    if (data.empty()) {
        return std::unexpected("Empty message data");
    }
    
    size_t offset = 0;
    
    // 解析Flag
    const uint8_t flag = data[offset++];
    const auto msgType = static_cast<MessageType>(flag & MSG_TYPE_MASK);
    const bool compressed = (flag & MSG_COMPRESS_ROUTE_MASK) != 0;
    
    // 解析Message ID
    uint32_t messageId = 0;
    if (msgType == MessageType::REQUEST || msgType == MessageType::RESPONSE) {
        auto idResult = decodeVarint(data.subspan(offset));
        if (!idResult) {
            return std::unexpected("Failed to decode message ID: " + idResult.error());
        }
        messageId = idResult->first;
        offset += idResult->second;
    }
    
    // 解析Route
    auto routeResult = decodeRoute(data.subspan(offset), compressed);
    if (!routeResult) {
        return std::unexpected("Failed to decode route: " + routeResult.error());
    }
    const std::string route = routeResult->first;
    offset += routeResult->second;
    
    // 解析JSON数据
    if (offset >= data.size()) {
        return std::unexpected("No JSON data in message");
    }
    
    try {
        const std::string jsonStr(reinterpret_cast<const char*>(data.data() + offset),
                                 data.size() - offset);
        json result = json::parse(jsonStr);
        
        // 添加元信息
        result["_meta"] = {
            {"type", static_cast<int>(msgType)},
            {"route", route},
            {"messageId", messageId}
        };
        
        return result;
    } catch (const json::exception& e) {
        return std::unexpected("JSON parse error: " + std::string(e.what()));
    }
}

// 编码请求消息
EncodeResult ProtocolHandler::encodeRequest(const std::string& route, 
                                           const json& data, 
                                           uint32_t requestId) {
    auto messageBytes = encodeMessage(MessageType::REQUEST, route, data, requestId);
    if (!messageBytes) {
        return messageBytes;
    }
    
    return encodePackage(PackageType::DATA, messageBytes.value());
}

// 编码通知消息
EncodeResult ProtocolHandler::encodeNotify(const std::string& route, const json& data) {
    auto messageBytes = encodeMessage(MessageType::NOTIFY, route, data);
    if (!messageBytes) {
        return messageBytes;
    }
    
    return encodePackage(PackageType::DATA, messageBytes.value());
}

// 变长整数编码
std::vector<uint8_t> ProtocolHandler::encodeVarint(uint32_t value) {
    std::vector<uint8_t> result;
    
    while (value >= 0x80) {
        result.push_back(static_cast<uint8_t>((value & 0x7F) | 0x80));
        value >>= 7;
    }
    result.push_back(static_cast<uint8_t>(value & 0x7F));
    
    return result;
}

// 变长整数解码
std::expected<std::pair<uint32_t, size_t>, std::string> 
ProtocolHandler::decodeVarint(std::span<const uint8_t> data) {
    uint32_t result = 0;
    size_t bytesRead = 0;
    int shift = 0;
    
    for (uint8_t byte : data) {
        if (shift >= 32) {
            return std::unexpected("Varint too large");
        }
        
        result |= static_cast<uint32_t>(byte & 0x7F) << shift;
        bytesRead++;
        
        if ((byte & 0x80) == 0) {
            return std::make_pair(result, bytesRead);
        }
        
        shift += 7;
    }
    
    return std::unexpected("Incomplete varint");
}

// 路由编码
std::vector<uint8_t> ProtocolHandler::encodeRoute(const std::string& route) {
    std::vector<uint8_t> result;
    
    // 路由长度 (1 byte)
    if (route.length() > 255) {
        result.push_back(255);
        // TODO: 支持路由压缩
    } else {
        result.push_back(static_cast<uint8_t>(route.length()));
    }
    
    // 路由内容
    const auto routeBytes = std::as_bytes(std::span{route});
    result.insert(result.end(),
                 reinterpret_cast<const uint8_t*>(routeBytes.data()),
                 reinterpret_cast<const uint8_t*>(routeBytes.data()) + routeBytes.size());
    
    return result;
}

// 路由解码
std::expected<std::pair<std::string, size_t>, std::string> 
ProtocolHandler::decodeRoute(std::span<const uint8_t> data, bool compressed) {
    if (data.empty()) {
        return std::unexpected("Empty route data");
    }
    
    if (compressed) {
        // TODO: 实现路由解压缩
        return std::unexpected("Route compression not implemented");
    }
    
    // 读取路由长度
    const uint8_t routeLength = data[0];
    if (data.size() < 1 + routeLength) {
        return std::unexpected("Route data incomplete");
    }
    
    // 读取路由内容
    const std::string route(reinterpret_cast<const char*>(data.data() + 1), routeLength);
    
    return std::make_pair(route, 1 + routeLength);
}

} // namespace fastgox::protocol
