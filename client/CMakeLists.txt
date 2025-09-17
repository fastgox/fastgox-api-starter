# FastGox Nano C++ Client
cmake_minimum_required(VERSION 3.28)
project(FastGoxClient 
    VERSION 1.0.0
    DESCRIPTION "FastGox Nano Framework C++ Client"
    LANGUAGES CXX
)

# 强制使用LLVM 18
set(CMAKE_CXX_COMPILER clang++)
set(CMAKE_C_COMPILER clang)

# C++23标准 (LLVM 18支持)
set(CMAKE_CXX_STANDARD 23)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

# 编译选项
set(CMAKE_CXX_FLAGS_DEBUG "-g -O0 -DDEBUG")
set(CMAKE_CXX_FLAGS_RELEASE "-O3 -DNDEBUG -march=native")

# 使用Ninja生成器
set(CMAKE_GENERATOR "Ninja" CACHE INTERNAL "" FORCE)

# 依赖库
find_package(Poco REQUIRED COMPONENTS Net Foundation)
find_package(nlohmann_json REQUIRED)

# 包含目录
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/include)

# 源文件
set(CLIENT_SOURCES
    src/nano_client.cpp
    src/protocol.cpp
    src/main.cpp
)

# 头文件
set(CLIENT_HEADERS
    include/nano_client.hpp
    include/protocol.hpp
    include/types.hpp
)

# 创建可执行文件
add_executable(fastgox_client ${CLIENT_SOURCES} ${CLIENT_HEADERS})

# 编译特性
target_compile_features(fastgox_client PRIVATE cxx_std_23)

# 编译选项 (LLVM 18优化)
target_compile_options(fastgox_client PRIVATE
    -Wall -Wextra -Wpedantic
    -Wno-unused-parameter
    -fcolor-diagnostics
    -fansi-escape-codes
    $<$<CONFIG:Debug>:-fsanitize=address -fsanitize=undefined>
    $<$<CONFIG:Release>:-flto -ffast-math>
)

# 链接选项
target_link_options(fastgox_client PRIVATE
    $<$<CONFIG:Debug>:-fsanitize=address -fsanitize=undefined>
    $<$<CONFIG:Release>:-flto>
)

# 链接库
target_link_libraries(fastgox_client PRIVATE
    Poco::Net
    Poco::Foundation
    nlohmann_json::nlohmann_json
)

# 预编译头文件
target_precompile_headers(fastgox_client PRIVATE
    <iostream>
    <string>
    <vector>
    <memory>
    <chrono>
    <functional>
    <nlohmann/json.hpp>
)

# 安装配置
install(TARGETS fastgox_client RUNTIME DESTINATION bin)
