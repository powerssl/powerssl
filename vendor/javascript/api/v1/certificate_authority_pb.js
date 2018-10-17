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

var google_api_annotations_pb = require('../../google/api/annotations_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
goog.exportSymbol('proto.powerssl.api.v1.CertificateAuthority', null, global);
goog.exportSymbol('proto.powerssl.api.v1.CreateCertificateAuthorityRequest', null, global);
goog.exportSymbol('proto.powerssl.api.v1.DeleteCertificateAuthorityRequest', null, global);
goog.exportSymbol('proto.powerssl.api.v1.GetCertificateAuthorityRequest', null, global);
goog.exportSymbol('proto.powerssl.api.v1.ListCertificateAuthoritiesRequest', null, global);
goog.exportSymbol('proto.powerssl.api.v1.ListCertificateAuthoritiesResponse', null, global);
goog.exportSymbol('proto.powerssl.api.v1.UpdateCertificateAuthorityRequest', null, global);

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
proto.powerssl.api.v1.CertificateAuthority = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.powerssl.api.v1.CertificateAuthority, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.powerssl.api.v1.CertificateAuthority.displayName = 'proto.powerssl.api.v1.CertificateAuthority';
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
proto.powerssl.api.v1.CertificateAuthority.prototype.toObject = function(opt_includeInstance) {
  return proto.powerssl.api.v1.CertificateAuthority.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.powerssl.api.v1.CertificateAuthority} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.CertificateAuthority.toObject = function(includeInstance, msg) {
  var f, obj = {

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
 * @return {!proto.powerssl.api.v1.CertificateAuthority}
 */
proto.powerssl.api.v1.CertificateAuthority.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.powerssl.api.v1.CertificateAuthority;
  return proto.powerssl.api.v1.CertificateAuthority.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.powerssl.api.v1.CertificateAuthority} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.powerssl.api.v1.CertificateAuthority}
 */
proto.powerssl.api.v1.CertificateAuthority.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
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
proto.powerssl.api.v1.CertificateAuthority.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.powerssl.api.v1.CertificateAuthority.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.powerssl.api.v1.CertificateAuthority} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.CertificateAuthority.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
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
proto.powerssl.api.v1.CreateCertificateAuthorityRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.powerssl.api.v1.CreateCertificateAuthorityRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.powerssl.api.v1.CreateCertificateAuthorityRequest.displayName = 'proto.powerssl.api.v1.CreateCertificateAuthorityRequest';
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
proto.powerssl.api.v1.CreateCertificateAuthorityRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.powerssl.api.v1.CreateCertificateAuthorityRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.powerssl.api.v1.CreateCertificateAuthorityRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.CreateCertificateAuthorityRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    certificateAuthority: (f = msg.getCertificateAuthority()) && proto.powerssl.api.v1.CertificateAuthority.toObject(includeInstance, f)
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
 * @return {!proto.powerssl.api.v1.CreateCertificateAuthorityRequest}
 */
proto.powerssl.api.v1.CreateCertificateAuthorityRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.powerssl.api.v1.CreateCertificateAuthorityRequest;
  return proto.powerssl.api.v1.CreateCertificateAuthorityRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.powerssl.api.v1.CreateCertificateAuthorityRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.powerssl.api.v1.CreateCertificateAuthorityRequest}
 */
proto.powerssl.api.v1.CreateCertificateAuthorityRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.powerssl.api.v1.CertificateAuthority;
      reader.readMessage(value,proto.powerssl.api.v1.CertificateAuthority.deserializeBinaryFromReader);
      msg.setCertificateAuthority(value);
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
proto.powerssl.api.v1.CreateCertificateAuthorityRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.powerssl.api.v1.CreateCertificateAuthorityRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.powerssl.api.v1.CreateCertificateAuthorityRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.CreateCertificateAuthorityRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCertificateAuthority();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.powerssl.api.v1.CertificateAuthority.serializeBinaryToWriter
    );
  }
};


/**
 * optional CertificateAuthority certificate_authority = 1;
 * @return {?proto.powerssl.api.v1.CertificateAuthority}
 */
proto.powerssl.api.v1.CreateCertificateAuthorityRequest.prototype.getCertificateAuthority = function() {
  return /** @type{?proto.powerssl.api.v1.CertificateAuthority} */ (
    jspb.Message.getWrapperField(this, proto.powerssl.api.v1.CertificateAuthority, 1));
};


