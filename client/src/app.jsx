'use strict';

import React from 'react';
import Router from 'react-router';
import { DefaultRoute, Route, RouteHandler } from 'react-router';

import Home from './components/Home.jsx';
import Minefield from './components/Minefield.jsx';

let App = React.createClass({
  render() {
    return (
      <RouteHandler />
    );
  }
});

let Routes = (
  <Route name="app" path="/" handler={App}>
    <Route name="home" path="/home" handler={Home}/>
    <Route name="minefield" path="/minefield" handler={Minefield}/>
    <DefaultRoute handler={Home} />
  </Route>
)

Router.run(Routes, function(Handler) {
  React.render(<Handler/>, document.body);
});
