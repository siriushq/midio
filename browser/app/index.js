import "babel-polyfill"
import "./less/main.less"
import "@fortawesome/fontawesome-free/css/all.css"
import "material-design-iconic-font/dist/css/material-design-iconic-font.min.css"

import React from "react"
import ReactDOM from "react-dom"
import { Router, Route } from "react-router-dom"
import { Provider } from "react-redux"

import history from "./js/history"
import configureStore from "./js/store/configure-store"
import hideLoader from "./js/loader"
import App from "./js/App"

const store = configureStore()

ReactDOM.render(
  <Provider store={store}>
    <Router history={history}>
      <App />
    </Router>
  </Provider>,
  document.getElementById("root")
)

hideLoader()
