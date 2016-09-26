import React, {Component} from 'react'
import TicketDescription from './Ticket/TicketDescription'
import TicketDetails from './Ticket/TicketDetails'

export default class TicketDetailPage extends Component {
  render() {
    return <div className="ticket container">
		<div className="ticket-heading container">
			<h3>{this.props.issueKey} <small>{this.props.summary}</small></h3>
		</div>

		<TicketDetails {...this.props} />
      <div className="container">
        <TicketDescription description={this.props.description} />
      </div>
    </div>
  }
}
