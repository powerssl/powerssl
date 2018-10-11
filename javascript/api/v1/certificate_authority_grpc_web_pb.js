/**
 * @fileoverview gRPC-Web generated client stub for powerssl.api.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('../../google/api/annotations_pb.js')

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
const proto = {};
proto.powerssl = {};
proto.powerssl.api = {};
proto.powerssl.api.v1 = require('./certificate_authority_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.powerssl.api.v1.CertificateAuthorityServiceClient =
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
proto.powerssl.api.v1.CertificateAuthorityServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.powerssl.api.v1.CertificateAuthorityServiceClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.powerssl.api.v1.CertificateAuthorityServiceClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.api.v1.CreateCertificateAuthorityRequest,
 *   !proto.powerssl.api.v1.CertificateAuthority>}
 */
const methodInfo_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.api.v1.CertificateAuthority,
  /** @param {!proto.powerssl.api.v1.CreateCertificateAuthorityRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.api.v1.CertificateAuthority.deserializeBinary
);


/**
 * @param {!proto.powerssl.api.v1.CreateCertificateAuthorityRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.api.v1.CertificateAuthority)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.api.v1.CertificateAuthority>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.api.v1.CertificateAuthorityServiceClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.api.v1.CertificateAuthorityService/Create',
      request,
      metadata,
      methodInfo_Create,
      callback);
};


/**
 * @param {!proto.powerssl.api.v1.CreateCertificateAuthorityRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.api.v1.CertificateAuthority>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.api.v1.CertificateAuthorityServicePromiseClient.prototype.create =
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
 *   !proto.powerssl.api.v1.DeleteCertificateAuthorityRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.api.v1.DeleteCertificateAuthorityRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.api.v1.DeleteCertificateAuthorityRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.api.v1.CertificateAuthorityServiceClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.api.v1.CertificateAuthorityService/Delete',
      request,
      metadata,
      methodInfo_Delete,
      callback);
};


/**
 * @param {!proto.powerssl.api.v1.DeleteCertificateAuthorityRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.api.v1.CertificateAuthorityServicePromiseClient.prototype.delete =
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
 *   !proto.powerssl.api.v1.GetCertificateAuthorityRequest,
 *   !proto.powerssl.api.v1.CertificateAuthority>}
 */
const methodInfo_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.api.v1.CertificateAuthority,
  /** @param {!proto.powerssl.api.v1.GetCertificateAuthorityRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.api.v1.CertificateAuthority.deserializeBinary
);


/**
 * @param {!proto.powerssl.api.v1.GetCertificateAuthorityRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.api.v1.CertificateAuthority)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.api.v1.CertificateAuthority>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.api.v1.CertificateAuthorityServiceClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.api.v1.CertificateAuthorityService/Get',
      request,
      metadata,
      methodInfo_Get,
      callback);
};


/**
 * @param {!proto.powerssl.api.v1.GetCertificateAuthorityRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.api.v1.CertificateAuthority>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.api.v1.CertificateAuthorityServicePromiseClient.prototype.get =
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
 *   !proto.powerssl.api.v1.ListCertificateAuthoritiesRequest,
 *   !proto.powerssl.api.v1.ListCertificateAuthoritiesResponse>}
 */
const methodInfo_List = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.api.v1.ListCertificateAuthoritiesResponse,
  /** @param {!proto.powerssl.api.v1.ListCertificateAuthoritiesRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.api.v1.ListCertificateAuthoritiesRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.api.v1.ListCertificateAuthoritiesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.api.v1.ListCertificateAuthoritiesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.api.v1.CertificateAuthorityServiceClient.prototype.list =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.api.v1.CertificateAuthorityService/List',
      request,
      metadata,
      methodInfo_List,
      callback);
};


/**
 * @param {!proto.powerssl.api.v1.ListCertificateAuthoritiesRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.api.v1.ListCertificateAuthoritiesResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.api.v1.CertificateAuthorityServicePromiseClient.prototype.list =
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
 *   !proto.powerssl.api.v1.UpdateCertificateAuthorityRequest,
 *   !proto.powerssl.api.v1.CertificateAuthority>}
 */
const methodInfo_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.api.v1.CertificateAuthority,
  /** @param {!proto.powerssl.api.v1.UpdateCertificateAuthorityRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.api.v1.CertificateAuthority.deserializeBinary
);


/**
 * @param {!proto.powerssl.api.v1.UpdateCertificateAuthorityRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.api.v1.CertificateAuthority)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.api.v1.CertificateAuthority>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.api.v1.CertificateAuthorityServiceClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.api.v1.CertificateAuthorityService/Update',
      request,
      metadata,
      methodInfo_Update,
      callback);
};


/**
 * @param {!proto.powerssl.api.v1.UpdateCertificateAuthorityRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.api.v1.CertificateAuthority>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.api.v1.CertificateAuthorityServicePromiseClient.prototype.update =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.update(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.powerssl.api.v1;

