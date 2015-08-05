var AppDispatcher = require('../dispatchers/app_dispatcher.js');
var assign = require('object-assign');
var EventEmitter = require('events').EventEmitter;
var AppConstants = require('../constants/app_constants.js');

/*
 * Minesfield.
 *
 * Status:
 *
 *     0 : default, not clicked,
 *    -1 : visited,
 *     2 : life,
 *    -2 : death
 *
 * Size: 7 by 7
 *
 * Mines position (x, y) from array can be calculated by,:
 *    x : index of an element / size of minesfield
 *    y : index of an element % size of minesfield
 *
 *  var mines = [
 *    0, -1, 0, 0, 0, 0, 0,
 *    0, 0, 0, 2, 0, 0, -1,
 *    0, -2, 0, 0, 0, 0, 0,
 *    0, 0, 0, -1, 0, 0, 0,
 *    0, 0, 0, -1, 0, 0, 0,
 *    0, 0, 0, 0, 0, 0, -1,
 *    0, 0, 0, 0, 0, 0, -1,
 *  ];
 */

var getMatrix = function(size) {
  return Array.apply(null, Array(size * size)).map(Number.prototype.valueOf, -9);
};

/* username */
var player = {
  'username' : '',
  'life' : 6,
  'size' : 0
};

var enemy = {
  'username' : '',
  'life' : 5,
  'size' : 0
};

var CHANGE_EVENT = "change";

var MinesStore = assign({}, EventEmitter.prototype, {
  emitChange: function() {
    this.emit(CHANGE_EVENT);
  },

  addChangeListener: function(callback) {
    this.on(CHANGE_EVENT, callback);
  },

  removeChangeListener: function(callback) {
    this.removeListener(CHANGE_EVENT, callback);
  },

  'addUsername': function(name) {
    player.username = name;
  },

  'getUsername': function() {
    return player.username;
  },

  'getEnemy': function() {
    return enemy;
  },

  'getUser': function() {
    return player;
  }
});

MinesStore.dispatchToken = AppDispatcher.register(function(payload) {
  var action = payload.action;

  switch(action.type) {
    case 'registered':
      player.life = action.payload.life;
      player.size = action.payload.size;
      MinesStore.emitChange();
      break;

    case 'ready':
      /* got an opponent */
      enemy.username = action.payload.name;
      enemy.life = action.payload.life;
      enemy.size = action.payload.size;
      /* only when ready, show the minefields */
      enemy.mines = getMatrix(enemy.size);
      player.mines = getMatrix(player.size);
      MinesStore.emitChange();
      break;

    case 'open':
      console.log('open', action);
      break;

    case 'result':
      var idx = +action.payload.idx,
        life = +action.payload.life,
        type = +action.payload.type;

      player.mines[idx] = type;
      player.life = life;
      MinesStore.emitChange();
      break;

    case 'enemy':
      var idx = +action.payload.idx,
      life = +action.payload.life,
      type = +action.payload.type;

      enemy.mines[idx] = type;
      enemy.life = life;
      MinesStore.emitChange();
      break;

    default:
      // do nothing
  }
});

module.exports = MinesStore;
