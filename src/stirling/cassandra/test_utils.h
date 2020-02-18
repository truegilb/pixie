#pragma once

#include <string>

namespace pl {
namespace stirling {
namespace cass {
namespace testutils {

template <typename TOpType, size_t N>
std::string CreateCQLEvent(TOpType op, const uint8_t (&a)[N], uint16_t stream) {
  static_assert(std::is_same_v<TOpType, cass::ReqOp> || std::is_same_v<TOpType, cass::RespOp>);

  std::string_view body = CreateCharArrayView<char>(a);
  size_t length = body.length();

  std::string hdr;
  hdr.resize(9);
  hdr[0] = 0x04;           // direction + version
  hdr[1] = 0x00;           // flags
  hdr[2] = (stream >> 8);  // stream
  hdr[3] = (stream >> 0);
  hdr[4] = static_cast<uint8_t>(op);  // opcode
  hdr[5] = length >> 24;              // length
  hdr[6] = length >> 16;
  hdr[7] = length >> 8;
  hdr[8] = length >> 0;

  if (std::is_same_v<TOpType, cass::RespOp>) {
    hdr[0] = hdr[0] | 0x80;
  }

  return absl::StrCat(hdr, body);
}

static constexpr int kCQLReqOpIdx = kCQLTable.ColIndex("req_op");
static constexpr int kCQLReqBodyIdx = kCQLTable.ColIndex("req_body");
static constexpr int kCQLRespOpIdx = kCQLTable.ColIndex("resp_op");
static constexpr int kCQLRespBodyIdx = kCQLTable.ColIndex("resp_body");
static constexpr int kCQLLatencyIdx = kCQLTable.ColIndex("latency_ns");

}  // namespace testutils
}  // namespace cass
}  // namespace stirling
}  // namespace pl
