(function() {

    'use strict';

    var esper = require('esper');
    var Attributes = require('./Attributes');
    var State = require('./State');

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
            this.state = new State(spec.state);
            this.attributes = new Attributes(spec.attributes);
            this.buffer = createCircleBuffer(24);
        }
        interpolate(update, t) {
            return new Organism({
                id: this.id,
                state: this.state.interpolate(update, t),
                attributes: this.attributes,
                buffer: this.buffer
            });
        }
        update(update) {
            this.state = new State(update.state);
            if (update.attributes) {
                this.positions = update.position;
            }
            if (update.attributes) {
                this.rotation = update.rotation;
            }
            if (update.attributes) {
                this.attributes = update.attributes;
            }
        }
        draw() {
            this.buffer.bind();
            this.buffer.draw();
            this.buffer.unbind();
        }
        matrix() {
            return this.state.matrix();
        }
    }

    module.exports = Organism;

}());
