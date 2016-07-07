(function() {

	'use strict';

	var _ = require('lodash');
	var esper = require('esper');
	var glm = require('gl-matrix');
	var Listener = require('./scripts/Listener');
	var Organism = require('./scripts/Organism');

	var FRAME_MS = 2000;

	var canvas;
	var gl;
	var view;
	var projection;
	var viewport;
	var shader;
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

	function onResize() {
		if (viewport) {
			var size = getWindowSize();
			viewport.resize(size.width , size.height);
			projection = glm.mat4.ortho(
				projection,
				0, 1,
				0, 1,
				-1, 1);
		}
	}

	window.addEventListener('resize', onResize);

	function render(organism, update, t) {
		if (update) {
			organism = organism.interpolate(update, t);
		}
		viewport.push();
		shader.push();

		shader.setUniform('uProjectionMatrix', projection);
		shader.setUniform('uViewMatrix', view);
		shader.setUniform('uModelMatrix', organism.matrix());

		organism.draw();

		shader.pop();
		viewport.pop();
	}

	function processFrame() {
		var stamp = Date.now();
		var delta = stamp - last;
		var t = Math.min(1.0, delta / FRAME_MS);
		_.forIn(organisms, organism => {
			// get update if it is available
			var update = updates[organism.id];
			// render the interpolated state
			render(organism, update, t);
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
		projection = glm.mat4.ortho(
			glm.mat4.create(),
			0, 1,
			0, 1,
			-1, 1);
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

			initializeState();

			shader = new esper.Shader({
				vert: 'shaders/organism.vert',
				frag: 'shaders/organism.frag'
			}, function() {
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
						last = Date.now();
						processFrame();
					});
			});

		}
	};

}());
