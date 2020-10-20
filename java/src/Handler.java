package io.grpc.fxwatcher;

import com.google.protobuf.ByteString;

public class Handler {

  public static String reply(ByteString input) {
    return "[Java] " + input.toStringUtf8();
  }

}
