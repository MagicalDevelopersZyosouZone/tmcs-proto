// package: tmcs_msg
// file: tmcs_rpc.proto

import * as jspb from "google-protobuf";

export class RPC extends jspb.Message {
  getRpcid(): number;
  setRpcid(value: number): void;

  getCode(): RPCCode;
  setCode(value: RPCCode): void;

  getMessage(): string;
  setMessage(value: string): void;

  hasTextdata(): boolean;
  clearTextdata(): void;
  getTextdata(): string;
  setTextdata(value: string): void;

  getDataCase(): RPC.DataCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RPC.AsObject;
  static toObject(includeInstance: boolean, msg: RPC): RPC.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RPC, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RPC;
  static deserializeBinaryFromReader(message: RPC, reader: jspb.BinaryReader): RPC;
}

export namespace RPC {
  export type AsObject = {
    rpcid: number,
    code: RPCCode,
    message: string,
    textdata: string,
  }

  export enum DataCase {
    DATA_NOT_SET = 0,
    TEXTDATA = 4,
  }
}

export enum RPCCode {
  PING = 0,
  PONG = 1,
  LINKCLICKED = 2000,
  CONTACTREQUEST = 2001,
  CONTACTONLINE = 2002,
  CONTACTOFFLINE = 2003,
}

