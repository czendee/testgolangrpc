// Copyright 2015, Google Inc.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

syntax = "proto3";

//option go_package = "google.golang.org/grpc/examples/multiplica/multiplica";
option go_package = "./";
option java_multiple_files = true;
option java_package = "io.grpc.examples.multiplica";
option java_outer_classname = "MultiplicaProto";

package multiplica;

import "google/api/annotations.proto";


// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayMultiplica (MultiplicaRequest) returns (MultiplicaReply) {
    option (google.api.http) = {
      post: "/v1/multiplica/greeter/say-multiplica"
      body: "*"
    };
  }
}

// The request message containing the user's name.
message MultiplicaRequest {
  string name = 1;
  string numero1 = 2;
  string numero2 = 3;
}

// The response message containing the greetings
message MultiplicaReply {
  string message = 1;
}

// message Info {
//   string version = 1;
//   string title = 2;
// }
// 
// message Schema {
//   string swagger = 1;
//   Info info = 2;
//   repeated string schemes = 3;
//   repeated string consumes = 4;
//   repeated string produces = 5;
// }
