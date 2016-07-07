(function() {

	'use strict';

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

	var organisms;
	var updates;
	var last;

	function createCircleBuffer(numSegments) {
		var COMPONENT_BYTE_SIZE = 4;
        var theta = (2 * Math.PI) / numSegments;
        var radius = 1.0;
        // precalculate sine and cosine
        var c = Math.cos(theta);
        var s = Math.sin(theta);
        var t;
        // start at angle = 0
        var x = radius;
        var y = 0;
        var buffer = new ArrayBuffer((numSegments + 2) * 2 * COMPONENT_BYTE_SIZE);
        var positions = new Float32Array(buffer);
        positions[0] = 0;
        positions[1] = 0;
        positions[positions.length-2] = radius;
        positions[positions.length-1] = 0;
        for(var i = 0; i < numSegments; i++) {
            positions[(i+1)*2] = x;
            positions[(i+1)*2+1] = y;
            // apply the rotation
            t = x;
            x = c * x - s * y;
            y = s * t + c * y;
        }
        var pointers = {
			0: {
	            size: 2,
	            type: 'FLOAT'
	        }
		};
        var options = {
            mode: 'TRIANGLE_FAN'
        };
        return new esper.VertexBuffer(positions, pointers, options);
    }

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
		var stamp = Date.now();
		var delta = stamp - last;


		last = stamp;
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
		orgs.forEach(org => {
			organisms[org.id] = new Organism(org);
		});
	}

	function handleUpdate(newUpdates) {
		// apply last updates to state
		updates.forEach(update => {
			organisms[update.id].update(update);
		});
		// store new updates to interpolate to
		updates = newUpdates;
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
						handleState(msg);
					} else if (msg.type === 'update') {
						console.log(msg);
						handleUpdate(msg);
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
