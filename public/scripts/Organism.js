(function() {

    'use strict';

    var esper = require('esper');
    var glm = require('gl-matrix');
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
            if (!spec) {
                throw 'No organism argument';
            }
            this.id = spec.id;
            this.state = new State(spec.state);
            this.attributes = new Attributes(spec.attributes);
            this.buffer = spec.buffer || createCircleBuffer(60);
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
            this.state.update(update.state);
        }
        draw() {
            this.buffer.bind();
            this.buffer.draw();
            this.buffer.unbind();
        }
        color() {
            var energy = this.state.energy;
            switch (this.state.type) {
                case 'alive':
                    return [0.2 * energy, 1.0 * energy, 0.3 * energy, 1.0];
                case 'reproducing':
                    return [0.9, 0.6, 0.4, 1.0];
                case 'dead':
                    return [0.4, 0.4, 0.4, 1.0];
                default:
                    return [1.0, 1.0, 0.0, 1.0];
            }
        }
        matrix() {
            var translation = this.state.position;
            var rotation = glm.quat.rotateZ(
                glm.quat.create(),
                glm.quat.identity(glm.quat.create()),
                this.state.rotation);
            var scale = glm.vec3.fromValues(
                this.state.size,
                this.state.size,
                this.state.size);
            return glm.mat4.fromRotationTranslationScale(
                glm.mat4.create(),
                // rotation
                rotation,
                // translation
                translation,
                // scale
                scale);
        }
        perception(weight) {
            var translation = this.state.position;
            var rotation = glm.quat.identity(glm.quat.create());
            var scale = glm.vec3.fromValues(
                this.attributes.perception * weight,
                this.attributes.perception * weight,
                this.attributes.perception * weight);
            return glm.mat4.fromRotationTranslationScale(
                glm.mat4.create(),
                // rotation
                rotation,
                // translation
                translation,
                // scale
                scale);
        }
        range() {
            var translation = this.state.position;
            var rotation = glm.quat.identity(glm.quat.create());
            var scale = glm.vec3.fromValues(
                this.state.size + this.attributes.range,
                this.state.size + this.attributes.range,
                this.state.size + this.attributes.range);
            return glm.mat4.fromRotationTranslationScale(
                glm.mat4.create(),
                // rotation
                rotation,
                // translation
                translation,
                // scale
                scale);
        }
    }

    module.exports = Organism;

}());
