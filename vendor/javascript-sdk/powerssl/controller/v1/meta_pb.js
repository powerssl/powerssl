/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.powerssl.controller.v1.Activity', null, global);
goog.exportSymbol('proto.powerssl.controller.v1.Activity.Name', null, global);
goog.exportSymbol('proto.powerssl.controller.v1.Activity.Workflow', null, global);
goog.exportSymbol('proto.powerssl.controller.v1.Error', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.powerssl.controller.v1.Activity = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.powerssl.controller.v1.Activity, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.powerssl.controller.v1.Activity.displayName = 'proto.powerssl.controller.v1.Activity';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.powerssl.controller.v1.Activity.prototype.toObject = function(opt_includeInstance) {
  return proto.powerssl.controller.v1.Activity.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.powerssl.controller.v1.Activity} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.controller.v1.Activity.toObject = function(includeInstance, msg) {
  var f, obj = {
    token: jspb.Message.getFieldWithDefault(msg, 1, ""),
    name: jspb.Message.getFieldWithDefault(msg, 2, 0),
    workflow: (f = msg.getWorkflow()) && proto.powerssl.controller.v1.Activity.Workflow.toObject(includeInstance, f),
    signature: jspb.Message.getFieldWithDefault(msg, 4, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.powerssl.controller.v1.Activity}
 */
proto.powerssl.controller.v1.Activity.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.powerssl.controller.v1.Activity;
  return proto.powerssl.controller.v1.Activity.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.powerssl.controller.v1.Activity} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.powerssl.controller.v1.Activity}
 */
proto.powerssl.controller.v1.Activity.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
      break;
    case 2:
      var value = /** @type {!proto.powerssl.controller.v1.Activity.Name} */ (reader.readEnum());
      msg.setName(value);
      break;
    case 3:
      var value = new proto.powerssl.controller.v1.Activity.Workflow;
      reader.readMessage(value,proto.powerssl.controller.v1.Activity.Workflow.deserializeBinaryFromReader);
      msg.setWorkflow(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setSignature(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.powerssl.controller.v1.Activity.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.powerssl.controller.v1.Activity.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.powerssl.controller.v1.Activity} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.controller.v1.Activity.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getName();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getWorkflow();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.powerssl.controller.v1.Activity.Workflow.serializeBinaryToWriter
    );
  }
  f = message.getSignature();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.powerssl.controller.v1.Activity.Name = {
  NAME_UNSPECIFIED: 0,
  ACME_CREATE_ACCOUNT: 101,
  ACME_CREATE_AUTHORIZATION: 102,
  ACME_CREATE_ORDER: 103,
  ACME_DEACTIVATE_ACCOUNT: 104,
  ACME_DEACTIVATE_AUTHORIZATION: 105,
  ACME_FINALIZE_ORDER: 106,
  ACME_GET_AUTHORIZATION: 107,
  ACME_GET_CERTIFICATE: 108,
  ACME_GET_CHALLENGE: 109,
  ACME_GET_ORDER: 110,
  ACME_REKEY_ACCOUNT: 111,
  ACME_REVOKE_CERTIFICATE: 112,
  ACME_UPDATE_ACCOUNT: 113,
  ACME_VALIDATE_CHALLENGE: 114,
  DNS_CREATE_RECORD: 201,
  DNS_DELETE_RECORD: 202,
  DNS_VERIFY_DOMAIN: 203
};


/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.powerssl.controller.v1.Activity.Workflow = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.powerssl.controller.v1.Activity.Workflow.repeatedFields_, null);
};
goog.inherits(proto.powerssl.controller.v1.Activity.Workflow, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.powerssl.controller.v1.Activity.Workflow.displayName = 'proto.powerssl.controller.v1.Activity.Workflow';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.powerssl.controller.v1.Activity.Workflow.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.powerssl.controller.v1.Activity.Workflow.prototype.toObject = function(opt_includeInstance) {
  return proto.powerssl.controller.v1.Activity.Workflow.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.powerssl.controller.v1.Activity.Workflow} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.controller.v1.Activity.Workflow.toObject = function(includeInstance, msg) {
  var f, obj = {
    activitiesList: jspb.Message.getRepeatedField(msg, 1)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.powerssl.controller.v1.Activity.Workflow}
 */
proto.powerssl.controller.v1.Activity.Workflow.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.powerssl.controller.v1.Activity.Workflow;
  return proto.powerssl.controller.v1.Activity.Workflow.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.powerssl.controller.v1.Activity.Workflow} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.powerssl.controller.v1.Activity.Workflow}
 */