/** @param {?proto.powerssl.api.v1.CertificateAuthority|undefined} value */
proto.powerssl.api.v1.CreateCertificateAuthorityRequest.prototype.setCertificateAuthority = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.powerssl.api.v1.CreateCertificateAuthorityRequest.prototype.clearCertificateAuthority = function() {
  this.setCertificateAuthority(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.powerssl.api.v1.CreateCertificateAuthorityRequest.prototype.hasCertificateAuthority = function() {
  return jspb.Message.getField(this, 1) != null;
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
proto.powerssl.api.v1.DeleteCertificateAuthorityRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.powerssl.api.v1.DeleteCertificateAuthorityRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.powerssl.api.v1.DeleteCertificateAuthorityRequest.displayName = 'proto.powerssl.api.v1.DeleteCertificateAuthorityRequest';
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
proto.powerssl.api.v1.DeleteCertificateAuthorityRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.powerssl.api.v1.DeleteCertificateAuthorityRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.powerssl.api.v1.DeleteCertificateAuthorityRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.DeleteCertificateAuthorityRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.powerssl.api.v1.DeleteCertificateAuthorityRequest}
 */
proto.powerssl.api.v1.DeleteCertificateAuthorityRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.powerssl.api.v1.DeleteCertificateAuthorityRequest;
  return proto.powerssl.api.v1.DeleteCertificateAuthorityRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.powerssl.api.v1.DeleteCertificateAuthorityRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.powerssl.api.v1.DeleteCertificateAuthorityRequest}
 */
proto.powerssl.api.v1.DeleteCertificateAuthorityRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
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
proto.powerssl.api.v1.DeleteCertificateAuthorityRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.powerssl.api.v1.DeleteCertificateAuthorityRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.powerssl.api.v1.DeleteCertificateAuthorityRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.DeleteCertificateAuthorityRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.powerssl.api.v1.DeleteCertificateAuthorityRequest.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.powerssl.api.v1.DeleteCertificateAuthorityRequest.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
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
proto.powerssl.api.v1.GetCertificateAuthorityRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.powerssl.api.v1.GetCertificateAuthorityRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.powerssl.api.v1.GetCertificateAuthorityRequest.displayName = 'proto.powerssl.api.v1.GetCertificateAuthorityRequest';
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
proto.powerssl.api.v1.GetCertificateAuthorityRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.powerssl.api.v1.GetCertificateAuthorityRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.powerssl.api.v1.GetCertificateAuthorityRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.GetCertificateAuthorityRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.powerssl.api.v1.GetCertificateAuthorityRequest}
 */
proto.powerssl.api.v1.GetCertificateAuthorityRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.powerssl.api.v1.GetCertificateAuthorityRequest;
  return proto.powerssl.api.v1.GetCertificateAuthorityRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.powerssl.api.v1.GetCertificateAuthorityRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.powerssl.api.v1.GetCertificateAuthorityRequest}
 */
