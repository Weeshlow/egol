(function() {

    'use strict';

    var esper = require('esper');

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

    class Organism {
        constructor(spec) {
            this.id = spec.id;
            this.position = spec.position;
            this.rotation = spec.rotation;
            this.state = spec.state;
            this.attributes = spec.attributes;
            this.buffer = createCircleBuffer(24);
        }
        interpolate(update, t) {
            // iterpolate between current state and update based on a t value from 0 to 1
            console.log(update, t);

            return new Organism();
        }
        apply(update) {
            // apply an update to the organism
            console.log(update);
        }
        draw() {
            this.buffer.bind();
            this.buffer.draw();
            this.buffer.unbind();
        }
    }

    module.exports = Organism;

}());
