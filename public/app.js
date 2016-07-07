(function() {

	'use strict';

	var esper = require('esper');
	var glm = require('gl-matrix');
	var Listener = require('./scripts/Listener');

	var FIELD_OF_VIEW = 60 * (Math.PI / 180);
	var MIN_Z = 0.1;
	var MAX_Z = 1000;

	var canvas;
	var gl;
	var view;
	var projection;
	var viewport;
	var listener;
	// var organisms;

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

	function processFrame() {
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
					if (msg.type === 'state') {
						console.log(msg);
					} else if (msg.type === 'update') {
						console.log(msg);
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
