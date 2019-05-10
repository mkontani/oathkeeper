/**
 * ORY Oathkeeper
 * ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.OryOathkeeper) {
      root.OryOathkeeper = {};
    }
    root.OryOathkeeper.SwaggerRuleHandler = factory(root.OryOathkeeper.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The SwaggerRuleHandler model module.
   * @module model/SwaggerRuleHandler
   * @version Latest
   */

  /**
   * Constructs a new <code>SwaggerRuleHandler</code>.
   * SwaggerRuleHandler swagger rule handler
   * @alias module:model/SwaggerRuleHandler
   * @class
   */
  var exports = function() {
    var _this = this;



  };

  /**
   * Constructs a <code>SwaggerRuleHandler</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/SwaggerRuleHandler} obj Optional instance to populate.
   * @return {module:model/SwaggerRuleHandler} The populated <code>SwaggerRuleHandler</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('config')) {
        obj['config'] = ApiClient.convertToType(data['config'], Object);
      }
      if (data.hasOwnProperty('handler')) {
        obj['handler'] = ApiClient.convertToType(data['handler'], 'String');
      }
    }
    return obj;
  }

  /**
   * Config contains the configuration for the handler. Please read the user guide for a complete list of each handler's available settings.
   * @member {Object} config
   */
  exports.prototype['config'] = undefined;
  /**
   * Handler identifies the implementation which will be used to handle this specific request. Please read the user guide for a complete list of available handlers.
   * @member {String} handler
   */
  exports.prototype['handler'] = undefined;



  return exports;
}));


