(function() {

	'use strict';

	var RETRY_INTERVAL = 5000;

	function getHost() {
		var loc = window.location;
		var new_uri;
		if (loc.protocol === 'https:') {
			new_uri = 'wss:';
		} else {
			new_uri = 'ws:';
		}
		return new_uri + '//' + loc.host + loc.pathname;
	}

	function establishConnection(requestor, callback) {
		requestor.socket = new WebSocket(getHost() + requestor.url);
		// on open
		requestor.socket.onopen = () => {
			requestor.isOpen = true;
			console.log('Websocket connection established');
			callback.apply(this, arguments);
		};
		// on message
		requestor.socket.onmessage = event => {
			var res = JSON.parse(event.data);
			requestor.listener(res);
		};
		// on close
		requestor.socket.onclose = () => {
			// log close only if connection was ever open
			if (requestor.isOpen) {
				console.warn('Websocket connection closed, attempting to re-connect in ' + RETRY_INTERVAL);
			}
			// flag as closed
			requestor.socket = null;
			requestor.isOpen = false;
			// attempt to re-establish connection
			setTimeout(() => {
				establishConnection(requestor, callback);
			}, RETRY_INTERVAL);
		};
	}

	class Listener {
		constructor(url, listener, callback) {
			this.url = url;
			this.listener = listener;
			this.isOpen = false;
			establishConnection(this, callback);
		}
		close() {
			this.socket.onclose = null;
			this.socket.close();
			this.socket = null;
		}
	}

	module.exports = Listener;

}());
