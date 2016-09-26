import React, {Component} from 'react'
import UserStub from '../Users/UserStub'

export default class TicketReporter extends Component {
	render() {
		return (
			<div className="row assignee-row">
				<div className="col-md-6">
					<span className="text-muted">Assignee:</span>
				</div>
				<div className="col-md-6">
					<UserStub {...this.props.assignee} />
				</div>
			</div> 
		)
	}
}
