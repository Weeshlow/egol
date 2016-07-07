(function() {

	'use strict';

	var _ = require('lodash');
	var esper = require('esper');
	var glm = require('gl-matrix');
	var Listener = require('./scripts/Listener');
	var Organism = require('./scripts/Organism');

	var FIELD_OF_VIEW = 60 * (Math.PI / 180);
	var MIN_Z = 0.1;
	var MAX_Z = 1000;

	var canvas;
	var gl;
	var view;
	var projection;
	var viewport;
	var listener;

	var organisms = {};
	var updates = {};
	var last;

	function getWindowSize() {
		var devicePixelRatio = window.devicePixelRatio || 1;
		return {
			width: window.innerWidth * devicePixelRatio,
			height: window.innerHeight * devicePixelRatio
		};
	}

	window.addEventListener('resize', () => {
		if (viewport) {
			var size = getWindowSize();
			viewport.resize(size.width , size.height);
			projection = glm.mat4.perspective(
				projection,
				FIELD_OF_VIEW,
				size.width / size.height,
				MIN_Z,
				MAX_Z);
		}
	});

	function render(organism, update, t) {
		if (update) {
			organism = organism.interpolate(update, t);
		}
		//organism.draw();
	}

	function processFrame() {
		var stamp = Date.now();
		var delta = stamp - last;
		_.forIn(organisms, organism => {
			// get update if it is available
			var update = updates[organism.id];
			// render the interpolated state
			render(organism, update, delta);
		});
		requestAnimationFrame(processFrame);
	}

	function initializeState() {
		// window size
		var size = getWindowSize();
		// viewport
		viewport = new esper.Viewport();
		viewport.resize(size.width, size.height);
		// view matrix
		view = glm.mat4.create(1);
		// projection matrix
		projection = glm.mat4.perspective(
			glm.mat4.create(),
			FIELD_OF_VIEW,
			size.width / size.height,
			MIN_Z,
			MAX_Z);
	}

	function handleState(orgs) {
		// clear current state
		organisms = {};
		_.forIn(orgs, org => {
			organisms[org.id] = new Organism(org);
		});
	}

	function handleUpdate(newUpdates) {
		// apply last updates to state
		_.forIn(updates, (update, id) => {
			organisms[id].update(update);
		});
		// store new updates to interpolate to
		updates = newUpdates;
		// update timestamp
		last = Date.now();
	}

	window.start = () => {
		// get canvas
		canvas = document.getElementById('glcanvas');
		// get WebGL context
		gl = esper.WebGLContext.get(canvas);
		// only continue if WebGL is available
		if (gl) {
			// create websocket connection
			listener = new Listener(
				'connect',
				// message handler
				msg => {
					console.log(msg);
					if (msg.type === 'state') {
						handleState(msg.data);
					} else if (msg.type === 'update') {
						handleUpdate(msg.data);
					}
				},
				// on connections
				() => {
					// initiaze rendering
					initializeState();
					// initiate draw loop
					processFrame();
				});
		}
	};

}());
