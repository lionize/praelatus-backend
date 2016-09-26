import React, {Component} from 'react'
import { Router, IndexRoute, Route } from 'react-router' 
import createHashHistory from 'history/lib/createHashHistory'
import Ticket from './Ticket'

const ticket = {
  id: 1,
  createdAt: "2016-07-11T21:18:02.1569Z",
  updatedAt: "2016-07-11T21:18:02.1569Z",
  issueKey: "TEST-1",
  summary: "This is a test issue #1",
  description: "A very find day for some testing.",
  issueType: "BUG",
  reporter: 1,
  assignee: 1,
  status: {
    name: "",
    type: 0
  }
}

export default class App extends Component {
  render() {
    return (
		// <Router history={createHashHistory({queryKey: false})}>
		// </Router>,
		// document.getElementById('app')
		<Ticket {...ticket} />
    )
  }
}

