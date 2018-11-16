/**
 * @fileoverview gRPC-Web generated client stub for powerssl.controller.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var powerssl_controller_v1_meta_pb = require('../../../powerssl/controller/v1/meta_pb.js')
const proto = {};
proto.powerssl = {};
proto.powerssl.controller = {};
proto.powerssl.controller.v1 = require('./integration_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.powerssl.controller.v1.IntegrationServiceClient =
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
proto.powerssl.controller.v1.IntegrationServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.powerssl.controller.v1.IntegrationServiceClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.powerssl.controller.v1.IntegrationServiceClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.RegisterIntegrationRequest,
 *   !proto.powerssl.controller.v1.Activity>}
 */
const methodInfo_Register = new grpc.web.AbstractClientBase.MethodInfo(
  powerssl_controller_v1_meta_pb.Activity,
  /** @param {!proto.powerssl.controller.v1.RegisterIntegrationRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  powerssl_controller_v1_meta_pb.Activity.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.RegisterIntegrationRequest} request The request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.Activity>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.IntegrationServiceClient.prototype.register =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/powerssl.controller.v1.IntegrationService/Register',
      request,
      metadata,
      methodInfo_Register);
};


module.exports = proto.powerssl.controller.v1;

