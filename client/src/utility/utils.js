'use strict';

var assign = require('object-assign');

/* Utility function to parse server message and get key value pairs.
 *  @param msg string
 *    e.g. "data:registered:[size=7]:[life=5]"
 *         "error:some error msg"
 *
 *  @return object
 *    {
 *      type: "registered",
 *      payload: {
 *        size: 7,
 *        life: 5
 *      }
 *    }
 *
 *    or,
 *
 *    {
 *      error: "some error msg"
 *    }
 */

exports.getObject = function(msg) {

  if (!msg || !msg.length) {
    return null;
  }

  var obj = {};

  msg = msg.split(':');

  /* parse error response */
  if (msg[0] === "error") {
    obj[msg[0]] = msg[1];
    return obj;
  }

  obj['type'] = msg[1];
  msg.shift();
  msg.shift();

  obj.payload = {};

  msg.forEach(function(e) {
    e = e.replace(/(\[|\])/g, '');

    var key = e.substr(0, e.indexOf('='));
    var value = e.substr(e.indexOf('=')+1, e.length);

    obj.payload[key] = value;
  });

  return obj;
}

exports.startTimer = function(duration, display) {
  if(!display) 
    display = document.querySelector('.timer');
  var start = Date.now(),
      diff,
      minutes,
      seconds;
  function timer() {
      // get the number of seconds that have elapsed since 
      // startTimer() was called
      diff = duration - (((Date.now() - start) / 1000) | 0);

      // does the same job as parseInt truncates the float
      minutes = (diff / 60) | 0;
      seconds = (diff % 60) | 0;

      minutes = minutes < 10 ? "0" + minutes : minutes;
      seconds = seconds < 10 ? "0" + seconds : seconds;

      display.textContent = minutes + ":" + seconds; 

      if (diff <= 0) {
          // add one second so that the count down starts at the full duration
          // example 05:00 not 04:59
          start = Date.now() + 1000;
      }
  };
  // we don't want to wait a full second before the timer starts
  timer();
  setInterval(timer, 1000);
}