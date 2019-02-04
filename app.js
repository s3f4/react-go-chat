let channels = [
    { "name": "Hardware Support" },
    { "name": "Software Support" }
]

class Channel extends React.Component {
    onClick() {
        console.log("i was clicked", this.props.name);
    }
    render() {
        return <li onClick={this.onClick.bind(this)}>{this.props.name}</li>
    }
}

class ChannelForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {};
    }
    onChange(e) {
        this.setState({
            channelName: e.target.value
        });
    }
    onSubmit(e) {
        let { channelName } = this.state;
        console.log(chanelName);
        channels.push({
            "name": channelName
        });
        this.setState({
            "channelName": ""
        });
        e.preventDefault();
    }
    render() {
        return <form onSubmit={this.onSubmit.bind(this)}>
            <input type="text" onChange={this.onChange.bind(this)} value={this.state.channelName} />
        </form>
    }
}

class ChannelList extends React.Component {
    render() {
        return <ul>
            {
                this.props.channels.map(channel => {
                    return <Channel name={channel.name} />
                })
            }
        </ul>
    }
}

class ChannelSection extends React.Component {
    render() {
        return <div>
            <ChannelList channels={channels} />
            <ChannelForm />
        </div>
    }
}

ReactDOM.render(<ChannelSection />, document.getElementById('app'));