proto.powerssl.controller.v1.Activity.Workflow.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addActivities(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.powerssl.controller.v1.Activity.Workflow.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.powerssl.controller.v1.Activity.Workflow.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.powerssl.controller.v1.Activity.Workflow} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.controller.v1.Activity.Workflow.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getActivitiesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
};


/**
 * repeated string activities = 1;
 * @return {!Array<string>}
 */
proto.powerssl.controller.v1.Activity.Workflow.prototype.getActivitiesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/** @param {!Array<string>} value */
proto.powerssl.controller.v1.Activity.Workflow.prototype.setActivitiesList = function(value) {
  jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {!string} value
 * @param {number=} opt_index
 */
proto.powerssl.controller.v1.Activity.Workflow.prototype.addActivities = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


proto.powerssl.controller.v1.Activity.Workflow.prototype.clearActivitiesList = function() {
  this.setActivitiesList([]);
};


/**
 * optional string token = 1;
 * @return {string}
 */
proto.powerssl.controller.v1.Activity.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.powerssl.controller.v1.Activity.prototype.setToken = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional Name name = 2;
 * @return {!proto.powerssl.controller.v1.Activity.Name}
 */
proto.powerssl.controller.v1.Activity.prototype.getName = function() {
  return /** @type {!proto.powerssl.controller.v1.Activity.Name} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {!proto.powerssl.controller.v1.Activity.Name} value */
proto.powerssl.controller.v1.Activity.prototype.setName = function(value) {
  jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional Workflow workflow = 3;
 * @return {?proto.powerssl.controller.v1.Activity.Workflow}
 */
proto.powerssl.controller.v1.Activity.prototype.getWorkflow = function() {
  return /** @type{?proto.powerssl.controller.v1.Activity.Workflow} */ (
    jspb.Message.getWrapperField(this, proto.powerssl.controller.v1.Activity.Workflow, 3));
};


/** @param {?proto.powerssl.controller.v1.Activity.Workflow|undefined} value */
proto.powerssl.controller.v1.Activity.prototype.setWorkflow = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.powerssl.controller.v1.Activity.prototype.clearWorkflow = function() {
  this.setWorkflow(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.powerssl.controller.v1.Activity.prototype.hasWorkflow = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional string signature = 4;
 * @return {string}
 */
proto.powerssl.controller.v1.Activity.prototype.getSignature = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.powerssl.controller.v1.Activity.prototype.setSignature = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.powerssl.controller.v1.Error = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.powerssl.controller.v1.Error, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.powerssl.controller.v1.Error.displayName = 'proto.powerssl.controller.v1.Error';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.powerssl.controller.v1.Error.prototype.toObject = function(opt_includeInstance) {
  return proto.powerssl.controller.v1.Error.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.powerssl.controller.v1.Error} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.controller.v1.Error.toObject = function(includeInstance, msg) {
  var f, obj = {
    message: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.powerssl.controller.v1.Error}
 */
proto.powerssl.controller.v1.Error.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.powerssl.controller.v1.Error;
  return proto.powerssl.controller.v1.Error.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.powerssl.controller.v1.Error} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.powerssl.controller.v1.Error}
 */
proto.powerssl.controller.v1.Error.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setMessage(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.powerssl.controller.v1.Error.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.powerssl.controller.v1.Error.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.powerssl.controller.v1.Error} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.controller.v1.Error.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMessage();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string message = 1;
 * @return {string}
 */
proto.powerssl.controller.v1.Error.prototype.getMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.powerssl.controller.v1.Error.prototype.setMessage = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


goog.object.extend(exports, proto.powerssl.controller.v1);