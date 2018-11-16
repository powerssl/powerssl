/**
 * @fileoverview gRPC-Web generated client stub for powerssl.controller.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var powerssl_controller_v1_meta_pb = require('../../../powerssl/controller/v1/meta_pb.js')
const proto = {};
proto.powerssl = {};
proto.powerssl.controller = {};
proto.powerssl.controller.v1 = require('./acme_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.powerssl.controller.v1.ACMEServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.powerssl.controller.v1.ACMEServiceClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.powerssl.controller.v1.ACMEServiceClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetCreateAccountRequestResponse>}
 */
const methodInfo_GetCreateAccountRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetCreateAccountRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetCreateAccountRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetCreateAccountRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetCreateAccountRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getCreateAccountRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetCreateAccountRequest',
      request,
      metadata,
      methodInfo_GetCreateAccountRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetCreateAccountRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getCreateAccountRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getCreateAccountRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetCreateAccountResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetCreateAccountResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetCreateAccountResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetCreateAccountResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setCreateAccountResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetCreateAccountResponse',
      request,
      metadata,
      methodInfo_SetCreateAccountResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetCreateAccountResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setCreateAccountResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setCreateAccountResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetDeactivateAccountRequestResponse>}
 */
const methodInfo_GetDeactivateAccountRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetDeactivateAccountRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetDeactivateAccountRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetDeactivateAccountRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetDeactivateAccountRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getDeactivateAccountRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetDeactivateAccountRequest',
      request,
      metadata,
      methodInfo_GetDeactivateAccountRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetDeactivateAccountRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getDeactivateAccountRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getDeactivateAccountRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetDeactivateAccountResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetDeactivateAccountResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetDeactivateAccountResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetDeactivateAccountResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setDeactivateAccountResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetDeactivateAccountResponse',
      request,
      metadata,
      methodInfo_SetDeactivateAccountResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetDeactivateAccountResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setDeactivateAccountResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setDeactivateAccountResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetRekeyAccountRequestResponse>}
 */
const methodInfo_GetRekeyAccountRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetRekeyAccountRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetRekeyAccountRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetRekeyAccountRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetRekeyAccountRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getRekeyAccountRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetRekeyAccountRequest',
      request,
      metadata,
      methodInfo_GetRekeyAccountRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetRekeyAccountRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getRekeyAccountRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getRekeyAccountRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetRekeyAccountResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetRekeyAccountResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetRekeyAccountResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetRekeyAccountResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setRekeyAccountResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetRekeyAccountResponse',
      request,
      metadata,
      methodInfo_SetRekeyAccountResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetRekeyAccountResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setRekeyAccountResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setRekeyAccountResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetUpdateAccountRequestResponse>}
 */
const methodInfo_GetUpdateAccountRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetUpdateAccountRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetUpdateAccountRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetUpdateAccountRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetUpdateAccountRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getUpdateAccountRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetUpdateAccountRequest',
      request,
      metadata,
      methodInfo_GetUpdateAccountRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetUpdateAccountRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getUpdateAccountRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getUpdateAccountRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetUpdateAccountResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetUpdateAccountResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetUpdateAccountResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetUpdateAccountResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setUpdateAccountResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetUpdateAccountResponse',
      request,
      metadata,
      methodInfo_SetUpdateAccountResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetUpdateAccountResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setUpdateAccountResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setUpdateAccountResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetCreateOrderRequestResponse>}
 */
const methodInfo_GetCreateOrderRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetCreateOrderRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetCreateOrderRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetCreateOrderRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetCreateOrderRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getCreateOrderRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetCreateOrderRequest',
      request,
      metadata,
      methodInfo_GetCreateOrderRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetCreateOrderRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getCreateOrderRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getCreateOrderRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetCreateOrderResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetCreateOrderResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetCreateOrderResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetCreateOrderResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setCreateOrderResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetCreateOrderResponse',
      request,
      metadata,
      methodInfo_SetCreateOrderResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetCreateOrderResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setCreateOrderResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setCreateOrderResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetFinalizeOrderRequestResponse>}
 */
const methodInfo_GetFinalizeOrderRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetFinalizeOrderRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetFinalizeOrderRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetFinalizeOrderRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetFinalizeOrderRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getFinalizeOrderRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetFinalizeOrderRequest',
      request,
      metadata,
      methodInfo_GetFinalizeOrderRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetFinalizeOrderRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getFinalizeOrderRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getFinalizeOrderRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetFinalizeOrderResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetFinalizeOrderResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetFinalizeOrderResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetFinalizeOrderResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setFinalizeOrderResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetFinalizeOrderResponse',
      request,
      metadata,
      methodInfo_SetFinalizeOrderResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetFinalizeOrderResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setFinalizeOrderResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setFinalizeOrderResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetGetOrderRequestResponse>}
 */
const methodInfo_GetGetOrderRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetGetOrderRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetGetOrderRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetGetOrderRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetGetOrderRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getGetOrderRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetGetOrderRequest',
      request,
      metadata,
      methodInfo_GetGetOrderRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetGetOrderRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getGetOrderRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getGetOrderRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetGetOrderResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetGetOrderResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetGetOrderResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetGetOrderResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setGetOrderResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetGetOrderResponse',
      request,
      metadata,
      methodInfo_SetGetOrderResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetGetOrderResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setGetOrderResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setGetOrderResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetCreateAuthorizationRequestResponse>}
 */
