import React, {Component} from 'react'
import TicketKey from './TicketKey'
import TicketType from './TicketType'
import TicketStatus from './TicketStatus'
import TicketAssignee from './TicketAssignee'
import TicketReporter from './TicketReporter'
import TicketDate from './TicketDate'

export default class TicketDetail extends Component {
	render() {
		return (
			<div className="details container">
				<div className="col-md-6 col-sm-6">
					<TicketStatus status={this.props.status} />
					<TicketDate title="Created Date" date={this.props.createdAt} />
					<TicketDate title="Updated Date" date={this.props.updatedAt} />
				</div>

				<div className="col-md-6 col-sm-6">
					<TicketReporter reporter={this.props.reporter} />
					<TicketAssignee assignee={this.props.assignee} />
					<TicketType ticketType={this.props.issueType} />
				</div>
			</div>
		)
	}
}
