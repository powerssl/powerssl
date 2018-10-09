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

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.exportSymbol('proto.powerssl.api.v1.ListMeta', null, global);
goog.exportSymbol('proto.powerssl.api.v1.ObjectMeta', null, global);
goog.exportSymbol('proto.powerssl.api.v1.TypeMeta', null, global);

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
proto.powerssl.api.v1.ListMeta = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.powerssl.api.v1.ListMeta, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.powerssl.api.v1.ListMeta.displayName = 'proto.powerssl.api.v1.ListMeta';
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
proto.powerssl.api.v1.ListMeta.prototype.toObject = function(opt_includeInstance) {
  return proto.powerssl.api.v1.ListMeta.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.powerssl.api.v1.ListMeta} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.ListMeta.toObject = function(includeInstance, msg) {
  var f, obj = {
    resourceVersion: jspb.Message.getFieldWithDefault(msg, 1, ""),
    pb_continue: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.powerssl.api.v1.ListMeta}
 */
proto.powerssl.api.v1.ListMeta.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.powerssl.api.v1.ListMeta;
  return proto.powerssl.api.v1.ListMeta.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.powerssl.api.v1.ListMeta} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.powerssl.api.v1.ListMeta}
 */
proto.powerssl.api.v1.ListMeta.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setResourceVersion(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setContinue(value);
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
proto.powerssl.api.v1.ListMeta.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.powerssl.api.v1.ListMeta.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.powerssl.api.v1.ListMeta} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.ListMeta.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResourceVersion();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getContinue();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string resource_version = 1;
 * @return {string}
 */
proto.powerssl.api.v1.ListMeta.prototype.getResourceVersion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.powerssl.api.v1.ListMeta.prototype.setResourceVersion = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string continue = 2;
 * @return {string}
 */
proto.powerssl.api.v1.ListMeta.prototype.getContinue = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.powerssl.api.v1.ListMeta.prototype.setContinue = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
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
proto.powerssl.api.v1.ObjectMeta = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.powerssl.api.v1.ObjectMeta, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.powerssl.api.v1.ObjectMeta.displayName = 'proto.powerssl.api.v1.ObjectMeta';
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
proto.powerssl.api.v1.ObjectMeta.prototype.toObject = function(opt_includeInstance) {
  return proto.powerssl.api.v1.ObjectMeta.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.powerssl.api.v1.ObjectMeta} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.ObjectMeta.toObject = function(includeInstance, msg) {
  var f, obj = {
    resourceVersion: jspb.Message.getFieldWithDefault(msg, 1, ""),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    uid: jspb.Message.getFieldWithDefault(msg, 3, ""),
    creationTimestamp: (f = msg.getCreationTimestamp()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    deletionTimestamp: (f = msg.getDeletionTimestamp()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : []
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
 * @return {!proto.powerssl.api.v1.ObjectMeta}
 */
proto.powerssl.api.v1.ObjectMeta.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.powerssl.api.v1.ObjectMeta;
  return proto.powerssl.api.v1.ObjectMeta.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.powerssl.api.v1.ObjectMeta} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.powerssl.api.v1.ObjectMeta}
 */
proto.powerssl.api.v1.ObjectMeta.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setResourceVersion(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setUid(value);
      break;
    case 4:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCreationTimestamp(value);
      break;
    case 5:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setDeletionTimestamp(value);
      break;
    case 6:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "");
         });
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
proto.powerssl.api.v1.ObjectMeta.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.powerssl.api.v1.ObjectMeta.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.powerssl.api.v1.ObjectMeta} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.ObjectMeta.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResourceVersion();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getUid();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getCreationTimestamp();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getDeletionTimestamp();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(6, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
};


/**
 * optional string resource_version = 1;
 * @return {string}
 */
proto.powerssl.api.v1.ObjectMeta.prototype.getResourceVersion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.powerssl.api.v1.ObjectMeta.prototype.setResourceVersion = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.powerssl.api.v1.ObjectMeta.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.powerssl.api.v1.ObjectMeta.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string uid = 3;
 * @return {string}
 */
proto.powerssl.api.v1.ObjectMeta.prototype.getUid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.powerssl.api.v1.ObjectMeta.prototype.setUid = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional google.protobuf.Timestamp creation_timestamp = 4;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.powerssl.api.v1.ObjectMeta.prototype.getCreationTimestamp = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 4));
};


/** @param {?proto.google.protobuf.Timestamp|undefined} value */
proto.powerssl.api.v1.ObjectMeta.prototype.setCreationTimestamp = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.powerssl.api.v1.ObjectMeta.prototype.clearCreationTimestamp = function() {
  this.setCreationTimestamp(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.powerssl.api.v1.ObjectMeta.prototype.hasCreationTimestamp = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional google.protobuf.Timestamp deletion_timestamp = 5;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.powerssl.api.v1.ObjectMeta.prototype.getDeletionTimestamp = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 5));
};


/** @param {?proto.google.protobuf.Timestamp|undefined} value */
proto.powerssl.api.v1.ObjectMeta.prototype.setDeletionTimestamp = function(value) {
  jspb.Message.setWrapperField(this, 5, value);
};


proto.powerssl.api.v1.ObjectMeta.prototype.clearDeletionTimestamp = function() {
  this.setDeletionTimestamp(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.powerssl.api.v1.ObjectMeta.prototype.hasDeletionTimestamp = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * map<string, string> labels = 6;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.powerssl.api.v1.ObjectMeta.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 6, opt_noLazyCreate,
      null));
};


proto.powerssl.api.v1.ObjectMeta.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
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
proto.powerssl.api.v1.TypeMeta = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.powerssl.api.v1.TypeMeta, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.powerssl.api.v1.TypeMeta.displayName = 'proto.powerssl.api.v1.TypeMeta';
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
proto.powerssl.api.v1.TypeMeta.prototype.toObject = function(opt_includeInstance) {
  return proto.powerssl.api.v1.TypeMeta.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.powerssl.api.v1.TypeMeta} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.TypeMeta.toObject = function(includeInstance, msg) {
  var f, obj = {
    apiVersion: jspb.Message.getFieldWithDefault(msg, 1, ""),
    kind: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.powerssl.api.v1.TypeMeta}
 */
proto.powerssl.api.v1.TypeMeta.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.powerssl.api.v1.TypeMeta;
  return proto.powerssl.api.v1.TypeMeta.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.powerssl.api.v1.TypeMeta} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.powerssl.api.v1.TypeMeta}
 */
proto.powerssl.api.v1.TypeMeta.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setApiVersion(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setKind(value);
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
proto.powerssl.api.v1.TypeMeta.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.powerssl.api.v1.TypeMeta.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.powerssl.api.v1.TypeMeta} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.TypeMeta.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getApiVersion();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getKind();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string api_version = 1;
 * @return {string}
 */
proto.powerssl.api.v1.TypeMeta.prototype.getApiVersion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.powerssl.api.v1.TypeMeta.prototype.setApiVersion = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string kind = 2;
 * @return {string}
 */
proto.powerssl.api.v1.TypeMeta.prototype.getKind = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.powerssl.api.v1.TypeMeta.prototype.setKind = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


goog.object.extend(exports, proto.powerssl.api.v1);