const methodInfo_GetCreateAuthorizationRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetCreateAuthorizationRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetCreateAuthorizationRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetCreateAuthorizationRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetCreateAuthorizationRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getCreateAuthorizationRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetCreateAuthorizationRequest',
      request,
      metadata,
      methodInfo_GetCreateAuthorizationRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetCreateAuthorizationRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getCreateAuthorizationRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getCreateAuthorizationRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetCreateAuthorizationResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetCreateAuthorizationResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetCreateAuthorizationResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetCreateAuthorizationResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setCreateAuthorizationResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetCreateAuthorizationResponse',
      request,
      metadata,
      methodInfo_SetCreateAuthorizationResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetCreateAuthorizationResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setCreateAuthorizationResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setCreateAuthorizationResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetDeactivateAuthorizationRequestResponse>}
 */
const methodInfo_GetDeactivateAuthorizationRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetDeactivateAuthorizationRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetDeactivateAuthorizationRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetDeactivateAuthorizationRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetDeactivateAuthorizationRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getDeactivateAuthorizationRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetDeactivateAuthorizationRequest',
      request,
      metadata,
      methodInfo_GetDeactivateAuthorizationRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetDeactivateAuthorizationRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getDeactivateAuthorizationRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getDeactivateAuthorizationRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetDeactivateAuthorizationResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetDeactivateAuthorizationResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetDeactivateAuthorizationResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetDeactivateAuthorizationResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setDeactivateAuthorizationResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetDeactivateAuthorizationResponse',
      request,
      metadata,
      methodInfo_SetDeactivateAuthorizationResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetDeactivateAuthorizationResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setDeactivateAuthorizationResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setDeactivateAuthorizationResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetGetAuthorizationRequestResponse>}
 */
const methodInfo_GetGetAuthorizationRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetGetAuthorizationRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetGetAuthorizationRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetGetAuthorizationRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetGetAuthorizationRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getGetAuthorizationRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetGetAuthorizationRequest',
      request,
      metadata,
      methodInfo_GetGetAuthorizationRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetGetAuthorizationRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getGetAuthorizationRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getGetAuthorizationRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetGetAuthorizationResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetGetAuthorizationResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetGetAuthorizationResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetGetAuthorizationResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setGetAuthorizationResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetGetAuthorizationResponse',
      request,
      metadata,
      methodInfo_SetGetAuthorizationResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetGetAuthorizationResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setGetAuthorizationResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setGetAuthorizationResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetGetChallengeRequestResponse>}
 */
const methodInfo_GetGetChallengeRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetGetChallengeRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetGetChallengeRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetGetChallengeRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetGetChallengeRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getGetChallengeRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetGetChallengeRequest',
      request,
      metadata,
      methodInfo_GetGetChallengeRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetGetChallengeRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getGetChallengeRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getGetChallengeRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetGetChallengeResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetGetChallengeResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetGetChallengeResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetGetChallengeResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setGetChallengeResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetGetChallengeResponse',
      request,
      metadata,
      methodInfo_SetGetChallengeResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetGetChallengeResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setGetChallengeResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setGetChallengeResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetValidateChallengeRequestResponse>}
 */
const methodInfo_GetValidateChallengeRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetValidateChallengeRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetValidateChallengeRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetValidateChallengeRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetValidateChallengeRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getValidateChallengeRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetValidateChallengeRequest',
      request,
      metadata,
      methodInfo_GetValidateChallengeRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetValidateChallengeRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getValidateChallengeRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getValidateChallengeRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetValidateChallengeResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetValidateChallengeResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetValidateChallengeResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetValidateChallengeResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setValidateChallengeResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetValidateChallengeResponse',
      request,
      metadata,
      methodInfo_SetValidateChallengeResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetValidateChallengeResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setValidateChallengeResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setValidateChallengeResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetGetCertificateRequestResponse>}
 */
const methodInfo_GetGetCertificateRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetGetCertificateRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetGetCertificateRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetGetCertificateRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetGetCertificateRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getGetCertificateRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetGetCertificateRequest',
      request,
      metadata,
      methodInfo_GetGetCertificateRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetGetCertificateRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getGetCertificateRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getGetCertificateRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetGetCertificateResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetGetCertificateResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetGetCertificateResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetGetCertificateResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setGetCertificateResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetGetCertificateResponse',
      request,
      metadata,
      methodInfo_SetGetCertificateResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetGetCertificateResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setGetCertificateResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setGetCertificateResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.Activity,
 *   !proto.powerssl.controller.v1.GetRevokeCertificateRequestResponse>}
 */
const methodInfo_GetRevokeCertificateRequest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.GetRevokeCertificateRequestResponse,
  /** @param {!proto.powerssl.controller.v1.Activity} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.GetRevokeCertificateRequestResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.GetRevokeCertificateRequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.GetRevokeCertificateRequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.getRevokeCertificateRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/GetRevokeCertificateRequest',
      request,
      metadata,
      methodInfo_GetRevokeCertificateRequest,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.Activity} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.GetRevokeCertificateRequestResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.getRevokeCertificateRequest =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getRevokeCertificateRequest(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.SetRevokeCertificateResponseRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_SetRevokeCertificateResponse = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.controller.v1.SetRevokeCertificateResponseRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.SetRevokeCertificateResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServiceClient.prototype.setRevokeCertificateResponse =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.ACMEService/SetRevokeCertificateResponse',
      request,
      metadata,
      methodInfo_SetRevokeCertificateResponse,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.SetRevokeCertificateResponseRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.ACMEServicePromiseClient.prototype.setRevokeCertificateResponse =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.setRevokeCertificateResponse(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.powerssl.controller.v1;

