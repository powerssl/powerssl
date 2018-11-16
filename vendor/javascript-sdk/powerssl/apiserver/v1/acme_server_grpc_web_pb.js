/**
 * @fileoverview gRPC-Web generated client stub for powerssl.apiserver.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('../../../google/api/annotations_pb.js')

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.powerssl = {};
proto.powerssl.apiserver = {};
proto.powerssl.apiserver.v1 = require('./acme_server_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.powerssl.apiserver.v1.ACMEServerServiceClient =
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
proto.powerssl.apiserver.v1.ACMEServerServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.powerssl.apiserver.v1.ACMEServerServiceClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.powerssl.apiserver.v1.ACMEServerServiceClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.apiserver.v1.CreateACMEServerRequest,
 *   !proto.powerssl.apiserver.v1.ACMEServer>}
 */
const methodInfo_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.apiserver.v1.ACMEServer,
  /** @param {!proto.powerssl.apiserver.v1.CreateACMEServerRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.apiserver.v1.ACMEServer.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.CreateACMEServerRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.apiserver.v1.ACMEServer)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.apiserver.v1.ACMEServer>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.ACMEServerServiceClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.ACMEServerService/Create',
      request,
      metadata,
      methodInfo_Create,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.CreateACMEServerRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.apiserver.v1.ACMEServer>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.ACMEServerServicePromiseClient.prototype.create =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.create(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.apiserver.v1.DeleteACMEServerRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.apiserver.v1.DeleteACMEServerRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.DeleteACMEServerRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.ACMEServerServiceClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.ACMEServerService/Delete',
      request,
      metadata,
      methodInfo_Delete,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.DeleteACMEServerRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.ACMEServerServicePromiseClient.prototype.delete =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.delete(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.apiserver.v1.GetACMEServerRequest,
 *   !proto.powerssl.apiserver.v1.ACMEServer>}
 */
const methodInfo_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.apiserver.v1.ACMEServer,
  /** @param {!proto.powerssl.apiserver.v1.GetACMEServerRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.apiserver.v1.ACMEServer.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.GetACMEServerRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.apiserver.v1.ACMEServer)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.apiserver.v1.ACMEServer>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.ACMEServerServiceClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.ACMEServerService/Get',
      request,
      metadata,
      methodInfo_Get,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.GetACMEServerRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.apiserver.v1.ACMEServer>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.ACMEServerServicePromiseClient.prototype.get =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.get(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.apiserver.v1.ListACMEServersRequest,
 *   !proto.powerssl.apiserver.v1.ListACMEServersResponse>}
 */
const methodInfo_List = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.apiserver.v1.ListACMEServersResponse,
  /** @param {!proto.powerssl.apiserver.v1.ListACMEServersRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.apiserver.v1.ListACMEServersResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.ListACMEServersRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.apiserver.v1.ListACMEServersResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.apiserver.v1.ListACMEServersResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.ACMEServerServiceClient.prototype.list =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.ACMEServerService/List',
      request,
      metadata,
      methodInfo_List,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.ListACMEServersRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.apiserver.v1.ListACMEServersResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.ACMEServerServicePromiseClient.prototype.list =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.list(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.apiserver.v1.UpdateACMEServerRequest,
 *   !proto.powerssl.apiserver.v1.ACMEServer>}
 */
const methodInfo_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.apiserver.v1.ACMEServer,
  /** @param {!proto.powerssl.apiserver.v1.UpdateACMEServerRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.apiserver.v1.ACMEServer.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.UpdateACMEServerRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.apiserver.v1.ACMEServer)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.apiserver.v1.ACMEServer>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.ACMEServerServiceClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.ACMEServerService/Update',
      request,
      metadata,
      methodInfo_Update,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.UpdateACMEServerRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.apiserver.v1.ACMEServer>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.ACMEServerServicePromiseClient.prototype.update =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.update(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.powerssl.apiserver.v1;

