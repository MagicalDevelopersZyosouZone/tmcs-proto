// package: tmcs_msg
// file: tmcs_error.proto

import * as jspb from "google-protobuf";

export class Error extends jspb.Message {
  getCode(): ErrorCode;
  setCode(value: ErrorCode): void;

  getMessage(): string;
  setMessage(value: string): void;

  getData(): string;
  setData(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Error.AsObject;
  static toObject(includeInstance: boolean, msg: Error): Error.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Error, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Error;
  static deserializeBinaryFromReader(message: Error, reader: jspb.BinaryReader): Error;
}

export namespace Error {
  export type AsObject = {
    code: ErrorCode,
    message: string,
    data: string,
  }
}

export enum ErrorCode {
  NONE = 0,
  INVALIDMESSAGE = 1000,
  MESSAGENOTDELIVERED = 1001,
  RECEIVERUNKNOWN = 1002,
  VERIFYERROR = 1003,
  INVALIDKEY = 1004,
}

