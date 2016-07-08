(function() {

	'use strict';

	var esper = require('esper');
	var glm = require('gl-matrix');
	var Attributes = require('./Attributes');
	var State = require('./State');
	var buffer = null;

	function createCircleBuffer(numSegments) {
		if (!buffer) {
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
			var buff = new ArrayBuffer((numSegments + 2) * 2 * COMPONENT_BYTE_SIZE);
			var positions = new Float32Array(buff);
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
			buffer = new esper.VertexBuffer(positions, pointers, options);
		}
		return buffer;
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
		size() {
			return 0.005 + (this.state.maturity * this.state.energy * 0.01);
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
			if (this.state.type === 'dead') {
				return [0.1, 0.1, 0.1, 1.0];
			}
			switch (this.attributes.family) {
				case 0:
					return [0.9, 0.6, 0.4, 1.0];
				case 1:
					return [0.2, 0.9, 0.2, 1.0];
				case 2:
					return [0.2, 0.2, 0.9, 1.0];
				case 3:
					return [0.9, 0.2, 0.2, 1.0];
				case 4:
					return [0.9, 0.2, 0.9, 1.0];
			}
		}
		matrix() {
			var translation = this.state.position;
			var rotation = glm.quat.identity(glm.quat.create());
			var scale = glm.vec3.fromValues(
				this.size(),
				this.size(),
				this.size());
			return glm.mat4.fromRotationTranslationScale(
				glm.mat4.create(),
				// rotation
				rotation,
				// translation
				translation,
				// scale
				scale);
		}
		perception() {
			var translation = this.state.position;
			var rotation = glm.quat.identity(glm.quat.create());
			var scale = glm.vec3.fromValues(
				this.size() + this.attributes.perception,
				this.size() + this.attributes.perception,
				this.size() + this.attributes.perception);
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
				this.size() + this.attributes.range,
				this.size() + this.attributes.range,
				this.size() + this.attributes.range);
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
