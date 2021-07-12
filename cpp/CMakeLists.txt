cmake_minimum_required(VERSION 3.5.1)

project(Server)

include(cmake/common.cmake)

# Proto file
get_filename_component(fxwatcher_proto "proto/fxwatcher.proto" ABSOLUTE)
get_filename_component(fxwatcher_proto_path "${fxwatcher_proto}" PATH)

# Generated sources
set(fxwatcher_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/fxwatcher.pb.cc")
set(fxwatcher_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/fxwatcher.pb.h")
set(fxwatcher_grpc_srcs "${CMAKE_CURRENT_BINARY_DIR}/fxwatcher.grpc.pb.cc")
set(fxwatcher_grpc_hdrs "${CMAKE_CURRENT_BINARY_DIR}/fxwatcher.grpc.pb.h")
add_custom_command(
      OUTPUT "${fxwatcher_proto_srcs}" "${fxwatcher_proto_hdrs}" "${fxwatcher_grpc_srcs}" "${fxwatcher_grpc_hdrs}"
      COMMAND ${_PROTOBUF_PROTOC}
      ARGS --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
        --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
        -I "${fxwatcher_proto_path}"
        --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
        "${fxwatcher_proto}"
      DEPENDS "${fxwatcher_proto}")

# Include generated *.pb.h files
include_directories("${CMAKE_CURRENT_BINARY_DIR}")

# fxwatcher_grpc_proto
add_library(fxwatcher_grpc_proto
  ${fxwatcher_grpc_srcs}
  ${fxwatcher_grpc_hdrs}
  ${fxwatcher_proto_srcs}
  ${fxwatcher_proto_hdrs})
target_link_libraries(fxwatcher_grpc_proto
  ${_REFLECTION}
  ${_GRPC_GRPCPP}
  ${_PROTOBUF_LIBPROTOBUF})

# handler
include_directories(src)
add_library(handler
  "src/handler.h"
  "src/handler.cc")
target_link_libraries(handler)

# Targets greeter_[async_](client|server)
foreach(_target
  server)
  add_executable(${_target} "${_target}.cc")
  target_link_libraries(${_target}
    handler
    fxwatcher_grpc_proto
    ${_REFLECTION}
    ${_GRPC_GRPCPP}
    ${_PROTOBUF_LIBPROTOBUF})
endforeach()

#set(source_files handler.cc handler.h)
#add_library(src ${source_files})
