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
proto.powerssl.apiserver.v1 = require('./certificate_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.powerssl.apiserver.v1.CertificateServiceClient =
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
proto.powerssl.apiserver.v1.CertificateServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.powerssl.apiserver.v1.CertificateServiceClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.powerssl.apiserver.v1.CertificateServiceClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.apiserver.v1.CreateCertificateRequest,
 *   !proto.powerssl.apiserver.v1.Certificate>}
 */
const methodInfo_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.apiserver.v1.Certificate,
  /** @param {!proto.powerssl.apiserver.v1.CreateCertificateRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.apiserver.v1.Certificate.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.CreateCertificateRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.apiserver.v1.Certificate)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.apiserver.v1.Certificate>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.CertificateServiceClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.CertificateService/Create',
      request,
      metadata,
      methodInfo_Create,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.CreateCertificateRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.apiserver.v1.Certificate>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.CertificateServicePromiseClient.prototype.create =
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
 *   !proto.powerssl.apiserver.v1.DeleteCertificateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.powerssl.apiserver.v1.DeleteCertificateRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.DeleteCertificateRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.CertificateServiceClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.CertificateService/Delete',
      request,
      metadata,
      methodInfo_Delete,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.DeleteCertificateRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.CertificateServicePromiseClient.prototype.delete =
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
 *   !proto.powerssl.apiserver.v1.GetCertificateRequest,
 *   !proto.powerssl.apiserver.v1.Certificate>}
 */
const methodInfo_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.apiserver.v1.Certificate,
  /** @param {!proto.powerssl.apiserver.v1.GetCertificateRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.apiserver.v1.Certificate.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.GetCertificateRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.apiserver.v1.Certificate)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.apiserver.v1.Certificate>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.CertificateServiceClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.CertificateService/Get',
      request,
      metadata,
      methodInfo_Get,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.GetCertificateRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.apiserver.v1.Certificate>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.CertificateServicePromiseClient.prototype.get =
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
 *   !proto.powerssl.apiserver.v1.ListCertificatesRequest,
 *   !proto.powerssl.apiserver.v1.ListCertificatesResponse>}
 */
const methodInfo_List = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.apiserver.v1.ListCertificatesResponse,
  /** @param {!proto.powerssl.apiserver.v1.ListCertificatesRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.apiserver.v1.ListCertificatesResponse.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.ListCertificatesRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.apiserver.v1.ListCertificatesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.apiserver.v1.ListCertificatesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.CertificateServiceClient.prototype.list =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.CertificateService/List',
      request,
      metadata,
      methodInfo_List,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.ListCertificatesRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.apiserver.v1.ListCertificatesResponse>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.CertificateServicePromiseClient.prototype.list =
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
 *   !proto.powerssl.apiserver.v1.UpdateCertificateRequest,
 *   !proto.powerssl.apiserver.v1.Certificate>}
 */
const methodInfo_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.apiserver.v1.Certificate,
  /** @param {!proto.powerssl.apiserver.v1.UpdateCertificateRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.apiserver.v1.Certificate.deserializeBinary
);


/**
 * @param {!proto.powerssl.apiserver.v1.UpdateCertificateRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.apiserver.v1.Certificate)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.apiserver.v1.Certificate>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.CertificateServiceClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.apiserver.v1.CertificateService/Update',
      request,
      metadata,
      methodInfo_Update,
      callback);
};


/**
 * @param {!proto.powerssl.apiserver.v1.UpdateCertificateRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.apiserver.v1.Certificate>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.apiserver.v1.CertificateServicePromiseClient.prototype.update =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.update(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.powerssl.apiserver.v1;

