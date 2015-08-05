var Dispatcher = require('flux').Dispatcher;
var assign = require('object-assign');

/* constants */
var AppConstants = require('../constants/app_constants.js');
var PayloadSources = AppConstants.PayloadSources;

var AppDispatcher = assign(new Dispatcher(), {
  /**
   * @param {object} action The details of the action, including the action's
   * type and additional data coming from the server.
   */
  handleServerAction: function(action) {
    var payload = {
      source: PayloadSources.SERVER_ACTION,
      action: action
    };
    this.dispatch(payload);
  }
});

module.exports = AppDispatcher;
