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

var google_protobuf_field_mask_pb = require('google-protobuf/google/protobuf/field_mask_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.powerssl = {};
proto.powerssl.apiserver = {};
proto.powerssl.apiserver.v1 = require('./user_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.powerssl.apiserver.v1.UserServiceClient =
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
proto.powerssl.apiserver.v1.UserServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.powerssl.apiserver.v1.UserServiceClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.powerssl.apiserver.v1.UserServiceClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.apiserver.v1.CreateUserRequest,
 *   !proto.powerssl.apiserver.v1.User>}
 */
const methodInfo_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.apiserver.v1.User,
  /** @param {!proto.powerssl.apiserver.v1.CreateUserRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.apiserver.v1.User.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.CreateUserRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.apiserver.v1.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.apiserver.v1.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.UserServiceClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.UserService/Create',
      request,
      metadata,
      methodInfo_Create,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.CreateUserRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.apiserver.v1.User>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.UserServicePromiseClient.prototype.create =
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
 *   !proto.powerssl.apiserver.v1.DeleteUserRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.apiserver.v1.DeleteUserRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.DeleteUserRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.UserServiceClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.UserService/Delete',
      request,
      metadata,
      methodInfo_Delete,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.DeleteUserRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.UserServicePromiseClient.prototype.delete =
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
 *   !proto.powerssl.apiserver.v1.GetUserRequest,
 *   !proto.powerssl.apiserver.v1.User>}
 */
const methodInfo_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.apiserver.v1.User,
  /** @param {!proto.powerssl.apiserver.v1.GetUserRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.apiserver.v1.User.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.GetUserRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.apiserver.v1.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.apiserver.v1.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.UserServiceClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.UserService/Get',
      request,
      metadata,
      methodInfo_Get,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.GetUserRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.apiserver.v1.User>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.UserServicePromiseClient.prototype.get =
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
 *   !proto.powerssl.apiserver.v1.ListUsersRequest,
 *   !proto.powerssl.apiserver.v1.ListUsersResponse>}
 */
const methodInfo_List = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.apiserver.v1.ListUsersResponse,
  /** @param {!proto.powerssl.apiserver.v1.ListUsersRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.apiserver.v1.ListUsersResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.ListUsersRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.apiserver.v1.ListUsersResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.apiserver.v1.ListUsersResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.UserServiceClient.prototype.list =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.UserService/List',
      request,
      metadata,
      methodInfo_List,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.ListUsersRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.apiserver.v1.ListUsersResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.UserServicePromiseClient.prototype.list =
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
 *   !proto.powerssl.apiserver.v1.UpdateUserRequest,
 *   !proto.powerssl.apiserver.v1.User>}
 */
const methodInfo_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.apiserver.v1.User,
  /** @param {!proto.powerssl.apiserver.v1.UpdateUserRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.apiserver.v1.User.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.UpdateUserRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.apiserver.v1.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.apiserver.v1.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.UserServiceClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.UserService/Update',
      request,
      metadata,
      methodInfo_Update,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.UpdateUserRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.apiserver.v1.User>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.UserServicePromiseClient.prototype.update =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.update(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.powerssl.apiserver.v1;

