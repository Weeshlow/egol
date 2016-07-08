(function() {

	'use strict';

	var glm = require('gl-matrix');

	class State {
		constructor(spec) {
			if (!spec) {
				throw 'No state argument';
			}
			// type
			this.type = spec.type;
			// energy
			this.energy = (spec.energy !== undefined) ? spec.energy : 1.0;
			// position
			this.position = spec.position ? glm.vec3.fromValues(
				spec.position.x || spec.position[0] || 0,
				spec.position.y || spec.position[1] || 0,
				spec.position.z || spec.position[2] || 0) : glm.vec3.create();
			// Maturity
			this.maturity = spec.maturity;
		}
		interpolate(update, t) {
			var state = update.state;

			// interpolate between current state and update based on a t value from 0 to 1
			// get distance vector
			let diff = glm.vec3.sub(glm.vec3.create(), state.position, this.position);
			// scale by t value
			diff = glm.vec3.scale(diff, diff, t);
			// get update position
			let position = glm.vec3.add(glm.vec3.create(), this.position, diff);

			// interpolate maturity
			let maturity = this.maturity * (1-t) + state.maturity * t;

			// interpolate energy
			let energy = this.energy * (1-t) + state.energy * t;

			return new State({
				type: state.type,
				energy: energy,
				position: position,
				maturity: maturity
			});
		}
		update(update) {
			this.type = update.type;
			this.energy = update.energy;
			this.position = update.position;
			this.maturity = update.maturity;
		}
	}

	module.exports = State;

}());