proto.powerssl.api.v1.GetCertificateAuthorityRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
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
proto.powerssl.api.v1.GetCertificateAuthorityRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.powerssl.api.v1.GetCertificateAuthorityRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.powerssl.api.v1.GetCertificateAuthorityRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.GetCertificateAuthorityRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.powerssl.api.v1.GetCertificateAuthorityRequest.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.powerssl.api.v1.GetCertificateAuthorityRequest.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
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
proto.powerssl.api.v1.ListCertificateAuthoritiesRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.powerssl.api.v1.ListCertificateAuthoritiesRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.displayName = 'proto.powerssl.api.v1.ListCertificateAuthoritiesRequest';
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
proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.powerssl.api.v1.ListCertificateAuthoritiesRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    pageSize: jspb.Message.getFieldWithDefault(msg, 1, 0),
    pageToken: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.powerssl.api.v1.ListCertificateAuthoritiesRequest}
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.powerssl.api.v1.ListCertificateAuthoritiesRequest;
  return proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.powerssl.api.v1.ListCertificateAuthoritiesRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.powerssl.api.v1.ListCertificateAuthoritiesRequest}
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPageSize(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPageToken(value);
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
proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.powerssl.api.v1.ListCertificateAuthoritiesRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPageSize();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getPageToken();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional int32 page_size = 1;
 * @return {number}
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.prototype.getPageSize = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.prototype.setPageSize = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string page_token = 2;
 * @return {string}
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.prototype.getPageToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.powerssl.api.v1.ListCertificateAuthoritiesRequest.prototype.setPageToken = function(value) {
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
proto.powerssl.api.v1.ListCertificateAuthoritiesResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.repeatedFields_, null);
};
goog.inherits(proto.powerssl.api.v1.ListCertificateAuthoritiesResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.displayName = 'proto.powerssl.api.v1.ListCertificateAuthoritiesResponse';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.repeatedFields_ = [1];



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
proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.powerssl.api.v1.ListCertificateAuthoritiesResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    certificateAuthoritiesList: jspb.Message.toObjectList(msg.getCertificateAuthoritiesList(),
    proto.powerssl.api.v1.CertificateAuthority.toObject, includeInstance),
    nextPageToken: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.powerssl.api.v1.ListCertificateAuthoritiesResponse}
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.powerssl.api.v1.ListCertificateAuthoritiesResponse;
  return proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.powerssl.api.v1.ListCertificateAuthoritiesResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.powerssl.api.v1.ListCertificateAuthoritiesResponse}
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.powerssl.api.v1.CertificateAuthority;
      reader.readMessage(value,proto.powerssl.api.v1.CertificateAuthority.deserializeBinaryFromReader);
      msg.addCertificateAuthorities(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setNextPageToken(value);
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
proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.powerssl.api.v1.ListCertificateAuthoritiesResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCertificateAuthoritiesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.powerssl.api.v1.CertificateAuthority.serializeBinaryToWriter
    );
  }
  f = message.getNextPageToken();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * repeated CertificateAuthority certificate_authorities = 1;
 * @return {!Array<!proto.powerssl.api.v1.CertificateAuthority>}
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.prototype.getCertificateAuthoritiesList = function() {
  return /** @type{!Array<!proto.powerssl.api.v1.CertificateAuthority>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.powerssl.api.v1.CertificateAuthority, 1));
};


/** @param {!Array<!proto.powerssl.api.v1.CertificateAuthority>} value */
proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.prototype.setCertificateAuthoritiesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.powerssl.api.v1.CertificateAuthority=} opt_value
 * @param {number=} opt_index
 * @return {!proto.powerssl.api.v1.CertificateAuthority}
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.prototype.addCertificateAuthorities = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.powerssl.api.v1.CertificateAuthority, opt_index);
};


proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.prototype.clearCertificateAuthoritiesList = function() {
  this.setCertificateAuthoritiesList([]);
};


/**
 * optional string next_page_token = 2;
 * @return {string}
 */
proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.prototype.getNextPageToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.powerssl.api.v1.ListCertificateAuthoritiesResponse.prototype.setNextPageToken = function(value) {
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
proto.powerssl.api.v1.UpdateCertificateAuthorityRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.powerssl.api.v1.UpdateCertificateAuthorityRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.displayName = 'proto.powerssl.api.v1.UpdateCertificateAuthorityRequest';
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
proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.powerssl.api.v1.UpdateCertificateAuthorityRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    certificateAuthority: (f = msg.getCertificateAuthority()) && proto.powerssl.api.v1.CertificateAuthority.toObject(includeInstance, f)
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
 * @return {!proto.powerssl.api.v1.UpdateCertificateAuthorityRequest}
 */
proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.powerssl.api.v1.UpdateCertificateAuthorityRequest;
  return proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.powerssl.api.v1.UpdateCertificateAuthorityRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.powerssl.api.v1.UpdateCertificateAuthorityRequest}
 */
proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = new proto.powerssl.api.v1.CertificateAuthority;
      reader.readMessage(value,proto.powerssl.api.v1.CertificateAuthority.deserializeBinaryFromReader);
      msg.setCertificateAuthority(value);
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
proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.powerssl.api.v1.UpdateCertificateAuthorityRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getCertificateAuthority();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.powerssl.api.v1.CertificateAuthority.serializeBinaryToWriter
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional CertificateAuthority certificate_authority = 2;
 * @return {?proto.powerssl.api.v1.CertificateAuthority}
 */
proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.prototype.getCertificateAuthority = function() {
  return /** @type{?proto.powerssl.api.v1.CertificateAuthority} */ (
    jspb.Message.getWrapperField(this, proto.powerssl.api.v1.CertificateAuthority, 2));
};


/** @param {?proto.powerssl.api.v1.CertificateAuthority|undefined} value */
proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.prototype.setCertificateAuthority = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.prototype.clearCertificateAuthority = function() {
  this.setCertificateAuthority(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.powerssl.api.v1.UpdateCertificateAuthorityRequest.prototype.hasCertificateAuthority = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.powerssl.api.v1);