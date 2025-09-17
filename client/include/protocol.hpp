#pragma once

#include "types.hpp"
#include <vector>
#include <span>
#include <expected>

namespace fastgox::protocol {

using EncodeResult = std::expected<std::vector<uint8_t>, std::string>;
using DecodeResult = std::expected<json, std::string>;

class ProtocolHandler {
public:
    // 编码Package层
    static EncodeResult encodePackage(PackageType type, std::span<const uint8_t> body);
    
    // 解码Package层
    static std::expected<std::pair<PackageType, std::vector<uint8_t>>, std::string> 
        decodePackage(std::span<const uint8_t> data);
    
    // 编码Message层
    static EncodeResult encodeMessage(MessageType type, 
                                     const std::string& route, 
                                     const json& data,
                                     uint32_t requestId = 0);
    
    // 解码Message层
    static DecodeResult decodeMessage(std::span<const uint8_t> data);
    
    // 编码请求消息
    static EncodeResult encodeRequest(const std::string& route, 
                                     const json& data, 
                                     uint32_t requestId);
    
    // 编码通知消息
    static EncodeResult encodeNotify(const std::string& route, const json& data);
    
private:
    // 变长整数编码
    static std::vector<uint8_t> encodeVarint(uint32_t value);
    
    // 变长整数解码
    static std::expected<std::pair<uint32_t, size_t>, std::string> 
        decodeVarint(std::span<const uint8_t> data);
    
    // 路由编码
    static std::vector<uint8_t> encodeRoute(const std::string& route);
    
    // 路由解码
    static std::expected<std::pair<std::string, size_t>, std::string> 
        decodeRoute(std::span<const uint8_t> data, bool compressed);
};

} // namespace fastgox::protocol
