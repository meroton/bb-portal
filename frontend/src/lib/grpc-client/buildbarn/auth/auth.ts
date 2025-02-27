// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v3.19.1
// source: buildbarn/auth/auth.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Value } from "../../google/protobuf/struct";
import { KeyValue } from "../../opentelemetry/proto/common/v1/common";

export const protobufPackage = "buildbarn.auth";

/**
 * Protobuf equivalent of the AuthenticationMetadata structure that is
 * used by the auth framework to store information on an authenticated
 * user.
 */
export interface AuthenticationMetadata {
  /**
   * Part of the authentication metadata that is safe to display
   * publicly (e.g., as part of logs or bb_browser).
   */
  public:
    | any
    | undefined;
  /**
   * OpenTelemetry tracing attributes to add to spans in which the
   * authentication took place (e.g., gRPC server call spans). All
   * attributes will have "auth." prepended to their names
   * automatically.
   */
  tracingAttributes: KeyValue[];
  /**
   * Part of the authentication metadata that should not be displayed
   * publicly. This field is useful for propagating information from the
   * authentication layer to the authorization layer, as this data can
   * be accessed by JMESPathExpressionAuthorizer.
   */
  private: any | undefined;
}

function createBaseAuthenticationMetadata(): AuthenticationMetadata {
  return { public: undefined, tracingAttributes: [], private: undefined };
}

export const AuthenticationMetadata: MessageFns<AuthenticationMetadata> = {
  encode(message: AuthenticationMetadata, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.public !== undefined) {
      Value.encode(Value.wrap(message.public), writer.uint32(10).fork()).join();
    }
    for (const v of message.tracingAttributes) {
      KeyValue.encode(v!, writer.uint32(18).fork()).join();
    }
    if (message.private !== undefined) {
      Value.encode(Value.wrap(message.private), writer.uint32(26).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): AuthenticationMetadata {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAuthenticationMetadata();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.public = Value.unwrap(Value.decode(reader, reader.uint32()));
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.tracingAttributes.push(KeyValue.decode(reader, reader.uint32()));
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.private = Value.unwrap(Value.decode(reader, reader.uint32()));
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AuthenticationMetadata {
    return {
      public: isSet(object?.public) ? object.public : undefined,
      tracingAttributes: globalThis.Array.isArray(object?.tracingAttributes)
        ? object.tracingAttributes.map((e: any) => KeyValue.fromJSON(e))
        : [],
      private: isSet(object?.private) ? object.private : undefined,
    };
  },

  toJSON(message: AuthenticationMetadata): unknown {
    const obj: any = {};
    if (message.public !== undefined) {
      obj.public = message.public;
    }
    if (message.tracingAttributes?.length) {
      obj.tracingAttributes = message.tracingAttributes.map((e) => KeyValue.toJSON(e));
    }
    if (message.private !== undefined) {
      obj.private = message.private;
    }
    return obj;
  },

  create(base?: DeepPartial<AuthenticationMetadata>): AuthenticationMetadata {
    return AuthenticationMetadata.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<AuthenticationMetadata>): AuthenticationMetadata {
    const message = createBaseAuthenticationMetadata();
    message.public = object.public ?? undefined;
    message.tracingAttributes = object.tracingAttributes?.map((e) => KeyValue.fromPartial(e)) || [];
    message.private = object.private ?? undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}
