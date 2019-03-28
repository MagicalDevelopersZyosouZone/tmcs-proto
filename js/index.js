"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var tmcs_msg_pb_1 = __importDefault(require("./tmcs_msg_pb"));
exports.TMCSMsg = tmcs_msg_pb_1.default;
var tmcs_error_pb_1 = __importDefault(require("./tmcs_error_pb"));
exports.TMCSError = tmcs_error_pb_1.default;
var tmcs_rpc_pb_1 = __importDefault(require("./tmcs_rpc_pb"));
exports.TMCSRPC = tmcs_rpc_pb_1.default;
exports.default = {
    TMCSMsg: tmcs_msg_pb_1.default,
    TMCSError: tmcs_error_pb_1.default,
    TMCSRPC: tmcs_rpc_pb_1.default,
};
