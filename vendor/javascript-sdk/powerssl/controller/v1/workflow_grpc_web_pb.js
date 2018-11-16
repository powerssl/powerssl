/**
 * @fileoverview gRPC-Web generated client stub for powerssl.controller.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var powerssl_controller_v1_integration_pb = require('../../../powerssl/controller/v1/integration_pb.js')
const proto = {};
proto.powerssl = {};
proto.powerssl.controller = {};
proto.powerssl.controller.v1 = require('./workflow_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.powerssl.controller.v1.WorkflowServiceClient =
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
proto.powerssl.controller.v1.WorkflowServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.powerssl.controller.v1.WorkflowServiceClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.powerssl.controller.v1.WorkflowServiceClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerssl.controller.v1.CreateWorkflowRequest,
 *   !proto.powerssl.controller.v1.Workflow>}
 */
const methodInfo_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerssl.controller.v1.Workflow,
  /** @param {!proto.powerssl.controller.v1.CreateWorkflowRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerssl.controller.v1.Workflow.deserializeBinary
);


/**
 * @param {!proto.powerssl.controller.v1.CreateWorkflowRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerssl.controller.v1.Workflow)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerssl.controller.v1.Workflow>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.WorkflowServiceClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerssl.controller.v1.WorkflowService/Create',
      request,
      metadata,
      methodInfo_Create,
      callback);
};


/**
 * @param {!proto.powerssl.controller.v1.CreateWorkflowRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerssl.controller.v1.Workflow>}
 *     The XHR Node Readable Stream
 */
proto.powerssl.controller.v1.WorkflowServicePromiseClient.prototype.create =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.create(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.powerssl.controller.v1;

